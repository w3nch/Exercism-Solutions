// Package weather helps report the current weather for cities in Goblinocus.
package weather

// CurrentCondition holds a short description of the weather,
// such as rain, snow, or sunshine.
var CurrentCondition string

// CurrentLocation holds the name of the city the weather report refers to.
var CurrentLocation string

// Forecast returns a short weather report for a city based on
// the provided location and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
