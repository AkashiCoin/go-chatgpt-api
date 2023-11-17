package cmd

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"github.com/AkashiCoin/go-chatgpt-api/cmd/flags"
	"github.com/AkashiCoin/go-chatgpt-api/internal/conf"
	"github.com/AkashiCoin/go-chatgpt-api/pkg/utils"
	"github.com/AkashiCoin/go-chatgpt-api/server/router"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func Execute(buildFS embed.FS, indexPage []byte) {
	Init()
	if !flags.Debug && !flags.Dev {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.LoggerWithWriter(log.StandardLogger().Out), gin.RecoveryWithWriter(log.StandardLogger().Out))
	router.SetRouter(r, buildFS, indexPage)
	var httpSrv, httpsSrv, unixSrv *http.Server

	if conf.Conf.Scheme.HttpPort != -1 {
		httpBase := fmt.Sprintf("%s:%d", conf.Conf.Scheme.Address, conf.Conf.Scheme.HttpPort)
		utils.Log.Infof("start HTTP server @ %s", httpBase)
		httpSrv = &http.Server{Addr: httpBase, Handler: r}
		go func() {
			err := httpSrv.ListenAndServe()
			if err != nil && err != http.ErrServerClosed {
				utils.Log.Fatalf("failed to start http: %s", err.Error())
			}
		}()
	}
	if conf.Conf.Scheme.HttpsPort != -1 {
		httpsBase := fmt.Sprintf("%s:%d", conf.Conf.Scheme.Address, conf.Conf.Scheme.HttpsPort)
		utils.Log.Infof("start HTTPS server @ %s", httpsBase)
		httpsSrv = &http.Server{Addr: httpsBase, Handler: r}
		go func() {
			err := httpsSrv.ListenAndServeTLS(conf.Conf.Scheme.CertFile, conf.Conf.Scheme.KeyFile)
			if err != nil && err != http.ErrServerClosed {
				utils.Log.Fatalf("failed to start https: %s", err.Error())
			}
		}()
	}
	if conf.Conf.Scheme.UnixFile != "" {
		utils.Log.Infof("start unix server @ %s", conf.Conf.Scheme.UnixFile)
		unixSrv = &http.Server{Handler: r}
		go func() {
			listener, err := net.Listen("unix", conf.Conf.Scheme.UnixFile)
			if err != nil {
				utils.Log.Fatalf("failed to listen unix: %+v", err)
			}
			// set socket file permission
			mode, err := strconv.ParseUint(conf.Conf.Scheme.UnixFilePerm, 8, 32)
			if err != nil {
				utils.Log.Errorf("failed to parse socket file permission: %+v", err)
			} else {
				err = os.Chmod(conf.Conf.Scheme.UnixFile, os.FileMode(mode))
				if err != nil {
					utils.Log.Errorf("failed to chmod socket file: %+v", err)
				}
			}
			err = unixSrv.Serve(listener)
			if err != nil && err != http.ErrServerClosed {
				utils.Log.Fatalf("failed to start unix: %s", err.Error())
			}
		}()
	}
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 1 second.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	utils.Log.Println("Shutdown server...")
	Release()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	if conf.Conf.Scheme.HttpPort != -1 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := httpSrv.Shutdown(ctx); err != nil {
				utils.Log.Fatal("HTTP server shutdown err: ", err)
			}
		}()
	}
	if conf.Conf.Scheme.HttpsPort != -1 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := httpsSrv.Shutdown(ctx); err != nil {
				utils.Log.Fatal("HTTPS server shutdown err: ", err)
			}
		}()
	}
	if conf.Conf.Scheme.UnixFile != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := unixSrv.Shutdown(ctx); err != nil {
				utils.Log.Fatal("Unix server shutdown err: ", err)
			}
		}()
	}
	wg.Wait()
	utils.Log.Println("Server exit")
}

func init() {
	flag.BoolVar(&flags.Debug, "debug", false, "start with debug mode")
	flag.BoolVar(&flags.Dev, "dev", false, "start with dev mode")
	flag.StringVar(&flags.DataDir, "data", "data", "data folder")
	flag.BoolVar(&flags.LogStd, "log-std", false, "Force to log to std")
}
