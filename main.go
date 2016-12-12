package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
	if os.Getenv("INFLUXDB_DATABASE") == "" {
		println("Please set INFLUXDB_DATABASE.")
		os.Exit(1)
	}

	if os.Getenv("INFLUXDB_URL") == "" {
		println("Please set INFLUXDB_URL.")
		os.Exit(1)
	}

	if len(os.Args) <= 1 {
		println("Usage: climate_reporter <code>")
		os.Exit(1)
	}

	url := fmt.Sprintf("http://opendata.netcetera.com:80/smn/smn/%s", os.Args[1])
	record := &SmnRecord{Code: os.Args[1]}
	getJson(url, record)

	influxConfig := &InfluxdbConfig{
		Url:      os.Getenv("INFLUXDB_URL"),
		Database: os.Getenv("INFLUXDB_DATABASE"),
		Username: os.Getenv("INFLUXDB_USERNAME"),
		Password: os.Getenv("INFLUXDB_PASSWORD")}

	convertedRecord := ConvertRecord(record)
	WriteToInflux(convertedRecord, influxConfig)
}
