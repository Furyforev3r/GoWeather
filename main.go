package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"flag"
	"os/exec"
	"runtime"
		
	"github.com/joho/godotenv"
)

type WeatherResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Main struct {
		Temp       float64 `json:"temp"`
		FeelsLike  float64 `json:"feels_like"`
		TempMin    float64 `json:"temp_min"`
		TempMax    float64 `json:"temp_max"`
		Pressure   int     `json:"pressure"`
		Humidity   int     `json:"humidity"`
		SeaLevel   int     `json:"sea_level"`
		GrndLevel  int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	TimeZone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func loadDotEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getWeather(location string) (*http.Response, error) {
	apiKey := os.Getenv("API_KEY")
	apiURL := "http://api.openweathermap.org/data/2.5/weather?q=" + location + "&appid=" + apiKey

	response, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching weather: %v", err)
	}
	return response, nil
}

func clearConsole() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printWeatherData(weatherData WeatherResponse) {
	fmt.Println("------------------------------------------------------------------")
	fmt.Printf("Weather for %s, %s:\n", weatherData.Name, weatherData.Sys.Country)
	fmt.Printf("Temperature: %.2f°C\n", weatherData.Main.Temp-273.15)
	fmt.Printf("Feels Like: %.2f°C\n", weatherData.Main.FeelsLike-273.15)
	fmt.Printf("Weather: %s\n", weatherData.Weather[0].Description)
	fmt.Printf("Wind Speed: %.2f m/s\n", weatherData.Wind.Speed)
	fmt.Printf("Humidity: %d%%\n", weatherData.Main.Humidity)
	fmt.Printf("Cloudiness: %d%%\n", weatherData.Clouds.All)
	fmt.Println("------------------------------------------------------------------")
}

func main() {
	loadDotEnv()
	
	location := flag.String("city", "London", "The city to fetch weather for")
	flag.Parse()
	
	loadDotEnv()

	response, err := getWeather(*location)
	
	if err != nil {
		log.Fatal("Error fetching weather data: " + err.Error())
	}
	defer response.Body.Close()

	var weatherData WeatherResponse
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&weatherData); err != nil {
		log.Fatal("Error parsing weather response: ", err)
	}

	clearConsole()
	printWeatherData(weatherData)
}
