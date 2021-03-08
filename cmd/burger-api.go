package main

import (
	"burger-api/internal/config"
	"burger-api/internal/db"
	"burger-api/internal/db/migration"
	"burger-api/internal/routes"
	"net/http"
	"time"
)

func main() {
	cfgFlags, err := config.ParseFlags()
	if err != nil {
		config.Logger.Fatal("Config Parse error:", err)
	}

	cfg, err := config.NewConfig(cfgFlags.Path)
	if err != nil {
		config.Logger.Fatal("Config Parse error:", err)
	}
	run(*cfg, cfgFlags)
}

func run(cfg config.Config, flags config.Flags) {
	db.InitDb(cfg)
	if flags.Migrate {
		migration.CreateCollections()
	}

	api := http.Server{
		Addr:         cfg.Web.Host + ":" + cfg.Web.Port,
		Handler:      routes.CreateHandlers(),
		ReadTimeout:  cfg.Web.Timeout.Read * time.Second,
		WriteTimeout: cfg.Web.Timeout.Write * time.Second,
	}

	if err := api.ListenAndServe(); err != nil {
		config.Logger.Println("ERROR", err)
	}
}
