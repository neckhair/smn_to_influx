package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const (
	MyDB     = "smn"
	username = "smn"
	password = "smn"
)

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {
	if len(os.Args) <= 1 {
		println("Usage: climate_reporter <code>")
		os.Exit(1)
	}

	url := fmt.Sprintf("http://opendata.netcetera.com:80/smn/smn/%s", os.Args[1])
	record := &SmnRecord{Code: os.Args[1]}
	getJson(url, record)

	convertedRecord := convertRecord(record)
	writeToInflux(convertedRecord)
}
