// config/db.go
package config

import (
	"log"
	"os"
)

func GetDBConnString() string {
	conn := os.Getenv("DATABASE_URL")
	if conn == "" {
		log.Fatal("DATABASE_URL not set")
	}
	return conn
}

