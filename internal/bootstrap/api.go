package bootstrap

import (
	"github.com/AkashiCoin/go-chatgpt-api/pkg/api"
	"github.com/AkashiCoin/go-chatgpt-api/pkg/api/chatgpt"
)

func InitApi() {
	api.Init()
	chatgpt.Init()
}
