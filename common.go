package goutil

import "math"

//FloatRound round for float64 `f` keeping `n` decimal
func FloatRound(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	offset := 0.5
	if f < 0 {
		offset = -offset
	}
	return math.Trunc((f+offset/pow10_n)*pow10_n) / pow10_n
}
