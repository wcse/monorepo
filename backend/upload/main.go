package main

import (
	"github.com/wcse/monorepo/backend/shared/config"
	"github.com/wcse/monorepo/backend/upload/server"
)

func main() {
	flags := config.ParseFlags()
	server.Run(flags)
}
