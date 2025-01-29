package data

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// initialize database

func init() {

	connStr := "admin:pass@tcp(localhost)/"
	db, err := sql.Open("mysql", connStr)
	CheckError(err)

	CheckError(db.Ping())
	log.Print("Connection To DB Succeeded!")

	_, _ = db.Exec("CREATE DATABASE IF NOT EXISTS api")
	_, _ = db.Exec("USE api")

	q := `
		CREATE TABLE IF NOT EXISTS users (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			isActive CHAR(1) NOT NULL DEFAULT 'Y'
			);
		`
	_, err = db.Exec(q)
	CheckError(err)

}
