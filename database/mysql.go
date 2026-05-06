package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"example/rest-api/config"
)

var DB *sql.DB

func Connect() {
	cfg := config.AppConfig

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	fmt.Println("DSN:", dsn) // Debug: Print the DSN

	db, err := sql.Open(cfg.Database.Driver, dsn)
	if err != nil {
		log.Fatal("DB open error:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("DB ping error:", err)
	}

	DB = db
	log.Println("✅ MySQL Connected")
}
