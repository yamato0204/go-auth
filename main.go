package main

import (
	"log"

	"github.com/yamato0204/go-up/app/db"
	"github.com/yamato0204/go-up/app/infra"
	"github.com/yamato0204/go-up/app/router"
	"github.com/yamato0204/go-up/app/usecase"
)

func main() {

	db := db.NewDB()
	infra := infra.NewInfra(db)
	usecase := usecase.NewUsecase(infra)

	e := router.GetRouter(usecase)

	err := e.Start(":8080")
	if err != nil{
		log.Fatal(err)
	}
}
