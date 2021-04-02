package loggio

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type JsonLog struct {
	Count    int32  `json:"count"`
	LogLevel string `json:"log_level"`
	Message  string `json:"message"`
}

type LogType struct {
	LogType  string `json:"string"`
	Periodms int32  `json:"period_ms"`
	Message  string `json:"message"`
}

func CreateJSONLog() {
	logCount := 0
	logObject := JsonLog{}
	logObject.LogLevel = "Information"
	logObject.Message = "Loggy is the greatest Logger out there."
	for {
		logObject.Count = int32(logCount)
		time.Sleep(time.Second * 5)
		item, err := json.Marshal(logObject)
		if err != nil {
			log.Fatalln("Fatal Error:", err)
		}
		fmt.Println(string(item))

		logCount++
	}
}

func CreateDefaultLog() {
	logCount := 0

	for {
		time.Sleep(time.Second * 5)
		fmt.Println("New Log Number: ", logCount)
		logCount++
	}
}
