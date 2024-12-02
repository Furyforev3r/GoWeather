# GoWeather

A simple command-line application written in Go that fetches the current weather information for a given city using the OpenWeather API.

## Features

- Fetches weather data including temperature, humidity, wind speed, and weather description.
- Supports input of any city name through command-line flags.

## Prerequisites

- Go 1.16 or higher
- An OpenWeather API key

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/furyforev3r/GoWeather.git
   cd GoWeather
   ```

2. Install dependencies:

   The project uses the `github.com/joho/godotenv` package to load environment variables.

   ```bash
   go mod tidy
   ```

3. Create a `.env` file in the root of the project and add your OpenWeather API key:

   ```env
   API_KEY=your_openweather_api_key_here
   ```

## Usage

To run the application, use the following command:

```bash
go run main.go --city="London"
```

or

```bash
go run main.go --city London
```


You can replace `"London"` with any city you want to get the weather for. The application will output the weather details, including temperature, wind speed, and humidity.

## Example Output

```
------------------------------------------------------------------
Weather for London, GB:
Temperature: 12.34°C
Feels Like: 10.45°C
Weather: Clear sky
Wind Speed: 2.34 m/s
Humidity: 65%
Cloudiness: 10%
------------------------------------------------------------------
```

## Notes

- The temperature is displayed in Celsius (calculated from the Kelvin value returned by the OpenWeather API).
- The application clears the console before showing the weather data to ensure the output is clean.
- If you don't specify a city, the application will default to "London".
