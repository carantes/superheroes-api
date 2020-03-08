package main

import (
	"log"

	"github.com/carantes/superheroes-api/bundles/superheroesbundle"
	"github.com/carantes/superheroes-api/core"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {
	// Load envs
	cfg := loadConfig()

	// Init DB
	db, err := initDB(cfg.DBType, cfg.DBConnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Init Bundles
	bundles := initBundles(db)

	// Configure Server with prefix and routes
	srv := core.NewServer(bundles, core.ServerOpts{APIPrefix: cfg.APIPrefix})

	// Start
	log.Fatal(srv.Start(cfg.APIAddr))
}

func initBundles(db *gorm.DB) []core.Bundle {
	log.Println("Loading bundles")
	return []core.Bundle{
		superheroesbundle.NewSuperheroesBundle(db),
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

func initDB(dbType string, dbConn string) (*gorm.DB, error) {
	db, err := gorm.Open(dbType, dbConn)
	if err != nil {
		return &gorm.DB{}, err
	}

	db.AutoMigrate(&superheroesbundle.Superhero{})
	return db, nil
}
