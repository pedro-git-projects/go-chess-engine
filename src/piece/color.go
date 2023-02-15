package piece

// type Color represents the chess piece colors
type Color int

// White and Black are enums representing chess piece colors
const (
	White Color = iota
	Black
)

// ColorString returns the corresponding color as a string
func ColorString(color Color) string {
	switch color {
	case White:
		return "white"
	case Black:
		return "black"
	default:
		return ""
	}
}
