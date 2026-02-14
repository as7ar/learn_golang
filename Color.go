package main

import "fmt"

type Color int

const (
	Red Color = iota
	Yellow
	Green
	Blue
	Purple
	Cyan
	Black
	White
)

func (c Color) String() string {
	switch c {
	case Red:
		return "\033[0;31mred\u001B[0m"
	case Yellow:
		return "\033[0;33myellow\u001B[0m"
	case Green:
		return "\033[0;32mgreen\u001B[0m"
	case Blue:
		return "\033[0;34mblue\u001B[0m"
	case Purple:
		return "\033[0;35mpurple\u001B[0m"
	case Cyan:
		return "\033[0;36mcyan\u001B[0m"
	case Black:
		return "\033[0;30mblack\033[0m"
	case White:
		return "\033[0;37mwhite\u001B[0m"
	default:
		panic(fmt.Errorf("unknown Color: %s", c))
	}
}
