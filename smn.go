package main

import (
	"strconv"
	"time"
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

type SmnRecordConverted struct {
	Code          string
	Time          time.Time
	Humidity      float64
	Precipitation float64
	Sunshine      float64
	Temperature   float64
	WindSpeed     float64
}

func convertRecord(record *SmnRecord) *SmnRecordConverted {
	recordConverted := new(SmnRecordConverted)

	recordConverted.Code = record.Code
	recordConverted.Temperature, _ = strconv.ParseFloat(record.Temperature, 32)
	recordConverted.Humidity, _ = strconv.ParseFloat(record.Humidity, 32)
	recordConverted.Sunshine, _ = strconv.ParseFloat(record.Sunshine, 32)
	recordConverted.WindSpeed, _ = strconv.ParseFloat(record.WindSpeed, 32)
	recordConverted.Time, _ = time.Parse(time.RFC3339, record.DateTime)

	return recordConverted
}
