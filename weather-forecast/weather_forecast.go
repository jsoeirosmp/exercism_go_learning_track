// Package weather provides methods for weather forecasting.
package weather

// CurrentCondition is the local weather condition.
var CurrentCondition string

// CurrentLocation is the current location.
var CurrentLocation string

// Forecast function returns a string containing the current location and weather condition
// based off of the two strings it receives (city and condition).
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
