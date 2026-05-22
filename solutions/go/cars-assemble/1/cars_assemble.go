package cars

const (
	successRateScale = 100.0

	minutesPerHour = 60

	carPrice      = 10000
	bundleSize    = 10
	bundlePrice   = 95000
)

// CalculateWorkingCarsPerHour returns the number of successfully
// produced cars per hour.
func CalculateWorkingCarsPerHour(
	productionRate int,
	successRate float64,
) float64 {
	return float64(productionRate) *
		(successRate / successRateScale)
}

// CalculateWorkingCarsPerMinute returns the number of successfully
// produced cars per minute.
func CalculateWorkingCarsPerMinute(
	productionRate int,
	successRate float64,
) int {
	carsPerHour := CalculateWorkingCarsPerHour(
		productionRate,
		successRate,
	)

	return int(carsPerHour / minutesPerHour)
}

// CalculateCost returns the total production cost.
func CalculateCost(carsCount int) uint {
	if carsCount <= 0 {
		return 0
	}

	bundles := carsCount / bundleSize
	remainingCars := carsCount % bundleSize

	totalCost := (bundles * bundlePrice) +
		(remainingCars * carPrice)

	return uint(totalCost)
}