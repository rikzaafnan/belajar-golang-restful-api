package app

import (
	"belajar-golang-restful-api/helper"
	"database/sql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_restful_api")
	// db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")
	helper.PanicIfError(err)

	// db.SetMaxIdleConns(5)
	// db.SetMaxOpenConns(20)
	// db.SetConnMaxLifetime(60 * time.Minute)
	// db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
