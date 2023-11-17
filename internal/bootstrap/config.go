package bootstrap

import (
	"encoding/json"
	"github.com/AkashiCoin/gin-template/cmd/flags"
	"github.com/AkashiCoin/gin-template/internal/conf"
	"github.com/AkashiCoin/gin-template/pkg/utils"
	"github.com/caarlos0/env/v9"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

const (
	ConfigFileName = "config.json"
)

func InitConfig() {
	configPath := filepath.Join(flags.DataDir, ConfigFileName)
	if !utils.Exists(configPath) {
		log.Infof("Conf file not exists, creating default Conf file")
		_, err := utils.CreateNestedFile(configPath)
		if err != nil {
			log.Fatalf("failed to create Conf file: %+v", err)
		}
		conf.Conf = conf.DefaultConfig()
		if !utils.WriteJsonToFile(configPath, conf.Conf) {
			log.Fatalf("failed to create default Conf file")
		}
	} else {
		ConfBytes, err := os.ReadFile(configPath)
		if err != nil {
			log.Fatalf("reading Conf file error: %+v", err)
		}
		conf.Conf = conf.DefaultConfig()
		err = json.Unmarshal(ConfBytes, conf.Conf)
		if err != nil {
			log.Fatalf("load Conf error: %+v", err)
		}
		// update Conf.json struct
		confBody, err := json.MarshalIndent(conf.Conf, "", "  ")
		if err != nil {
			log.Fatalf("marshal Conf error: %+v", err)
		}
		err = os.WriteFile(configPath, confBody, 0o777)
		if err != nil {
			log.Fatalf("update Conf struct error: %+v", err)
		}
	}
	if !conf.Conf.Force {
		confFromEnv()
	}
}

func confFromEnv() {
	log.Infof("load config from env...")
	if err := env.ParseWithOptions(conf.Conf, env.Options{}); err != nil {
		log.Fatalf("load config from env error: %+v", err)
	}
}
