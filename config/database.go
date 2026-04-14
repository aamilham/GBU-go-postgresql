package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Driver PostgreSQL
)

var DB *sql.DB

func ConnectDB() {
	//load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal meload file .envL: ", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	//set postgresql connection url
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", 
		user, password, host, port, dbname)
	
	//koneksi ke postgresql
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Gagal membuka koneksi ke database:", err)
	}

	//test koneksi dengan ping ke database
	err = DB.Ping()
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	fmt.Println("Berhasil terhubung ke PostgreSQL menggunakan .env!")
}