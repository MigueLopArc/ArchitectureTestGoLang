package data

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

var (
	data *DatabaseContext
	once sync.Once
)

// DatabaseContext manages the connection to the database.
type DatabaseContext struct {
	DB *sql.DB
}

// New returns a new instance of Data with the database connection ready.
func New() *DatabaseContext {
	once.Do(initDB)

	return data
}

func getConnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return sql.Open("postgres", psqlInfo)
}

// initialize the data variable with the connection to the database.
func initDB() {
	db, err := getConnection()
	if err != nil {
		log.Panic(err)
	}

	data = &DatabaseContext{
		DB: db,
	}
}

// Close closes the resources used by data.
func Close() error {
	if data == nil {
		return nil
	}

	return data.DB.Close()
}
