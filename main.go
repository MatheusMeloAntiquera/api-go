package main

import (
	"log"

	"github.com/MatheusMeloAntiquera/api-go/src/database"
	"github.com/MatheusMeloAntiquera/api-go/src/helpers"
	Router "github.com/MatheusMeloAntiquera/api-go/src/routes"
)

type User struct {
	ID   uint
	Name string
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	helpers.LoadDotEnv()
	database.ConnectToDatabase()
	Router.Run()
}
