package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"fmt"
)

var Db sql.DB

func init() {
	userName := os.Getenv("DB_USERNAME")

	connString := fmt.Sprintf("%s:@(%s:%s)/diploradar?parseTime=true", userName, os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

	db, err := sql.Open(os.Getenv("DB_CONNECTION"), connString)

	Db = *db

	if err != nil {
		fmt.Println(err.Error())
	}

	err = Db.Ping()

	if err != nil {
		fmt.Println(err.Error())
	}
}