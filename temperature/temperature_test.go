package temperature

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestParseCelsius(t *testing.T) {
	tests := []struct {
		name    string
		raw     string
		want    int16
		wantErr error
	}{
		{name: "zero", raw: "0", want: 0},
		{name: "min int16", raw: "-32768", want: -32768},
		{name: "max int16", raw: "32767", want: 32767},
		{name: "negative", raw: "-10", want: -10},
		{name: "positive", raw: "25", want: 25},
		{name: "positive overflow", raw: "32768", wantErr: strconv.ErrRange},
		{name: "negative overflow", raw: "-32769", wantErr: strconv.ErrRange},
		{name: "invalid string", raw: "cold", wantErr: strconv.ErrSyntax},
		{name: "empty string", raw: "", wantErr: strconv.ErrSyntax},
		{name: "fractional input", raw: "21.5", wantErr: strconv.ErrSyntax},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCelsius(tt.raw)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("ParseCelsius(%q) error = %v, want %v", tt.raw, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("ParseCelsius(%q) = %d, want %d", tt.raw, got, tt.want)
			}
			if tt.wantErr != nil && !strings.Contains(err.Error(), strconv.Quote(tt.raw)) {
				t.Errorf("error %q does not contain original input %q", err, tt.raw)
			}
		})
	}
}

func TestToFahrenheit(t *testing.T) {
	tests := []struct {
		name    string
		celsius int16
		want    float64
	}{
		{name: "zero", celsius: 0, want: 32},
		{name: "min int16", celsius: -32768, want: -58950.4},
		{name: "max int16", celsius: 32767, want: 59012.6},
		{name: "negative", celsius: -10, want: 14},
		{name: "matching scales", celsius: -40, want: -40},
		{name: "fractional result", celsius: 21, want: 69.8},
		{name: "positive", celsius: 25, want: 77},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToFahrenheit(tt.celsius)
			if math.IsNaN(got) || math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("ToFahrenheit(%d) = %g, want %g", tt.celsius, got, tt.want)
			}
		})
	}
}
