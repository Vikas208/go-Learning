package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Vikas208/social-media-api/loaders"
	"github.com/Vikas208/social-media-api/routers"

	"github.com/joho/godotenv"
)

func serverConnect() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from: ", err)
		}
	}()
	fmt.Printf("🚀 SERVER LISTENING AT ADDR %s", os.Getenv("PORT"))
	fmt.Println()
	err := http.ListenAndServe(":"+os.Getenv("PORT"), routers.Routers())
	if err != nil {
		log.Println("Error server not starting", err)
	}
}

func loadDependency() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Recovered from: ", err)
		}
	}()
	file := loaders.LogFileSetup()
	loaders.ConnectToDatabase()
	defer file.Close()
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	loadDependency()
	serverConnect()
}
