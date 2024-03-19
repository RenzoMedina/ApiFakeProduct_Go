package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

const (
	dns = "usi1p2fhaqjrbcxd:7b3gioYphTLYaZoRbF8P@tcp(byr2cqsjyxldxj3oabah-mysql.services.clever-cloud.com)/byr2cqsjyxldxj3oabah?parseTime=true"
)

func NewConnection() {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", dns)
		if err != nil {
			log.Fatalf("Can't open database %v", err)
		}
		if err = db.Ping(); err != nil {
			log.Fatalf("Can't do ping %v", err)
		}
		fmt.Println("Connection successfully!!!")
	})
}

func Pool() *sql.DB {
	return db
}

/*
? Function for validated the null value
*/
/*
func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}*/
