package game

// type turn represents the current chess piece color turn
type turn int

// White and Black are enums representing chess piece color's turn
const (
	None turn = iota
	White
	Black
)

func (turn turn) String() string {
	switch turn {
	case White:
		return "white's turn"
	case Black:
		return "black's turn"
	default:
		return ""
	}
}
