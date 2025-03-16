package main

import (
	"Golang/weather/geo"
	"Golang/weather/weather"
	"flag"
	"fmt"
)

func main() {
	// name := flag.String("name", "Anton", "Name of user") // (название флага, дефолтное значение, описание)
	// age := flag.Int("age", 18, "Age of user")
	// isAdmin := flag.Bool("isAdmin", false, "Is user admin")
	city := flag.String("city", "", "City of user")
	format := flag.Int("format", 1, "Формат вывода погоды")
	flag.Parse()
	fmt.Println(*city)
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(geoData)
	weatherData := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)

}
