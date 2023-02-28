package main

import (
	"sheet-retrieve/config"
	"sheet-retrieve/web"
)

func main() {
	conf := config.LoadConfig("", "")
	webServer := web.NewWebServer(conf)
	webServer.Start()
}
