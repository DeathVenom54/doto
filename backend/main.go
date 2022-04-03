package main

import (
	"github.com/DeathVenom54/doto-backend/router"
	"github.com/joho/godotenv"
	logger "github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file:\n%s\n", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("Error while setting up logger:\n%s\n", err)
	}

	logger.Info("Listening for requests on :3000")
	err = http.ListenAndServe(":3000", router.Router)
	if err != nil {
		logger.Fatalln(err)
	}
}

func setupLogger() error {
	environment := os.Getenv("ENVIRONMENT")

	if environment == "PROD" {
		logger.SetFormatter(&logger.JSONFormatter{
			PrettyPrint: true,
		})
	} else {
		logger.SetFormatter(&logger.TextFormatter{
			FullTimestamp: true,
		})
	}

	return nil
}
