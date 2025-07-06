package test

import (
	"database/sql"
	"hermawansafrin/belajar-golang-restful-api/helper"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func SetupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "kledo:someRandomPassword@tcp(localhost:3306)/belajar_golang_restful_api_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
