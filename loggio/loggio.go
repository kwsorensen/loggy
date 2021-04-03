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

func CreateJSONLog(interval int) {
	logCount := 0
	logObject := JsonLog{}
	logObject.LogLevel = "Information"
	logObject.Message = "Loggy is the greatest Logger out there."
	for {
		logObject.Count = int32(logCount)
		time.Sleep(time.Millisecond * time.Duration(interval))
		item, err := json.Marshal(logObject)
		if err != nil {
			log.Fatalln("Fatal Error:", err)
		}
		fmt.Println(string(item))

		logCount++
	}
}

func CreateDefaultLog(interval int) {
	logCount := 0

	for {
		time.Sleep(time.Millisecond * time.Duration(interval))
		fmt.Println("New Log Number: ", logCount)
		logCount++
	}
}
