package conf

import (
	"regexp"
	"time"
)

var (
	StartTime = time.Now().Unix()
	AppName   = "go-chatgpt-api"
	Version   = "dev"
)

var (
	Conf *Config
)

var PrivacyReg []*regexp.Regexp
