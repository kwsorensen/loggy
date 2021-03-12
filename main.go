package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type jsonLog struct {
	Count   int32  `json:"count"`
	Message string `json:"message"`
}

func startLogging() {
	counter := 0
	for {
		time.Sleep(time.Second * 5)
		log.Println("Logging Line: ", counter, "Line")
		counter++
	}
}

func startLoggingJSON() {
	counter := 0
	logObject := jsonLog{}
	logObject.Message = "Hello string"
	for {
		logObject.Count = int32(counter)
		time.Sleep(time.Second * 5)
		item, err := json.Marshal(logObject)
		if err != nil {
			log.Fatalln("Fatal Error:", err)
		}

		fmt.Println(string(item))
		log.Println(string(item))

		counter++
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	// If the file doesn't exist, create it or append to the file
	os.MkdirAll("/mnt/c/logs/app2/logs", 0666)
	file, err := os.OpenFile("/mnt/c/json/mixed-objects.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetFlags(0)
	log.SetOutput(file)
	log.Println("Log Generator Starting!")

	go startLogging()
	go startLoggingJSON()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
