package store

import (
	"database/sql"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetConnection() *sql.DB {
	return db
}

func NewMySQLConnection() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	protocol := os.Getenv("DB_PROTOCOL")
	once.Do(func() {
		var err error
		db, err = sql.Open(
			"mysql",
			user+":"+password+"@"+protocol+"("+host+":"+port+")/"+database+"?parseTime=true")
		if err != nil {
			log.Fatalf("Error opening database: %v", err)
		}
		if err := db.Ping(); err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}
		log.Println("Connected to MySQL database successfully.")
	})
}
