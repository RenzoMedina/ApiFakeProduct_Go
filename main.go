package main

import (
	"fakeproduct/model"
	"fakeproduct/storage"
	"log"
)

func main() {

	storage.NewConnection()
	mysql := storage.NewMySql(storage.Pool())
	newmysql := model.NewServices(mysql)

	if err := newmysql.Migrate(); err != nil {
		log.Fatalf("model.Migrate() %v", err)

	}
}
