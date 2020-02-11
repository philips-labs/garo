package main

import (
	"encoding/json"
	"log"
)

func dieOnErr(format string, err error) {
	if err != nil {
		log.Fatalf(format, err)
	}
}

func printJSON(format string, data interface{}) {
	d, err := json.MarshalIndent(data, "", "  ")
	dieOnErr("Failed json marshalling: %s", err)
	log.Printf(format, d)
}
