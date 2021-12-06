package cars

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	//panic("SuccessRate not implemented")
	if speed == 0 {
		return 0.0
	} else if speed <= 4 {
		return 1.0
	} else if speed <= 8 {
		return 0.9
	} else if speed <= 10 {
		return 0.77
	} else {
		return 0.0
	}
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	//panic("CalculateProductionRatePerHour not implemented") 221
	var baseRate = 221.0 * SuccessRate(speed)
	return float64(speed) * baseRate
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	//panic("CalculateProductionRatePerMinute not implemented")
	return int(CalculateProductionRatePerHour(speed)) / 60
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	//panic("CalculateLimitedProductionRatePerHour not implemented")
	production := CalculateProductionRatePerHour(speed)
	if production <= limit {
		return production
	} else {
		return limit
	}
}
