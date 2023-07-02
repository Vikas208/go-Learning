package loaders

import (
	"log"
	"os"
)

func LogFileSetup() *os.File {
	err := os.MkdirAll(os.Getenv("LOGS_DIR"), 0755)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Directory created: ", os.Getenv("LOGS_DIR"))
	filepath := os.Getenv("LOGS_DIR") + "/log.txt"

	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("File created: ", filepath)

	log.SetOutput(file)
	return file
}
