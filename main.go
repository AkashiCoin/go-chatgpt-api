package main

import (
	"embed"
	"github.com/AkashiCoin/gin-template/cmd"
)

//go:embed web/dist
var buildFS embed.FS

//go:embed web/dist/index.html
var indexPage []byte

func main() {
	cmd.Execute(buildFS, indexPage)
}
