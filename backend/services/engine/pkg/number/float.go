package number

import (
	"math"
	"strconv"
)

func Float64ToString(n float64) string {
	return strconv.FormatFloat(n, 'f', -1, 64)
}

func roundNumber(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func Round(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(roundNumber(num*output)) / output
}

func StringToFloat64(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return f, err
	}
	return f, nil
}
