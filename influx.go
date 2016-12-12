package main

import (
	"github.com/influxdata/influxdb/client/v2"
	"log"
)

func writeToInflux(record *SmnRecordConverted) {
	// Make client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
		// Username: username,
		// Password: password,
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "m",
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	// Create a point and add to batch
	tags := map[string]string{"station": record.Code}
	fields := map[string]interface{}{
		"temperature": record.Temperature,
		"humidity":    record.Humidity,
		"windspeed":   record.WindSpeed,
		"sunshine":    record.Sunshine,
	}

	pt, err := client.NewPoint("climate", tags, fields, record.Time)

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
