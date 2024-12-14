package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func NewDatabase(cfg Config) (*sql.DB, error) {
	//Build connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName)

	//Open connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	//Set up connection pool settings
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 5)

	//Test connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Printf("Connected to database %s on %s:%d", cfg.DbName, cfg.Host, cfg.Port)
	return db, nil

}
