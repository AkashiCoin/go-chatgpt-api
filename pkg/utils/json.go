package utils

import (
	"encoding/json"
	stdjson "encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
)

// WriteJsonToFile write struct to json file
func WriteJsonToFile(dst string, data interface{}, std ...bool) bool {
	str, err := json.MarshalIndent(data, "", "  ")
	if len(std) > 0 && std[0] {
		str, err = stdjson.MarshalIndent(data, "", "  ")
	}
	if err != nil {
		log.Errorf("failed convert Conf to []byte:%s", err.Error())
		return false
	}
	err = os.WriteFile(dst, str, 0777)
	if err != nil {
		log.Errorf("failed to write json file:%s", err.Error())
		return false
	}
	return true
}
