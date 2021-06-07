package ansi

// Style contains the bits for all colors and styles
type Style int

// Bitmask
/*
 0	Red Foreground
 1	Green Foreground
 2  Blue Foreground
 3	Bright Foreground
 4	Foreground specified

 5	Red Background
 6  Green Background
 7	Blue Background
 8	Bright Background
 9  Background specified

 10 Bold
 11 Faint
 12 Italic
 13 Inverse
 14	Hidden
*/

// Colors and flags for all different styles
const (
	Default Style = 0

	Black   Style = 0b10000
	Red     Style = 0b10001
	Green   Style = 0b10010
	Yellow  Style = 0b10011
	Blue    Style = 0b10100
	Magenta Style = 0b10101
	Cyan    Style = 0b10110
	White   Style = 0b10111
	Bright  Style = 0b11000

	Bold       Style = 1 << 10
	Faint      Style = 1 << 11
	Italic     Style = 1 << 12
	Inverse    Style = 1 << 13
	Hide       Style = 1 << 14
	Strikethru Style = 1 << 15
)

// Background uses the current foreground color as the background
func (s Style) Background() Style {
	if s&0b11111 != 0 {
		return (s & 0b11111) << 5 // extract foreground and shift into background
	}
	return 0
}

// NewStyle create a new style from the list of arguments
func NewStyle(styles ...Style) Style {
	var style Style
	for _, v := range styles {

		// clear fore and background if setting a replacement color (dont want bits to overwrite) (keep brightness)
		if v&0b01000 == 0 && v&0b10000 != 0 { // not setting bright, but a color
			style = style &^ 0b00111 //) | (style & 0b01000)
		}
		if v&0b11111_00000 != 0 {
			style = (style &^ 0b11111_00000) | (style & 0b01000_00000)
		}

		style |= v
	}
	return style
}

func (s Style) flags2values() []int {
	var values = make([]int, 0, 4) // 4 = default initial capacity of attributes
	if s&Bold != 0 {
		values = append(values, 1)
	}
	if s&Faint != 0 {
		values = append(values, 2)
	}
	if s&Italic != 0 {
		values = append(values, 3)
	}
	if s&Inverse != 0 {
		values = append(values, 7)
	}
	if s&Hide != 0 {
		values = append(values, 8)
	}
	if s&Strikethru != 0 {
		values = append(values, 9)
	}
	return values
}

// String converts the Style into the corresponding ANSI-Escape sequence
func (s Style) String() string {

	if s == Default {
		return "\033[0;39;49m"
	}

	values := s.flags2values()

	// Effects
	if len(values) == 0 {
		values = append(values, 0)
	}
	// colors
	if fg := makeColor(s, 30); fg != -1 {
		values = append(values, fg)
	}
	if bg := makeColor(s>>5, 40); bg != -1 {
		values = append(values, bg)
	}

	// create ANSI escape-sequence
	var runes = make([]rune, 0, 8) // 8 is default initial capacity of escape-sequence
	runes = append(runes, '\033', '[')
	for i, n := range values {
		if i > 0 {
			runes = append(runes, ';')
		}
		if n >= 100 {
			runes = append(runes, rune('0'+(n/100%10)))
		}
		if n >= 10 {
			runes = append(runes, rune('0'+(n/10%10)))
		}
		runes = append(runes, rune('0'+(n%10)))
	}
	return string(append(runes, 'm'))
}

func makeColor(s Style, index int) int {
	color := int(s)
	if color&0b10000 == 0 {
		return -1
	}
	if color&0b01000 != 0 {
		index += 60
	}

	return index + color&0b111
}
