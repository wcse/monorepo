package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/wcse/monorepo/backend/feed/server"
	"github.com/wcse/monorepo/backend/shared/config"
)

func main() {
	flags := config.ParseFlags()
	server.Run(flags)
}
