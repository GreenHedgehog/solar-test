package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/GreenHedgehog/solar-test/db"
	"github.com/GreenHedgehog/solar-test/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type config struct {
	DB     db.Config `json:"db"`
	Server struct {
		Port int `json:"port"`
	} `json:"server"`
}

var (
	errNoConfig      = errors.New("No config")
	errInvalidConfig = errors.New("Invalid config")
)

func main() {
	var conf config

	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(errNoConfig, err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&conf)
	if err != nil {
		log.Fatal(errInvalidConfig, err)
	}

	err = db.Init(&conf.DB)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handlers.Add(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.Server.Port)))

}
