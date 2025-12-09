package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	
	_ "github.com/lib/pq"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() *Config {
	port, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))
	
	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     port,
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "orderdb"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func ConnectDB(cfg *Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	
	log.Println("Successfully connected to order database")
	return db, nil
}