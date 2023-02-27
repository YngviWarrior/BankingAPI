package utils

import "math"

func ToFixed(num float64, precision int64) float64 {
	return math.Floor(num*math.Pow(10, float64(precision))) / math.Pow(10, float64(precision))
}
