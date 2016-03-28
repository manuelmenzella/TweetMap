package tools

import (
	"math"
)

func Round(n float64) float64 {
	if n > 0 {
		return math.Floor(n + 0.5)
	} else {
		return math.Ceil(n - 0.5)
	}
}
