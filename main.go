package main

import (
	Connection "./database"
	"./model"
	"./repository"
	"./configs"
)

func main(){
	db := Connection.GetDB()

	db.AutoMigrate(&model.Contact{})

	defer db.Close()

	contactRepository := repository.NewContactRepository(db)

	route := configs.SetupRoutes(contactRepository)

	route.Run(":8000")
}