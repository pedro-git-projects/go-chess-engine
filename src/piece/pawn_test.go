package piece_test

import (
	"reflect"
	"testing"

	"github.com/pedro-git-projects/go-chess/src/board"
	"github.com/pedro-git-projects/go-chess/src/utils"
)

func TestMovement(t *testing.T) {
	b := board.New()
	state := b.State()

	a4 := state[utils.NewCoordinate('a', 4)]

	pawn := b.PieceAt(utils.NewCoordinate('a', 2))
	pawn.Move(utils.NewCoordinate('a', 4), b)

	newstate := b.State()
	newa4 := newstate[utils.NewCoordinate('a', 4)]
	newa2 := newstate[utils.NewCoordinate('a', 2)]

	if reflect.DeepEqual(newa4, a4) {
		t.Errorf("current: %v and original: %v should not be equal", newa4, a4)
	}

	if reflect.DeepEqual(newa4, nil) {
		t.Error("piece did not move")
	}

	if !reflect.DeepEqual(newa2, nil) {
		t.Error("piece is occuping two spots")
	}
}

func TestCalculateLegalMoves(t *testing.T) {
	b := board.New()
	pawn := b.PieceAt(utils.NewCoordinate('a', 2))
	pawn.CalculateLegalMoves(b)
	got := pawn.LegalMoves()

	expect := []utils.Coordinate{
		utils.NewCoordinate('a', 3),
		utils.NewCoordinate('a', 4),
	}

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expected %v but got %v", expect, got)
	}
}
