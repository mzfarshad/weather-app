package location

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"weather-app/city"
	"weather-app/model"
)

const ipInfoUrl = "https://ipinfo.io/json?token="

func FetchWeatherByLocation(lat, lon float64) {
	url := fmt.Sprintf("%s?lat=%f&lon=%f&appid=%s&units=metric", city.BaseUrl, lat, lon, city.ApiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error fetching weather data: %v\n", err)
		return
	}
	defer resp.Body.Close()
	var weather model.WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		fmt.Printf("error decoding weather data: %v\n", err)
		return
	}
	if len(weather.Weather) > 0 {
		fmt.Printf("Weather in %s: %.2f°C, %s\n", weather.Name, weather.Main.Temp, weather.Weather[0].Description)
	} else {
		fmt.Printf("Weather in %s: %.2f°C, no weather description available\n", weather.Name, weather.Main.Temp)
	}
}

func GetLocationFromIP() (float64, float64, error) {
	ipInfoApiKey := os.Getenv("IPINFO_API_KEY")
	if ipInfoApiKey == "" {
		return 0, 0, fmt.Errorf("IPINFO_API_KEY environment variable is not set")
	}

	resp, err := http.Get(ipInfoUrl + ipInfoApiKey)
	if err != nil {
		return 0, 0, fmt.Errorf("error fetching IP info: %v", err)
	}
	defer resp.Body.Close()

	var ipInfo model.IPInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&ipInfo); err != nil {
		return 0, 0, fmt.Errorf("error decoding IP info: %v", err)
	}

	var lat, lon float64
	fmt.Sscanf(ipInfo.Loc, "%f,%f", &lat, &lon)
	return lat, lon, nil
}
