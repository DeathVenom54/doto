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

	PORT := os.Getenv("PORT")
	if PORT == "" {
		logger.Fatalln("env PORT not found")
	}

	logger.Infof("Listening for requests on %s\n", PORT)
	err = http.ListenAndServe(PORT, router.Router)
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
