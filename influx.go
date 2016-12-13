package main

import (
	"github.com/influxdata/influxdb/client/v2"
	"log"
)

type InfluxdbConfig struct {
	Username string
	Password string
	Url      string
	Database string
}

func WriteToInflux(record *SmnRecordConverted, config *InfluxdbConfig) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     config.Url,
		Username: config.Username,
		Password: config.Password,
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  config.Database,
		Precision: "m",
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	// Create a point and add to batch
	tags := map[string]string{"station": record.Code}
	fields := map[string]interface{}{
		"temperature":   record.Temperature,
		"humidity":      record.Humidity,
		"windspeed":     record.WindSpeed,
		"sunshine":      record.Sunshine,
		"precipitation": record.Precipitation,
		"gustpeak":      record.GustPeak,
	}

	pt, err := client.NewPoint("smn", tags, fields, record.Time)

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	bp.AddPoint(pt)

	// Write the batch
	err = c.Write(bp)

	if err != nil {
		log.Fatalln("Error: ", err)
	}
}
