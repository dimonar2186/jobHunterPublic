package dbcrud

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
)

var db *sql.DB

// ConnectDB activates a connection to DB
func ConnectDB() (*sql.DB, error) {
	//dbUser := os.Getenv("DB_USER")
	//dbPassword := os.Getenv("DB_PWD")
	//dbName := os.Getenv("DB_NAME")

	dbUser := "jobHunterBasic_su"
	dbPassword := "A211186dssu"
	dbName := "jobHunterBasic"

	//if dbUser == "" || dbPassword == "" || dbName == "" {
	//	return nil, errors.New("database credentials are not set")
	//}
	if dbUser == "" {
		return nil, errors.New("database cred1 are not set")
	}
	if dbPassword == "" {
		return nil, errors.New("database cred2 are not set")
	}
	if dbName == "" {
		return nil, errors.New("database cred3 are not set")
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, dbUser, dbPassword, dbName)
	//Open connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	//Check connection by db.ping
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return db, nil
}

// GetDB returns active DB-connection
func GetDB() *sql.DB {
	if db != nil {
		return db
	}
	return nil
}

// CloseDB closes active DB-connection
func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

func OpenTransaction() (*sql.Tx, error) {
	db, err := ConnectDB() //connect to the database
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB %w", err)
	}
	defer db.Close()

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil) //transaction start
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction %w", err)
	}
	return tx, nil
}
