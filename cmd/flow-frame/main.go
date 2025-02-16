package main

import (
	"flow-frame/internal/auth"
	myCli "flow-frame/internal/cli"
	"flow-frame/internal/config"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	authURL := server.GetAuthURL(cfg.ClientID, cfg.ClientSecret)

	err = exec.Command("open", authURL).Start()
	if err != nil {
		log.Fatal(err)
	}

	authChan := make(chan string)
	go server.StartAuthServer(cfg.Port, authChan)
	code := <-authChan
	fmt.Println("Received OAuth code", code)

	app := myCli.CreateCLIApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
