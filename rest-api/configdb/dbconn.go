package configdb

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() mysql.Config {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "rootroot",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "Employee",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	var err error
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return cfg

}
