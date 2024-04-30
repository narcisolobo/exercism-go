package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	d := successRate / 100
	return float64(productionRate) * d
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	d := successRate / 100
	perHour := float64(productionRate) * d
	return int(perHour / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	mult := carsCount / 10
	remainder := carsCount % 10
	costPer10 := 95000
	costPer1 := 10000
	return uint(mult*costPer10 + remainder*costPer1)
}
