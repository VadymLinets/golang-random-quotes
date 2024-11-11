package tools

import "math/rand/v2"

type numbers interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Sum[U any, T numbers](arr []U, f func(U) T) T {
	var accumulator T
	for _, value := range arr {
		accumulator += f(value)
	}

	return accumulator
}

func RandomPercent(maxValue float64) float64 {
	return rand.Float64() * maxValue
}
