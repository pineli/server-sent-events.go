package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"time"
)

type Sensors struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

func getSensorValues() *Sensors {
	data := &Sensors{
		Temperature: generateTemperature(),
		Humidity:    generateHumidity(),
	}

	return data
}

func generateTemperature() float64 {
	rand.Seed(time.Now().UnixNano())
	currentTemperature := rand.Float64()*120 - 20

	for i := 0; i < 10; i++ {
		step := rand.Float64()*2 - 1
		currentTemperature += step

		if currentTemperature < -20 {
			currentTemperature = -20
		} else if currentTemperature > 100 {
			currentTemperature = 100
		}
	}

	return math.Round(currentTemperature*100) / 100
}

func generateHumidity() float64 {
	var last float64 = 50
	var maxDelta float64 = 5

	for {
		delta := rand.Float64()*maxDelta*2 - maxDelta
		newHumidity := last + delta

		if newHumidity < 0 {
			newHumidity = 0
		} else if newHumidity > 100 {
			newHumidity = 100
		}

		last = newHumidity
		waitTime := time.Duration(rand.Intn(500)+1) * time.Millisecond
		time.Sleep(waitTime)

		return newHumidity
	}
}

func main() {
	http.HandleFunc("/sensor-data-events", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		for {
			// Enviando data e hora
			// currentTime := time.Now()
			// fmt.Fprintf(w, "event: sensor-data\n")
			// fmt.Fprintf(w, "data: %s\n\n", currentTime.Format("2006-01-02 15:04:05"))
			// w.(http.Flusher).Flush()
			// time.Sleep(1000 * time.Millisecond)

			// Simlando 2 sensores IoT
			dataString, _ := json.Marshal(getSensorValues())
			fmt.Fprintf(w, "event: sensor-data\n")
			fmt.Fprintf(w, "data: %s\n\n", dataString)
			w.(http.Flusher).Flush()
			time.Sleep(1000 * time.Millisecond)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "page/index.html")
	})
	fmt.Println("Servidor executando na porta 8081")

	http.ListenAndServe(":8081", nil)
}
