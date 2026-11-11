package main

import (
	"github.com/kodarov/go-core/temperature"
)

func main() {
	celsius, err := temperature.ParseCelsius("20")
	if err != nil {
		println(err.Error())
	} else {
		println(celsius)
	}
	value, err := temperature.ParseCelsius("32769")
	if err != nil {
		println(err.Error())
	} else {
		println(value)
	}

	value, err = temperature.ParseCelsius("cold")
	if err != nil {
		println(err.Error())
	} else {
		println(value)
	}

	fahrenheit := temperature.ToFahrenheit(25)
	println(fahrenheit)
	fahrenheit = temperature.ToFahrenheit(21)
	println(fahrenheit)
}
