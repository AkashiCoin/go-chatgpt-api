package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AkashiCoin/go-chatgpt-api/internal/conf"
	"github.com/bogdanfinn/fhttp/cookiejar"
	"io"
	"strings"
	"time"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

const (
	ChatGPTApiPrefix    = "/chatgpt"
	ImitateApiPrefix    = "/imitate/v1"
	ChatGPTApiUrlPrefix = "https://chat.openai.com"

	PlatformApiPrefix    = "/platform"
	PlatformApiUrlPrefix = "https://api.openai.com"

	gpt4Model                          = "gpt-4"
	Auth                               = "auth"
	defaultErrorMessageKey             = "errorMessage"
	AuthorizationHeader                = "Authorization"
	XAuthorizationHeader               = "X-Authorization"
	ContentType                        = "application/x-www-form-urlencoded"
	UserAgent                          = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36"
	Auth0Url                           = "https://auth0.openai.com"
	LoginUsernameUrl                   = Auth0Url + "/u/login/identifier?state="
	LoginPasswordUrl                   = Auth0Url + "/u/login/password?state="
	ParseUserInfoErrorMessage          = "failed to parse user login info"
	GetAuthorizedUrlErrorMessage       = "failed to get authorized url"
	EmailInvalidErrorMessage           = "email is not valid"
	EmailOrPasswordInvalidErrorMessage = "email or password is not correct"
	GetAccessTokenErrorMessage         = "failed to get access token"
	defaultTimeoutSeconds              = 600 // 10 minutes

	EmailKey                       = "email"
	AccountDeactivatedErrorMessage = "account %s is deactivated"

	ReadyHint = "service go-chatgpt-api is ready"

	refreshPuidErrorMessage = "failed to refresh PUID"
)

var (
	Client         tls_client.HttpClient
	ArkoseClient   tls_client.HttpClient
	PUID           string
	ProxyUrl       string
	arkoseTokenUrl string
)

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthLogin interface {
	GetAuthorizedUrl(csrfToken string) (string, int, error)
	GetState(authorizedUrl string) (string, int, error)
	CheckUsername(state string, username string) (int, error)
	CheckPassword(state string, username string, password string) (string, int, error)
	GetAccessToken(code string) (string, int, error)
	GetAccessTokenFromHeader(c *gin.Context) (string, int, error)
}

func Init() {
	arkoseTokenUrl = conf.Conf.ArkoseTokenUrl
	Client, _ = tls_client.NewHttpClient(tls_client.NewNoopLogger(), []tls_client.HttpClientOption{
		tls_client.WithCookieJar(tls_client.NewCookieJar()),
		tls_client.WithTimeoutSeconds(defaultTimeoutSeconds),
		tls_client.WithClientProfile(profiles.Okhttp4Android13),
	}...)
	ArkoseClient = getHttpClient()

	setupPUID()
}

func NewHttpClient() tls_client.HttpClient {
	client := getHttpClient()

	ProxyUrl = conf.Conf.Proxy
	if ProxyUrl != "" {
		client.SetProxy(ProxyUrl)
	}

	return client
}

func getHttpClient() tls_client.HttpClient {
	client, _ := tls_client.NewHttpClient(tls_client.NewNoopLogger(), []tls_client.HttpClientOption{
		tls_client.WithCookieJar(tls_client.NewCookieJar()),
		tls_client.WithClientProfile(profiles.Okhttp4Android13),
	}...)
	return client
}

