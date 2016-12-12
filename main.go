package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type SmnStation struct {
	Code      string
	Name      string
	Ch1903Y   int
	Ch1903X   int
	Lat       int
	Lng       int
	Elevation int
}

type SmnRecord struct {
	Station       SmnStation
	Code          string
	DateTime      string
	Temperature   string
	Sunshine      string
	Precipitation string
	WindDirection string
	WindSpeed     string
	QnhPressure   string
	GustPeak      string
	Humidity      string
	QfePressure   string
	QffPressure   string
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {
	station := os.Args[1]
	url := fmt.Sprintf("http://opendata.netcetera.com:80/smn/smn/%s", station)

	record := new(SmnRecord)
	getJson(url, record)
	println(record.Temperature)
}
