package main

import (
	"fmt"
	"github.com/DeathVenom54/doto-backend/db"
	"github.com/DeathVenom54/doto-backend/router"
	"github.com/DeathVenom54/doto-backend/snowflake"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	logger "github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
)

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		logger.Fatalln("env PORT not found")
	}

	logger.Infof("Listening for requests on %s\n", PORT)
	err := http.ListenAndServe(PORT, router.Router)
	if err != nil {
		logger.Fatalln(err)
	}
}

func init() {
	err := setupLogger()
	if err != nil {
		log.Fatalf("Error while setting up logger:\n%s\n", err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file:\n%s\n", err)
	}

	snowflake.StartSnowflakeNode()

	err = db.Connect()
	if err != nil {
		logger.Fatalf("Error while connecting to db:\n%s\n", err)
	}
}

func setupLogger() error {
	environment := os.Getenv("ENVIRONMENT")

	if environment == "PROD" {
		fmt.Println("Environment: Production")
		logger.SetFormatter(&logger.JSONFormatter{
			PrettyPrint: true,
		})
		logger.SetLevel(logger.InfoLevel)
	} else {
		fmt.Println("Environment: Development")
		logger.SetFormatter(&logger.TextFormatter{
			FullTimestamp: true,
		})
		logger.SetLevel(logger.DebugLevel)
	}

	return nil
}
