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
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")
	once.Do(func() {
		var err error
		db, err = sql.Open(
			"mysql",
			user+":"+password+"@tcp(localhost:3306)/"+database+"?parseTime=true")
		if err != nil {
			log.Fatalf("Error opening database: %v", err)
		}
		if err := db.Ping(); err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}
		log.Println("Connected to MySQL database successfully.")
	})
}
