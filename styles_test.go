package ansi

import (
	"strings"
	"testing"
)

func TestNewStyle(t *testing.T) {
	tests := []struct {
		name   string
		styles []Style
		want   Style
	}{
		{"default", []Style{Default}, Default},
		{"bold+red", []Style{Bold, Red}, Bold | Red},
		{"red+bold", []Style{Red, Bold}, Bold | Red},
		{"swedish-1", []Style{Yellow, Blue.Background()}, Yellow | Blue<<5},
		{"swedish-2", []Style{Bright, Yellow, Blue.Background()}, Bright | Yellow | Blue<<5},
		{"swedish-3", []Style{Yellow, Bright, Blue.Background()}, Bright | Yellow | Blue<<5},
		{"swedish-4", []Style{Yellow, Bright, Blue.Background().BrightBack(true)}, Bright | Yellow | Blue<<5 | Bright<<5},
		{"swedish-5", []Style{Blue.Background().BrightBack(true), Yellow, Bright}, Bright | Yellow | Blue<<5 | Bright<<5},
		{"swedish-6", []Style{Blue.Background().BrightBack(true), Yellow, Bright}, Bright | Yellow | Blue<<5 | Bright<<5},
		{"red->green", []Style{Red.GreenFore(true)}, Green},
		{"!red->green", []Style{Red.GreenFore(false)}, Red},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStyle(tt.styles...); got != tt.want {
				t.Errorf("Style() = %b, want %b", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name   string
		styles []Style
		want   string
	}{
		{"default", []Style{Default}, "\033[0;39;49m"},
		{"bold+red", []Style{Bold, Red}, "\033[1;31m"},
		{"swedish-1", []Style{Yellow, Blue.Background()}, "\033[0;33;44m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStyle(tt.styles...).String(); got != tt.want {
				got = strings.Replace(got, "\033", "^", -1)
				tt.want = strings.Replace(tt.want, "\033", "^", -1)
				t.Errorf("Style() = %s, want %s", got, tt.want)
			}
		})
	}
}
