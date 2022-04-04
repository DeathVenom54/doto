package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	logger "github.com/sirupsen/logrus"
	"os"
)

var DB *sqlx.DB

func init() {
	var err error
	DB, err = Connect()
	if err != nil {
		logger.Fatalf("Error while connecting to db:\n%s\n", err)
	}
}

func Connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", os.Getenv("DB_CONN_STRING"))
	if err != nil {
		return nil, err
	}

	db.MustExec(schema)
	logger.Infoln("Connected to database")
	return db, err
}
