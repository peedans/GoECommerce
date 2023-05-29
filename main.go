package main

import (
	"github.com/peedans/GoEcommerce/config"
	"github.com/peedans/GoEcommerce/modules/servers"
	"github.com/peedans/GoEcommerce/pkg/databases"
	"os"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.LoadConfig(envPath())

	db := databases.DbConnect(cfg.Db())
	defer db.Close()

	servers.NewServer(cfg, db).Start()
}
