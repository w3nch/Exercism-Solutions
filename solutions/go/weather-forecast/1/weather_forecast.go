// Package weather provides weather forecast information.
package weather

var (
	// CurrentCondition stores the current weather condition.
	CurrentCondition string

	// CurrentLocation stores the current weather location.
	CurrentLocation string
)

// Forecast returns the current location and its weather condition.
func Forecast(city, condition string) string {
	CurrentLocation = city
	CurrentCondition = condition

	return CurrentLocation + " - current weather condition: " + CurrentCondition
}