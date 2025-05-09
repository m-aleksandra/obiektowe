package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB // global db

type OpenWeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Name string `json:"name"`
}

type Weather struct {
	gorm.Model
	City        string
	Temperature float64
	Description string
}

type WeatherRequest struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

type WeatherProxy struct{}

func (wp *WeatherProxy) Fetch(lat, lon string) (*OpenWeatherResponse, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("Missing API_KEY")
	}

	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=metric",
		lat, lon, apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var parsed OpenWeatherResponse
	if err := json.Unmarshal(bodyBytes, &parsed); err != nil {
		return nil, err
	}

	return &parsed, nil
}

func weatherHandler(c echo.Context) error {
	req := new(WeatherRequest)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "Invalid JSON body")
	}

	if req.Lat == "" || req.Lon == "" {
		return c.String(http.StatusBadRequest, "Missing lat or lon")
	}

	proxy := WeatherProxy{}
	result, err := proxy.Fetch(req.Lat, req.Lon)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error calling external API")
	}

	entry := Weather{
		City:        result.Name,
		Temperature: result.Main.Temp,
		Description: result.Weather[0].Description,
	}

	db.Create(&entry)

	return c.JSON(http.StatusOK, entry)
}

func loadInitialData() {
	initialData := []Weather{
		{City: "Warsaw", Temperature: 18.5, Description: "Sunny"},
		{City: "Gdansk", Temperature: 15.2, Description: "Cloudy"},
		{City: "Krakow", Temperature: 16.7, Description: "Rainy"},
	}

	var count int64
	db.Model(&Weather{}).Count(&count)
	if count == 0 {
		for _, entry := range initialData {
			db.Create(&entry)
		}
	}
}

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("weather.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&Weather{})
	loadInitialData()

	e := echo.New()
	e.POST("/weather", weatherHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
