package main

import (
	"github.com/NYTimes/gziphandler"
	config "github.com/calebtracey/config-yaml"
	"github.com/calebtracey/rugby-crawler-api/internal/routes"
	"github.com/calebtracey/rugby-crawler-api/internal/routes/openapi"
	log "github.com/sirupsen/logrus"
)

var (
	configPath = "local_config.yaml"
)

const Port = "6080"

func main() {
	defer panicQuit()

	appConfig := config.New(configPath)
	facade, err := initializeDAO(*appConfig)

	if err != nil {
		log.Error(err)
		panicQuit()
	}
	handler := routes.Handler{Service: facade}
	router := handler.InitializeRoutes()

	openapi.RegisterOpenAPI(router)
	c := CorsHandler()

	log.Fatal(ListenAndServe(Port, gziphandler.GzipHandler(c.Handler(router))))
}

func panicQuit() {
	if r := recover(); r != nil {
		log.Errorf("I panicked and am quitting: %v", r)
		log.Error("I should be alerting someone...")
	}
}
