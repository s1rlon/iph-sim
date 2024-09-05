package game

import (
	"fmt"
)

func formatNumber(value float64) string {
	switch {
	case value >= 1_000_000_000_000:
		return fmt.Sprintf("%.2fT", value/1_000_000_000_000)
	case value >= 1_000_000_000:
		return fmt.Sprintf("%.2fB", value/1_000_000_000)
	case value >= 1_000_000:
		return fmt.Sprintf("%.2fM", value/1_000_000)
	case value >= 1_000:
		return fmt.Sprintf("%.2fK", value/1_000)
	default:
		return fmt.Sprintf("%.2f", value)
	}
}
