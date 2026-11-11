package temperature

import (
	"fmt"
	"strconv"
)

func ParseCelsius(raw string) (int16, error) {
	i, err := strconv.ParseInt(raw, 10, 16)
	if err != nil {
		return 0, fmt.Errorf("failed to parse %q celsius value: %w", raw, err)
	}
	return int16(i), nil
}

func ToFahrenheit(c int16) float64 {
	return float64(c)*9/5 + 32
}
