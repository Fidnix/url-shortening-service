package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB
var once sync.Once

func InitDB() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the file .env")
	}

	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%v:%v", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}

	fmt.Println(cfg.FormatDSN())
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("Couldn't config database")
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal("Couldn't connect database")
	}
}

func GetDB() *sql.DB {
	once.Do(InitDB)
	return db
}
