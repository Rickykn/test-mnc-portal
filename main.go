package main

import (
	"fmt"
	"log"

	"github.com/Rickykn/buddyku-app.git/routers"
	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("Buddyku-app")

	errEnv := godotenv.Load()

	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	server := routers.Server()

	err := server.Run()
	if err != nil {
		panic(err)
	}
}
