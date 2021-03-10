package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func startLogging() {
	counter := 0
	for {
		time.Sleep(time.Second * 5)
		log.Println("Logging Line: ", counter, "Line")
		counter++
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	// If the file doesn't exist, create it or append to the file
	os.MkdirAll("/mnt/c/logs/app2/logs", 0666)
	file, err := os.OpenFile("/mnt/c/logs/app2/logs/logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.Println("Log Generator Starting!")

	go startLogging()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
