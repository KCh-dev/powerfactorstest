package goapp

import (
	"fmt"
	"goapp/internal/config"
	"goapp/internal/pkg/httpsrv"
	"goapp/internal/pkg/strgen"
	"log"
	"os"
)

func Start(cfg *config.Config, exitChannel chan os.Signal) error {
	var (
		strChan = make(chan string, 100)    // String channel with max parallel counter processes.
		strCli  = strgen.New(strChan)       // String generator.
		httpSrv = httpsrv.New(strChan, cfg) // HTTP server, now passing config
	)

	// Start String Generator.
	if err := strCli.Start(); err != nil {
		return fmt.Errorf("failed to start string generator: %w", err)
	}
	defer strCli.Stop()

	// Start HTTP server.
	if err := httpSrv.Start(); err != nil {
		return fmt.Errorf("failed to start HTTP server: %w", err)
	}
	defer httpSrv.Stop()

	log.Println("GoApp Started")
	defer log.Println("GoApp Stopped")

	<-exitChannel

	return nil
}
