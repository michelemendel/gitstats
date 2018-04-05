// ANSI Escape Sequences (Codes)
package main

// Text attributes
// 0	All attributes off
// 1	Bold on
// 4	Underscore (on monochrome display adapter only)
// 5	Blink on
// 7	Reverse video on
// 8	Concealed on

// Foreground colors
// 30	Black
// 31	Red
// 32	Green
// 33	Yellow
// 34	Blue
// 35	Magenta
// 36	Cyan
// 37	White

// Background colors
// 40	Black
// 41	Red
// 42	Green
// 43	Yellow
// 44	Blue
// 45	Magenta
// 46	Cyan
// 47	White

// Text Attributes
const (
	AttrOff = "\033[0m"
)

// Foreground Color
const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

// Background Color
const (
	BgBlue = "\033[44m"
)
