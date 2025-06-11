package main

import "fmt"

func greetings() {
	fmt.Println("-------- Welcome to Unit Converter --------")
	fmt.Println("-------------------------------------------")
}

// (0°C × 9/5) + 32 = 32°F
func celsiusToFahrenheit(celcius float32) float32 {

	f := (celcius * (9.0 / 5.0)) + 32
	return f
}

// (0°F − 32) × 5/9 = -17.78°C
func fahrenheitToCelsius(fahrenheit float32) float32 {

	c := (fahrenheit - 32) * (5.0 / 9.0)
	return c
}

// multiply the length value by 3.281
func metreToFeet(meter float32) float32 {
	f := meter * 3.281
	return f
}

//  divide the length value by 3.281
func feetToMetre(feet float32) float32 {
	m := feet / 3.281
	return m
}

func getUnit() string {
	fmt.Println("Which Unit value you have? Type the number")
	fmt.Println("1. Celsius to Fahrenheit")
	fmt.Println("2. Fahrenheit to Celsius")
	fmt.Println("3. Meter to Feet")
	fmt.Println("4. Feet to Meter")
	var unit string
	fmt.Scanln(&unit)
	return unit
}

func getUnitValue(unitNumber string) (float32, string) {
	var unit string
	if unitNumber == "1" {
		unit = "Celsius"
	}
	if unitNumber == "2" {
		unit = "Fahrenheit"
	}
	if unitNumber == "3" {
		unit = "Meter"
	}
	if unitNumber == "4" {
		unit = "Feet"
	}
	fmt.Printf("Type the %s value :- ", unit)
	var value float32
	fmt.Scanln(&value)
	return value, unit
}

func printOutput(unitValue float32, unit string, convertedValue float32, convertedUnit string) {
	fmt.Printf("%2f %s is Equal to %2f %s", unitValue, unit, convertedValue, convertedUnit)
}

func main() {
	greetings()
	unit := getUnit()
	value, unitName := getUnitValue(unit)

	switch unit {
	case "1":
		result := celsiusToFahrenheit(value)
		printOutput(value, unitName, result, "Fahrenheit")
	case "2":
		result := fahrenheitToCelsius(value)
		printOutput(value, unitName, result, "Celsius")
	case "3":
		result := metreToFeet(value)
		printOutput(value, unitName, result, "Feet")
	case "4":
		result := feetToMetre(value)
		printOutput(value, unitName, result, "Meter")
	}
}
