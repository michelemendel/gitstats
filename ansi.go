// ANSI Escape Sequences (Codes)
package main

// Ansi Escape Sequences (Codes)
var Ansi = struct {
	// Text Attributes
	AttrOff, Bold, Underscore, Blink, ReverseVideo, Concealed string
	// Foreground Color
	Black, Red, Green, Yellow, Blue, Magenta, Cyan, White string
	// Background Color
	BgBlue string
}{
	// Text Attributes
	AttrOff:      "\033[0m",
	Bold:         "\033[1m",
	Underscore:   "\033[4m",
	Blink:        "\033[5m",
	ReverseVideo: "\033[7m",
	Concealed:    "\033[8m",
	// Foreground Color
	Black:   "\033[30m",
	Red:     "\033[31m",
	Green:   "\033[32m",
	Yellow:  "\033[33m",
	Blue:    "\033[34m",
	Magenta: "\033[35m",
	Cyan:    "\033[36m",
	White:   "\033[37m",
	// Background Color
	BgBlue: "\033[44m",
}

// Background colors
// 40	Black
// 41	Red
// 42	Green
// 43	Yellow
// 44	Blue
// 45	Magenta
// 46	Cyan
// 47	White
