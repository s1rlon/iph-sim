package game

import (
	"fmt"
)

func FormatNumber(value float64) string {
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

func Add(a, b int) int {
	return a + b
}

func GetOreValue(oreName string) float64 {
	if value, exists := OreValues[oreName]; exists {
		return value
	}
	return 0.0
}

func FormatTime(seconds float64) string {
	hours := int(seconds) / 3600
	minutes := (int(seconds) % 3600) / 60
	secs := int(seconds) % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, secs)
}
