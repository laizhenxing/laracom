package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateDBConnection() (*gorm.DB, error) {
	dsn := os.Getenv("DB_DSN")

	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	return gorm.Open("mysql", fmt.Sprintf(dsn, user, password, host, port, dbName))
}
