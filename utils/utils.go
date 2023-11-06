package utils

import "math"

func ConvertToBillions(n float64) float64 {
	return math.Round(n/1e8*100) / 100
}
