package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/Ashmita152/testInternBoilerPlate/checker"
	"github.com/Ashmita152/testInternBoilerPlate/server"
)

type Config struct {
	listenUrl string
}

func main() {
	config := Config{}
	app := kingpin.New("status-check", "A server to track the health of logz.io region endpoints")
	app.Flag("listen.url", "The URL at which this server listens").Default("localhost:8080").StringVar(&config.listenUrl)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	endpoints := []string{
		"app.logz.io",
		"app-au.logz.io",
		"app-ca.logz.io",
		"app-eu.logz.io",
		"app-nl.logz.io",
		"app-uk.logz.io",
		"app-wa.logz.io",
	}

	// background go-routine for polling the health of region endpoints
	go checker.PollEndpoints(endpoints)

	handler := http.NewServeMux()
	handler.Handle("/status", server.NewStatusHandler())

	server := http.Server{
		Handler: handler,
	}

	listener, err := net.Listen("tcp", config.listenUrl)
	if err != nil {
		log.Fatalf("Error while listening on %s: %s\n", config.listenUrl, err)
	} else {
		log.Printf("Listening on %s\n", config.listenUrl)
	}
	defer listener.Close()

	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("Error while serving request: %s\n", err)
	}
}
