package main

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"
)

type jsonLog struct {
	Count   int32  `json:"count"`
	Message string `json:"message"`
}

func startLogging(wg *sync.WaitGroup) {
	counter := 0
	for {
		time.Sleep(time.Second * 5)
		log.Println("Logging Line: ", counter, "Line")
		counter++
	}
	wg.Done()
}

func startLoggingJSON(wg *sync.WaitGroup) {
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
		log.Println(string(item))

		counter++
	}
	wg.Done()
}

func main() {
	// get file path as first CLI argument
	logpath := "./logs.txt"
	if len(os.Args) > 1 {
		logpath = os.Args[1]
	}

	log.Println("Logging to:", logpath)

	// If the file doesn't exist, create it or append to the file
	/*file, err := os.OpenFile(logpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	*/
	log.SetFlags(0)
	// log.SetOutput(file)
	log.Println("Log Generator Starting!")

	// set up wait group so app will block for goroutines
	var wg sync.WaitGroup
	wg.Add(2)
	go startLogging(&wg)
	go startLoggingJSON(&wg)
	wg.Wait()
}
