// ((( WAZIUP )))
//
//
// Johann Forster, 2019
//
package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Waziup/wazigate-health/health"
)

var apiAddr string

func main() {

	// apiAddr comes from the WAZIGATE_ADDR env variable
	apiAddr = os.Getenv("WAZIGATE_ADDR")
	if apiAddr == "" {
		apiAddr = "127.0.0.1:880"
	}

	// a HTTP server for configuration
	addr := os.Getenv("WAZIAPP_ADDR")
	if addr != "" {
		server := http.HandlerFunc(Serve)
		go log.Fatal(http.ListenAndServe(addr, server))
	}

	// First, create the Health sensor at the API
	CreateHealthSensor()

	// Post some initial data and telemtry
	PostHealth(map[string]interface{}{
		"telemetry": health.Telemetry(),
		"boottime":  health.Boottime(),
	})

	counter := 0
	ticker := time.NewTicker(10 * time.Second)

	for true {

		<-ticker.C // wait

		data := map[string]interface{}{
			"cpu": health.CPU(),
		}

		if counter == 0 {
			data["disk"] = health.Disk()
		}

		PostHealth(data)
		counter++

		if counter == 30 {
			counter = 0
		}
	}
}

// CreateHealthSensor declares the "health" sensor.
// We call `POST /sensors` on the API to create the sensor.
func CreateHealthSensor() {
	sensor := map[string]string{
		"id":   "health",
		"name": "Health",
	}
	data, _ := json.Marshal(sensor)
	body := bytes.NewBuffer(data)

	log.Println("Creating Health sensor...")

	resp, err := http.Post("http://"+apiAddr+"/sensors", "application/json; charset=utf-8", body)
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}
	if resp.StatusCode == 409 {
		log.Println("The sensor was already created.")
	} else if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Fatalf("Response was not ok: %d %s", resp.StatusCode, resp.Status)
	}

	log.Printf("Successfully created sensor (Status: %s).", resp.Status)
}

// PostHealth creates new health data.
// We call `POST /sensors/health/value` to insert a new datapoint.
func PostHealth(value interface{}) {
	data, _ := json.Marshal(value)
	body := bytes.NewBuffer(data)
	resp, err := http.Post("http://"+apiAddr+"/sensors/health/value", "application/json; charset=utf-8", body)
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Fatalf("Response was not ok: %d %s", resp.StatusCode, resp.Status)
	}
	log.Println(resp.StatusCode, string(data))
}
