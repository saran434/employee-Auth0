package main

import (
	"log"
	"net"
	"net/http"
	"strconv"

	"employeejwt/internal/config"
	setupmod "employeejwt/internal/setup.go"
)

func main() {
	serverConfig := config.LoadServer()
	listenAddress := net.JoinHostPort(serverConfig.Host, strconv.Itoa(serverConfig.Port))
	serverURL := serverConfig.Protocol + "://" + listenAddress
	log.Println("Server Protocol:", serverConfig.Protocol)
	log.Println("Server Host:", serverConfig.Host)
	log.Println("Server Port:", serverConfig.Port)
	log.Println("Server URL:", serverURL)
	r, err := setupmod.SetupRouter()
	if err != nil {
		log.Fatal("configure authentication: ", err)
	}
	setupmod.SetupDatabase()
	log.Fatal(http.ListenAndServe(listenAddress, r))
}
