package ansi

import "testing"

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
