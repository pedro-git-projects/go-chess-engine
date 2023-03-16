package piece_test

import (
	"reflect"
	"testing"

	"github.com/pedro-git-projects/go-chess/src/board"
	"github.com/pedro-git-projects/go-chess/src/utils"
)

func TestKnightCalculateLegalMovements(t *testing.T) {
	b := board.New()
	knight := b.PieceAt(utils.NewCoordinate('b', 1))
	knight.CalculateLegalMoves(b)
	got := knight.LegalMoves()
	expect := []utils.Coordinate{
		utils.NewCoordinate('c', 3),
		utils.NewCoordinate('a', 3),
	}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expected %v but got %v", expect, got)
	}

	b.MovePiece(utils.NewCoordinate('b', 1), utils.NewCoordinate('a', 3))
	knight = b.PieceAt(utils.NewCoordinate('a', 3))
	knight.CalculateLegalMoves(b)
	got = knight.LegalMoves()
	expect = []utils.Coordinate{
		utils.NewCoordinate('b', 5),
		utils.NewCoordinate('b', 1),
		utils.NewCoordinate('c', 4),
	}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expected %v but got %v", expect, got)
	}
}
