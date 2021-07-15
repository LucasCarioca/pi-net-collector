package main

import (
	"flag"
	"fmt"
	"github.com/LucasCarioca/pi-net-collector/pkg/config"
	"github.com/LucasCarioca/pi-net-collector/pkg/server"
	"os"
)

func getEnv() string {
	env := os.Getenv("ENV")
	if env == "" {
		envFlag := flag.String("e", "dev", "")
		flag.Usage = func() {
			fmt.Println("Usage: server -e {mode}")
			os.Exit(1)
		}
		flag.Parse()
		env = *envFlag
	}
	return env
}

func main() {
	config.Init(getEnv())
	server.Init(config.GetConfig())
}
