package utils

import (
	"fmt"
	"strconv"
)

func Trucate(num float64, decimal int64) float64 {
	s := "%." + fmt.Sprintf("%d", decimal) + "f"
	truncated, _ := strconv.ParseFloat(fmt.Sprintf(s, num), 64)

	return truncated
}
