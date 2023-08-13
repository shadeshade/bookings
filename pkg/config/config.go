package config

import (
	"log"
)

type AppConfig struct {
	InfoLog      *log.Logger
	InProduction bool
}
