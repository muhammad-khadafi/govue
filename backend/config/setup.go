package config

import (
	_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

var DB *sqlx.DB

func ConnectDatabase() (*sqlx.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dataSourceName := os.Getenv("DB_USERNAME") + "/" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("SERVICE_NAME")
	db, err := sqlx.Connect(os.Getenv("DB_DRIVER"), dataSourceName)
	DB = db

	DB_MAX_OPEN_CONN, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONN"))
	DB_MAX_IDLE_CONN, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONN"))
	DB_CONN_MAX_LIFETIME, _ := strconv.Atoi(os.Getenv("DB_CONN_MAX_LIFETIME"))
	DB_IDLE_MAX_LIFETIME, _ := strconv.Atoi(os.Getenv("DB_IDLE_MAX_LIFETIME"))
	//use nano second, 1 sec = 1000000000 ns
	DB.SetMaxOpenConns(DB_MAX_OPEN_CONN)
	DB.SetConnMaxLifetime(time.Minute * time.Duration(DB_CONN_MAX_LIFETIME))
	DB.SetMaxIdleConns(DB_MAX_IDLE_CONN)
	DB.SetConnMaxIdleTime(time.Minute * time.Duration(DB_IDLE_MAX_LIFETIME))
	return DB, err
}
