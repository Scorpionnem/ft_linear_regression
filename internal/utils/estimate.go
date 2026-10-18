package utils

func EstimatePrice(km float64, theta0 float64, theta1 float64) float64 {
	return (theta0 + (theta1 * km))
}