func Proxy(c *gin.Context) {
	url := c.Request.URL.Path
	if strings.Contains(url, ChatGPTApiPrefix) {
		url = strings.ReplaceAll(url, ChatGPTApiPrefix, ChatGPTApiUrlPrefix)
	} else if strings.Contains(url, ImitateApiPrefix) {
		url = strings.ReplaceAll(url, ImitateApiPrefix, ChatGPTApiUrlPrefix+"/backend-api")
	} else if strings.Contains(url, PlatformApiPrefix) {
		url = strings.ReplaceAll(url, PlatformApiPrefix, PlatformApiUrlPrefix)
	} else {
		url = fmt.Sprintf("%s%s", ChatGPTApiUrlPrefix, url)
	}

	method := c.Request.Method
	queryParams := c.Request.URL.Query().Encode()
	if queryParams != "" {
		url += "?" + queryParams
	}
	var req *http.Request
	if method == http.MethodGet {
		req, _ = http.NewRequest(http.MethodGet, url, nil)
	} else {
		body, _ := io.ReadAll(c.Request.Body)
		req, _ = http.NewRequest(method, url, bytes.NewReader(body))
	}

	for _, cookie := range c.Request.Cookies() {
		req.AddCookie(&http.Cookie{
			Name:  cookie.Name,
			Value: cookie.Value,
		})
	}
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set(AuthorizationHeader, GetAccessToken(c))
	jar, _ := cookiejar.New(nil)
	Client.SetCookieJar(jar)
	resp, err := Client.Do(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ReturnMessage(err.Error()))
		return
	}

	defer resp.Body.Close()
	c.Status(resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusUnauthorized {
			logger.Error(fmt.Sprintf(AccountDeactivatedErrorMessage, c.GetString(EmailKey)))
		}

		responseMap := make(map[string]interface{})
		json.NewDecoder(resp.Body).Decode(&responseMap)
		c.AbortWithStatusJSON(resp.StatusCode, responseMap)
		return
	}
	for _, setCookie := range resp.Header.Values("Set-Cookie") {
		c.Writer.Header().Add("Set-Cookie", setCookie)
	}
	io.Copy(c.Writer, resp.Body)
}

func ReturnMessage(msg string) gin.H {
	logger.Warn(msg)

	return gin.H{
		defaultErrorMessageKey: msg,
	}
}

func GetAccessToken(c *gin.Context) string {
	accessToken := c.Request.Header.Get(AuthorizationHeader)
	if !strings.HasPrefix(accessToken, "Bearer") {
		return "Bearer " + accessToken
	}

	return accessToken
}

func GetArkoseToken(model string) (string, error) {
	version := "/gpt3"
	if strings.HasPrefix(model, gpt4Model) {
		version = "/gpt4"
	}
	if strings.HasPrefix(model, Auth) {
		version = "/auth"
	}
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", arkoseTokenUrl, version), nil)
	resp, err := ArkoseClient.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	defer resp.Body.Close()
	responseMap := make(map[string]interface{})
	json.NewDecoder(resp.Body).Decode(&responseMap)
	arkoseToken, ok := responseMap["token"]
	if !ok || arkoseToken == "" {
		return "", errors.New(fmt.Sprintf("failed to get arkose token: %v", responseMap))
	}

	return arkoseToken.(string), nil
}

func setupPUID() {
	username := conf.Conf.Session.Username
	password := conf.Conf.Session.Password
	accessToken := conf.Conf.Session.AccessToken
	puid := conf.Conf.Session.Puid
	if accessToken != "" {
		go func() {
			for {
				authenticator := NewAuthenticator("", "", ProxyUrl)
				authenticator.Result.AccessToken = accessToken

				puid, err := authenticator.GetPUID()
				if err != nil {
					logger.Error(refreshPuidErrorMessage)
					return
				}
				PUID = puid

				time.Sleep(time.Hour * 24 * 3)
			}
		}()
	} else if puid != "" {
		PUID = puid
	} else if username != "" && password != "" {
		go func() {
			for {
				authenticator := NewAuthenticator(username, password, ProxyUrl)
				if err := authenticator.Begin(); err != nil {
					logger.Warn(fmt.Sprintf("%s: %s", refreshPuidErrorMessage, err.Details))
					return
				}

				accessToken := authenticator.GetAccessToken()
				if accessToken == "" {
					logger.Error(refreshPuidErrorMessage)
					return
				}

				puid, err := authenticator.GetPUID()
				if err != nil {
					logger.Error(refreshPuidErrorMessage)
					return
				}

				PUID = puid

				time.Sleep(time.Hour * 24 * 7)
			}
		}()
	}
}
