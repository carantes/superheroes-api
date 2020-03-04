package core

import "os"

// Config struct
type Config struct {
	APIAddr   string
	APIPrefix string
}

// Load data from env variables
func (c *Config) Load() {
	c.APIAddr = os.Getenv("API_ADDR")
	c.APIPrefix = os.Getenv("API_PREFIX")
}
