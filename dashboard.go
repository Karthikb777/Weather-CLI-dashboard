package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// URL for the data - from openweathermap.org
const url = "https://api.openweathermap.org/data/2.5/onecall?lat=<LATITUDE>&lon=<LONGITUDE>&units=metric&exclude=minutely,daily&appid=<YOUR API KEY>"

var (
	pressure, clouds, humidity, temp, feels_like float64
)

// decoding json data
func parser(body []uint8) map[string]interface{} {
	temp := make(map[string]interface{})
	err := json.Unmarshal(body, &temp)
	if err == nil {
		return temp
	}
	return temp
}

// cherrypicking the data that we want
func (d *Data) allocate(data map[string]interface{}) {
	for k, _ := range data {
		switch k {
		case "current":
			d.current = data["current"]
		case "hourly":
			d.hourly = data["hourly"]
		}
	}
}

type Data struct {
	current, hourly interface{}
}

func main() {
	data := Data{}
	res, err := http.Get(url)
	if err == nil {
		body, errors := ioutil.ReadAll(res.Body)
		if errors == nil {
			res := parser(body)
			data.allocate(res)
		}
	}
	currentData := data.current.(map[string]interface{})
	for k, v := range currentData {
		//	fmt.Printf("%v %v\n", k, v)
		switch k {
		case "temp":
			temp = v.(float64)
		case "feels_like":
			feels_like = v.(float64)
		case "pressure":
			pressure = v.(float64)
		case "clouds":
			clouds = v.(float64)
		case "humidity":
			humidity = v.(float64)
		}
	}

	/*	hourlyData := data.hourly.([]interface{})
		for k, v := range hourlyData {
			fmt.Println(k, v)
			fmt.Println("\n\n")
		} */

	fmt.Printf("\n  Hello! \n  Current weather:\n -----------------------\n  Temperature: %v°C \n  Feels like: %v°C \n  Pressure: %v hPa \n  Clouds: %v percent\n  Humidity: %v percent\n ----------------------- \n", temp, feels_like, pressure, clouds, humidity)
	fmt.Println("\n")

}
