package ansi

// Effects

// Bold sets the Bold-bit according to the bool-argument
func (s Style) Bold(f bool) Style {
	if f {
		return s | Bold
	}
	return s &^ Bold
}

// Faint sets the Faint-bit according to the bool-argument
func (s Style) Faint(f bool) Style {
	if f {
		return s | Faint
	}
	return s &^ Faint
}

// Italic sets the Italic-bit according to the bool-argument
func (s Style) Italic(f bool) Style {
	if f {
		return s | Italic
	}
	return s &^ Italic
}

// Inverse sets the Inverse-bit according to the bool-argument
func (s Style) Inverse(f bool) Style {
	if f {
		return s | Inverse
	}
	return s &^ Inverse
}

// Hide sets the Hide-bit according to the bool-argument
func (s Style) Hide(f bool) Style {
	if f {
		return s | Hide
	}
	return s &^ Hide
}

// Strikethru sets the Strikethru-bit according to the bool-argument
func (s Style) Strikethru(f bool) Style {
	if f {
		return s | Strikethru
	}
	return s &^ Strikethru
}

// Foreground colors

// BlackFore set the foreground to black if the argument is 'true'
func (s Style) BlackFore(f bool) Style {
	if f {
		return s&^0b111 | Black
	}
	return s
}

// RedFore set the foreground to red if the argument is 'true'
func (s Style) RedFore(f bool) Style {
	if f {
		return s&^0b111 | Red
	}
	return s
}

// GreenFore set the foreground to green if the argument is 'true'
func (s Style) GreenFore(f bool) Style {
	if f {
		return s&^0b111 | Green
	}
	return s
}

// YellowFore set the foreground to yellow if the argument is 'true'
func (s Style) YellowFore(f bool) Style {
	if f {
		return s&^0b111 | Yellow
	}
	return s
}

// BlueFore set the foreground to blue if the argument is 'true'
func (s Style) BlueFore(f bool) Style {
	if f {
		return s&^0b111 | Blue
	}
	return s
}

// MagentaFore set the foreground to magenta if the argument is 'true'
func (s Style) MagentaFore(f bool) Style {
	if f {
		return s&^0b111 | Magenta
	}
	return s
}

// CyanFore set the foreground to cyan if the argument is 'true'
func (s Style) CyanFore(f bool) Style {
	if f {
		return s&^0b111 | Cyan
	}
	return s
}

// WhiteFore set the foreground to white if the argument is 'true'
func (s Style) WhiteFore(f bool) Style {
	if f {
		return s&^0b111 | White
	}
	return s
}

// BrightFore set the foreground to bright if the argument is 'true'
func (s Style) BrightFore(f bool) Style {
	if f {
		return s | Bright
	}
	return s &^ 0b1000
}

// Background colors

// BlackBack set the background to black if the argument is 'true'
func (s Style) BlackBack(f bool) Style {
	if f {
		return s&^0b111_00000 | Black<<5
	}
	return s
}

// RedBack set the background to red if the argument is 'true'
func (s Style) RedBack(f bool) Style {
	if f {
		return s&^0b111_00000 | Red<<5
	}
	return s
}

// GreenBack set the background to green if the argument is 'true'
func (s Style) GreenBack(f bool) Style {
	if f {
		return s&^0b111_00000 | Green<<5
	}
	return s
}

// YellowBack set the background to yellow if the argument is 'true'
func (s Style) YellowBack(f bool) Style {
	if f {
		return s&^0b111_00000 | Yellow<<5
	}
	return s
}

// BlueBack set the background to blue if the argument is 'true'
func (s Style) BlueBack(f bool) Style {
	if f {
		return s&^0b111_00000 | Blue<<5
	}
	return s
}

// MagentaBack set the background to magenta if the argument is 'true'
func (s Style) MagentaBack(f bool) Style {
	if f {
		return s&^0b111_00000 | Magenta<<5
	}
	return s
}

// CyanBack set the background to cyan if the argument is 'true'
func (s Style) CyanBack(f bool) Style {
	if f {
		return s&^0b111_00000 | Cyan<<5
	}
	return s
}

// WhiteBack set the background to white if the argument is 'true'
func (s Style) WhiteBack(f bool) Style {
	if f {
		return s&^0b111_00000 | White<<5
	}
	return s
}

// BrightBack set the background to bright if the argument is 'true'
func (s Style) BrightBack(f bool) Style {
	if f {
		return s | Bright<<5
	}
	return s &^ 0b1000
}
