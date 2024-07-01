# Weather App

Weather App is a simple Go application that fetches weather information either for specified cities or based on the user's current IP location.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [Fetch Weather by City](#fetch-weather-by-city)
  - [Fetch Weather by IP Location](#fetch-weather-by-ip-location)
- [Environment Variables](#environment-variables)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/mzfarshad/weather-app.git
   cd weather-app

## Install dependencies using Go modules:

```
  go mod tidy
```

## Usage

####  Weather App can fetch weather information in two ways: by specifying cities or by using the current IP location.

### Fetch Weather by City

#### To fetch weather based on specific cities, run:

```
  go run main.go <city1> <city2>...
```
#### Replace <city_1>, <city_2>, etc. with the names of the cities you want to check.

### Fetch Weather by IP Location

#### To fetch weather based on your current IP location, simply run:

```
  go run main.go
```
#### This will use your current IP address to determine your location and fetch weather information accordingly.

## Environment Variables

#### Make sure you have set up the following environment variables:

#### OPENWEATHERMAP_API_KEY: Your OpenWeatherMap API key
#### IPINFO_API_KEY: Your IPInfo API key
#### These keys are necessary for fetching weather data from OpenWeatherMap and determining your IP location using IPInfo
