package vt10x

// ANSI color values
const (
	Black Color = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	LightGrey
	DarkGrey
	LightRed
	LightGreen
	LightYellow
	LightBlue
	LightMagenta
	LightCyan
	White
)

// Default colors are potentially distinct to allow for special behavior.
// For example, a transparent background. Otherwise, the simple case is to
// map default colors to another color.
const (
	DefaultFG Color = 1<<24 + iota
	DefaultBG
	DefaultCursor
)

// Color maps to the ANSI colors [0, 16) and the xterm colors [16, 256).
type Color uint32

// ANSI returns true if Color is within [0, 16).
func (c Color) ANSI() bool {
	return (c < 16)
}

// IsDefault returns true if this is a default color (DefaultFG, DefaultBG, DefaultCursor).
func (c Color) IsDefault() bool {
	return c >= DefaultFG && c <= DefaultCursor
}

// IsIndexed returns true if this is an indexed color (0-255).
func (c Color) IsIndexed() bool {
	return c < 256
}

// IsRGB returns true if this is an RGB color (not indexed and not default).
func (c Color) IsRGB() bool {
	return c < DefaultFG && c >= 256
}

// RGB returns the red, green, blue components if this is an RGB color.
// For non-RGB colors, returns 0, 0, 0.
func (c Color) RGB() (r, g, b uint8) {
	if !c.IsRGB() {
		return 0, 0, 0
	}
	return uint8((c >> 16) & 0xFF), uint8((c >> 8) & 0xFF), uint8(c & 0xFF)
}

// Index returns the color index if this is an indexed color.
// For non-indexed colors, returns 0.
func (c Color) Index() uint8 {
	if c.IsIndexed() {
		return uint8(c)
	}
	return 0
}
