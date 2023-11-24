package main

import (
	"log"
	"math/rand"
	"time"

	weatherController "github.com/bangpenn/assignment-3/controllers"
	"github.com/bangpenn/assignment-3/models"
)

type WeatherData struct {
	Water       float64
	Wind        float64
	WaterStatus string
	WindStatus  string
}

func main() {
	models.ConnectDatabase()
	updateDataWithTicker()
}

func updateDataWithTicker() {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			weatherData := generateRandomWeatherData()
			printWeatherData(weatherData)
			weatherController.Update(weatherData.Water, weatherData.Wind, weatherData.WaterStatus, weatherData.WindStatus)
		}
	}
}

func generateRandomWeatherData() WeatherData {
	water := rand.Float64()*100 + 1
	wind := rand.Float64()*100 + 1

	waterStatus := getStatus(water)
	windStatus := getStatus(wind)

	return WeatherData{
		Water:       water,
		Wind:        wind,
		WaterStatus: waterStatus,
		WindStatus:  windStatus,
	}
}

func getStatus(value float64) string {
	if value < 5 {
		return "aman"
	} else if value >= 6 && value <= 8 {
		return "siaga"
	} else {
		return "bahaya"
	}
}

func printWeatherData(data WeatherData) {
	log.Printf("Water: %.2f m (%s), Wind: %.2f m/s (%s)\n", data.Water, data.WaterStatus, data.Wind, data.WindStatus)
}
