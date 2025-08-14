package db

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	conn *sql.DB
	once   sync.Once
)

func DbContion() (*sql.DB, error) {
	var err error
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	once.Do(func() {
		conn, err = sql.Open("mysql", dns)
		if err != nil {
			fmt.Println("we get error while fetching db connection")
			panic(err)

		}
		createTable := `
            CREATE TABLE IF NOT EXISTS numbers (
                id INT AUTO_INCREMENT PRIMARY KEY,
                num VARCHAR(255) NOT NULL UNIQUE
            );
        `
        _, err = conn.Exec(createTable)
        if err != nil {
            fmt.Println("error creating table:", err)
            panic(err)
        }
		
		
		
	})

	return conn, err
}
