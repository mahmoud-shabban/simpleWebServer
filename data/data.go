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

// create user adm@'%' identified by 'Pass@123#';

func init() {

	connStr := "admin:pass@tcp(localhost)/"
	db, err := sql.Open("mysql", connStr)
	CheckError(err)

	CheckError(db.Ping())
	log.Print("Connection To DB Succeeded!")

	// time.Sleep(10 * time.Second)
	// q := `
	// 	CREATE DATABASE IF NOT EXISTS api;
	// 	USE api;
	// 	CREATE TABLE IF NOT EXISTS users (
	// 		id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	// 		username VARCHAR(255) NOT NULL,
	// 		email VARCHAR(255) NOT NULL,
	// 		isActive CHAR(1) NOT NULL DEFAULT 'Y'
	// 		);
	// 	`

	q := `
	USE api;
	CREATE TABLE IF NOT EXISTS staff (
	id     int NOT NULL AUTO_INCREMENT PRIMARY KEY,
	fname  varchar(255) NOT NULL,
	lname  varchar(255) NOT NULL,
	title  varchar(255) NOT NULL,
	isActive CHAR(1) NOT NULL DEFAULT 'Y'

);

	`
	_, err = db.Exec(q)
	CheckError(err)

}
