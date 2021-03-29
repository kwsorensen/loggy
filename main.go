package main

import (
	"encoding/json"
	"fmt"
	"log"
	loggio "loggio"
	"net/http"

	"github.com/google/uuid"
)

var Logs []Log

type Log struct {
	Type string    `json:"Type"`
	Id   uuid.UUID `json:"id"`
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to loggy!\n")
}

func createLog(logType string) {
	var logStruct Log
	logUUID, err := uuid.NewUUID() // generate new UUID for the log
	if err != nil {
		log.Fatalln("Fatal Error Creating UUID:", err)
	}
	logStruct.Id = logUUID
	logStruct.Type = logType

	Logs = append(Logs, logStruct)

}

func returnLoggers(w http.ResponseWriter, r *http.Request) {

	data, err := json.Marshal(Logs)
	if err != nil {
		log.Fatal("Error Returning Logs: ", err)
	}
	fmt.Fprintln(w, string(data))

}

func defaultLogHandler(w http.ResponseWriter, r *http.Request) {
	createLog("Default")

	fmt.Fprintf(w, "Default Log Handler Event started")
	log.Println("New Default Log Handler Starting...")
	go loggio.CreateDefaultLog()
}

func jsonLogHandler(w http.ResponseWriter, r *http.Request) {
	createLog("JSON")

	fmt.Fprintf(w, "JSON Log Handler Event started")
	log.Println("New JSON Log Handler Starting...")
	go loggio.CreateJSONLog()
}

func main() {

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/createDefaultLog", defaultLogHandler)
	http.HandleFunc("/createJSONLog", jsonLogHandler)
	http.HandleFunc("/getLogs", returnLoggers)

	log.Println("Log Generator Starting!")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
