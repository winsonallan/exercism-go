// Package weather provides tools to forecast weather.
package weather

var (
    // CurrentCondition stores the current weather condition that is being forecasted.
	CurrentCondition string

    // CurrentLocation stores the current location of the weather that is being forecasted.
	CurrentLocation  string
)

// Forecast returns the weather forecast of a location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
