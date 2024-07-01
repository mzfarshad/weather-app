package main

import (
	"fmt"
	"os"
	"sync"
	"weather-app/city"
	"weather-app/location"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("error loading .env file: %v\n", err)
		return
	}

	city.ApiKey = os.Getenv("WEATHER_API_KEY")
	if city.ApiKey == "" {
		fmt.Println("WEATHER_API_KEY environment variable is not set!")
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: weather-app <city_1> <city_2> ... ")
		lat, lon, err := location.GetLocationFromIP()
		if err != nil {
			fmt.Printf("error getting location from IP: %v\n", err)
			return
		}
		location.FetchWeatherByLocation(lat, lon)
		return
	}

	cities := os.Args[1:]
	var wg sync.WaitGroup
	wg.Add(len(cities))

	for _, cty := range cities {
		go city.FetchWeather(cty, &wg)
	}
	wg.Wait()
}
