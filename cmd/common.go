package cmd

import (
	"github.com/AkashiCoin/go-chatgpt-api/internal/bootstrap"
	"github.com/AkashiCoin/go-chatgpt-api/internal/db"
	"github.com/AkashiCoin/go-chatgpt-api/pkg/utils"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strconv"
)

func Init() {
	bootstrap.InitConfig()
	bootstrap.InitApi()
	bootstrap.Log()
	bootstrap.InitDB()
}

func Release() {
	db.Close()
}

var pid = -1
var pidFile string

func initDaemon() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exPath := filepath.Dir(ex)
	_ = os.MkdirAll(filepath.Join(exPath, "daemon"), 0700)
	pidFile = filepath.Join(exPath, "daemon/pid")
	if utils.Exists(pidFile) {
		bytes, err := os.ReadFile(pidFile)
		if err != nil {
			log.Fatal("failed to read pid file", err)
		}
		id, err := strconv.Atoi(string(bytes))
		if err != nil {
			log.Fatal("failed to parse pid data", err)
		}
		pid = id
	}
}
