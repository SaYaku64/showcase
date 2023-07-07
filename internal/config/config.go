package config

import (
	"os"

	"github.com/SaYaku64/showcase/internal/broker"
	"github.com/SaYaku64/showcase/internal/broker/gin"
	"github.com/SaYaku64/showcase/internal/logger"
	"github.com/SaYaku64/showcase/internal/logger/golog"
	"github.com/SaYaku64/showcase/internal/logger/otel"
)

type Config struct {
	Logger logger.ILogger
	Broker broker.IBroker
}

// Create - creates global config with settings
func Create() *Config {
	//
	//
	// there are can be uploading info from .env file or using terraform tfvars etc
	//
	//

	config := Config{}

	return &config
}

func (cfg *Config) SetLogger() {
	switch os.Getenv(cfgLogs) {
	case "otel":
		cfg.Logger = otel.Init()
	default:
		fallthrough
	case "golog":
		cfg.Logger = golog.Init()
	}
}

func (cfg *Config) SetBroker() {
	switch os.Getenv(cfgBroker) {
	case "gin":
		cfg.Broker = gin.New()
	case "nats":
		// 
	case "ws":
		// 
	default:
		cfg.Logger.Fatal("UNKNOWN BROKER")
	}
}
