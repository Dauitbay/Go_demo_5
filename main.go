package main

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"flag"
	"fmt"
)

func main() {
	fmt.Println("New project")
	city := flag.String("city", "", "User's city")
	format := flag.Int("format", 1, "Weather output format")
	flag.Parse()
	fmt.Println(*city)
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(geoData)
	weatherData := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)
	// fmt.Println(*format)
}

// r := strings.NewReader("Hello i am a reader")
// block := make([]byte, 4)
// for {
// 	_, err := r.Read(block)
// 	fmt.Printf("%q\n", block)
// 	if err == io.EOF{
// 		break
// 	}
// }
