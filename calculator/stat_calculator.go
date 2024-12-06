package calculator

import (
	"math"
)

func CalculateStat(base, iv, ev, level int, natureModifier float64) int {
	// Apply the formula
	stat := ((2*base + iv + ev/4) * level / 100) + 5
	// Apply the nature modifier
	statWithNature := float64(stat) * natureModifier
	// Return the stat as an integer (rounded down)
	return int(math.Floor(statWithNature))
}