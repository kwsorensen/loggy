package main

import (
	"encoding/json"
	"fmt"
	"log"
	loggio "loggio"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
)

var Logs []Log
var Config InitConfig

type Log struct {
	Type string    `json:"Type"`
	Id   uuid.UUID `json:"id"`
}

type InitConfig struct {
	JSON_INTERVAL_MS, DEFAULT_INTERVAL_MS int
}

func startupConfig() {
	fmt.Println("Starting Configuration")
	ENV_JSON_INTERVAL_MS, exists := os.LookupEnv("JSON_INTERVAL_MS")
	if exists {
		if value, err := strconv.Atoi(ENV_JSON_INTERVAL_MS); err == nil {
			Config.JSON_INTERVAL_MS = value
		}
	} else {
		Config.JSON_INTERVAL_MS = 1000
	}

	ENV_DEFAULT_INTERVAL_MS, exists := os.LookupEnv("DEFAULT_INTERVAL_MS")
	if exists {

		if value, err := strconv.Atoi(ENV_DEFAULT_INTERVAL_MS); err == nil {
			Config.DEFAULT_INTERVAL_MS = value
		}
	} else {
		Config.DEFAULT_INTERVAL_MS = 1000
	}

	go loggio.CreateDefaultLog(Config.DEFAULT_INTERVAL_MS)
	go loggio.CreateJSONLog(Config.JSON_INTERVAL_MS)
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
	go loggio.CreateDefaultLog(1000)
}

func jsonLogHandler(w http.ResponseWriter, r *http.Request) {
	createLog("JSON")

	fmt.Fprintf(w, "JSON Log Handler Event started")
	log.Println("New JSON Log Handler Starting...")
	go loggio.CreateJSONLog(1000)
}

func main() {
	startupConfig()

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/createDefaultLog", defaultLogHandler)
	http.HandleFunc("/createJSONLog", jsonLogHandler)
	http.HandleFunc("/getLogs", returnLoggers)

	log.Println("Log Generator API Starting!")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
