package city

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"weather-app/model"
)

const BaseUrl = "https://api.openweathermap.org/data/2.5/weather"

var ApiKey string

func FetchWeather(city string, wg *sync.WaitGroup) {
	defer wg.Done()
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", BaseUrl, city, ApiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error fetching weather data: %v\n", err)
		return
	}
	defer resp.Body.Close()

	var weather model.WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		fmt.Printf("error decoding weather data for %s: %v\n", city, err)
		return
	}
	if len(weather.Weather) > 0 {
		fmt.Printf("Weather in %s: %.2f°C, %s\n", weather.Name, weather.Main.Temp, weather.Weather[0].Description)
	} else {
		fmt.Printf("Weather in %s: %.2f°C, no weather description available\n", weather.Name, weather.Main.Temp)
	}
}
