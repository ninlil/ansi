package ansi

// Effects

func (s Style) Bold(f bool) Style {
	if f {
		return s | Bold
	}
	return s &^ Bold
}

func (s Style) Faint(f bool) Style {
	if f {
		return s | Faint
	}
	return s &^ Faint
}

func (s Style) Italic(f bool) Style {
	if f {
		return s | Italic
	}
	return s &^ Italic
}

func (s Style) Inverse(f bool) Style {
	if f {
		return s | Inverse
	}
	return s &^ Inverse
}

func (s Style) Hide(f bool) Style {
	if f {
		return s | Hide
	}
	return s &^ Hide
}

func (s Style) Strikethru(f bool) Style {
	if f {
		return s | Strikethru
	}
	return s &^ Strikethru
}

// Foreground colors

func (s Style) BlackFore(f bool) Style {
	if f {
		return s&^0b111 | Black
	}
	return s
}

func (s Style) RedFore(f bool) Style {
	if f {
		return s&^0b111 | Red
	}
	return s
}

func (s Style) GreenFore(f bool) Style {
	if f {
		return s&^0b111 | Green
	}
	return s
}

func (s Style) YellowFore(f bool) Style {
	if f {
		return s&^0b111 | Yellow
	}
	return s
}

func (s Style) BlueFore(f bool) Style {
	if f {
		return s&^0b111 | Blue
	}
	return s
}

func (s Style) MagentaFore(f bool) Style {
	if f {
		return s&^0b111 | Magenta
	}
	return s
}

func (s Style) CyanFore(f bool) Style {
	if f {
		return s&^0b111 | Cyan
	}
	return s
}

func (s Style) WhiteFore(f bool) Style {
	if f {
		return s&^0b111 | White
	}
	return s
}

func (s Style) BrightFore(f bool) Style {
	if f {
		return s | Bright
	}
	return s &^ 0b1000
}

// Background colors

func (s Style) BlackBack(f bool) Style {
	if f {
		return s&^0b111_00000 | Black<<5
	}
	return s
}

func (s Style) RedBack(f bool) Style {
	if f {
		return s&^0b111_00000 | Red<<5
	}
	return s
}

func (s Style) GreenBack(f bool) Style {
	if f {
		return s&^0b111_00000 | Green<<5
	}
	return s
}

func (s Style) YellowBack(f bool) Style {
	if f {
		return s&^0b111_00000 | Yellow<<5
	}
	return s
}

func (s Style) BlueBack(f bool) Style {
	if f {
		return s&^0b111_00000 | Blue<<5
	}
	return s
}

func (s Style) MagentaBack(f bool) Style {
	if f {
		return s&^0b111_00000 | Magenta<<5
	}
	return s
}

func (s Style) CyanBack(f bool) Style {
	if f {
		return s&^0b111_00000 | Cyan<<5
	}
	return s
}

func (s Style) WhiteBack(f bool) Style {
	if f {
		return s&^0b111_00000 | White<<5
	}
	return s
}

func (s Style) BrightBack(f bool) Style {
	if f {
		return s | Bright<<5
	}
	return s &^ 0b1000
}
