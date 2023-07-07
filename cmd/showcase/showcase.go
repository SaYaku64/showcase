package main

import (
	"github.com/SaYaku64/showcase/internal/config"
)

func main() {
	cfg := config.Create()

	cfg.SetLogger()
	cfg.SetBroker()

	
}
