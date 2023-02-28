package main

import (
	"sheet-retrieve/config"
	"sheet-retrieve/internal/adapters/meld"
	"sheet-retrieve/utils"
)

func main() {
	config := config.LoadConfig("", "")
	logger := utils.SetupLogger(config.Logger.LogLevel, config.Logger.LogPath+config.UserList.LogFile)
	meldServer := meld.NewMeldServer(config, logger)
	meldServer.Start()
}
