package conf

import (
	"regexp"
	"time"
)

var (
	StartTime = time.Now().Unix()
	AppName   = "gin-template"
	Version   = "dev"
)

var (
	Conf *Config
)

var PrivacyReg []*regexp.Regexp
