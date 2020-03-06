package main

import (
	"log"

	"github.com/carantes/superheroes-api/bundles/superheroesbundle"
	"github.com/carantes/superheroes-api/core"
	"github.com/joho/godotenv"
)

func main() {
	// Load envs
	cfg := loadConfig()

	// Init Bundles
	bundles := initBundles()

	// Start Server with prefix and bundle routes
	srv := core.NewServer(bundles, core.ServerOpts{APIPrefix: cfg.APIPrefix})
	log.Fatal(srv.Start(cfg.APIAddr))
}

func initBundles() []core.Bundle {
	log.Println("Loading bundles")
	return []core.Bundle{
		superheroesbundle.NewSuperheroesBundle(),
	}
}

func loadConfig() *core.Config {
	log.Println("Loading configs from .env")
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file not found")
	}

	return core.GetConfig()
}
