package familytree

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func Connection() *sql.Conn {
	Conf := mysql.Config{
		User:                 "root",
		Passwd:               "",
		DBName:               "family_tree",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", Conf.FormatDSN())
	if err != nil {
		fmt.Printf("Error %s", err)
	}
	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Error happened during pinging")
	}
	// fmt.Println("Connected!")

	conn, err := db.Conn(context.Background())
	if err != nil {
		fmt.Println("Error happened during connection")
	}
	return conn
}
