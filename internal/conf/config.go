package conf

import (
	"github.com/AkashiCoin/go-chatgpt-api/cmd/flags"
	"path/filepath"
)

type Database struct {
	Type        string `json:"type" env:"DB_TYPE"`
	Host        string `json:"host" env:"DB_HOST"`
	Port        int    `json:"port" env:"DB_PORT"`
	User        string `json:"user" env:"DB_USER"`
	Password    string `json:"password" env:"DB_PASS"`
	Name        string `json:"name" env:"DB_NAME"`
	DBFile      string `json:"db_file" env:"DB_FILE"`
	TablePrefix string `json:"table_prefix" env:"DB_TABLE_PREFIX"`
	SSLMode     string `json:"ssl_mode" env:"DB_SSL_MODE"`
}

type Scheme struct {
	Address      string `json:"address" env:"ADDR"`
	HttpPort     int    `json:"http_port" env:"HTTP_PORT"`
	HttpsPort    int    `json:"https_port" env:"HTTPS_PORT"`
	ForceHttps   bool   `json:"force_https" env:"FORCE_HTTPS"`
	CertFile     string `json:"cert_file" env:"CERT_FILE"`
	KeyFile      string `json:"key_file" env:"KEY_FILE"`
	UnixFile     string `json:"unix_file" env:"UNIX_FILE"`
	UnixFilePerm string `json:"unix_file_perm" env:"UNIX_FILE_PERM"`
}

type LogConfig struct {
	Enable     bool   `json:"enable" env:"LOG_ENABLE"`
	Name       string `json:"name" env:"LOG_NAME"`
	MaxSize    int    `json:"max_size" env:"MAX_SIZE"`
	MaxBackups int    `json:"max_backups" env:"MAX_BACKUPS"`
	MaxAge     int    `json:"max_age" env:"MAX_AGE"`
	Compress   bool   `json:"compress" env:"COMPRESS"`
}

type SessionConfig struct {
	Username     string `json:"username" env:"USERNAME"`
	Password     string `json:"password" env:"PASSWORD"`
	AccessToken  string `json:"access_token" env:"ACCESS_TOKEN"`
	SessionToken string `json:"session_token" env:"SESSION_TOKEN"`
	Puid         string `json:"puid" env:"PUID"`
}

type Config struct {
	Force          bool          `json:"force"`
	Cdn            string        `json:"cdn" env:"CDN"`
	Proxy          string        `json:"proxy" env:"PROXY"`
	AutoContinue   bool          `json:"auto_continue" env:"AUTO_CONTINUE"`
	ArkoseTokenUrl string        `json:"arkose_token_url" env:"ARKOSE_TOKEN_URL"`
	Session        SessionConfig `json:"session"`
	Scheme         Scheme        `json:"scheme"`
	Database       Database      `json:"database"`
	Log            LogConfig     `json:"log"`
}

func DefaultConfig() *Config {
	logPath := filepath.Join(flags.DataDir, "log/log.log")
	dbPath := filepath.Join(flags.DataDir, "data.db")
	return &Config{
		Scheme: Scheme{
			Address:    "0.0.0.0",
			UnixFile:   "",
			HttpPort:   8755,
			HttpsPort:  -1,
			ForceHttps: false,
			CertFile:   "",
			KeyFile:    "",
		},
		Database: Database{
			Type:        "sqlite3",
			Port:        0,
			TablePrefix: "x_",
			DBFile:      dbPath},
		Log: LogConfig{
			Enable:     true,
			Name:       logPath,
			MaxSize:    50,
			MaxBackups: 30,
			MaxAge:     28,
		},
	}
}
