package core

import (
	"os"
)

// Config struct
type Config struct {
	APIAddr             string
	APIPrefix           string
	SuperheroAPIBaseURL string
	SuperheroAPIToken   string
	SuperheroAPITimeout int
}

// GetConfig instance
func GetConfig() *Config {
	c := &Config{}
	c.Load()
	return c
}

// Load data from env variables
func (c *Config) Load() {
	c.APIAddr = os.Getenv("API_ADDR")
	c.APIPrefix = os.Getenv("API_PREFIX")
	c.SuperheroAPIBaseURL = os.Getenv("SUPERHEROAPI_URL")
	c.SuperheroAPIToken = os.Getenv("SUPERHEROAPI_TOKEN")
	c.SuperheroAPITimeout = GetUtils().ParseInt(os.Getenv("SUPERHEROAPI_TIMEOUT"))
}
