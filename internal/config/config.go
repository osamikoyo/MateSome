package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port     int
	Hostname string
}

func Get() Config {
	ports := os.Getenv("PORT")
	port, _ := strconv.Atoi(ports)
	host := os.Getenv("HOSTNAME")
	return Config{
		Port:     port,
		Hostname: host,
	}
}
