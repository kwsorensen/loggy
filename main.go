package main

import (
	"log"
	"os"
)

func setupLogger() {

}

func main() {
	// If the file doesn't exist, create it or append to the file
	os.MkdirAll("/mnt/c/logs/app1/", 0666)
	file, err := os.OpenFile("/mnt/c/logs/app1/logs/logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.Println("Hello world!")
}
