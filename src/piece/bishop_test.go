package piece_test

import (
	"reflect"
	"testing"

	"github.com/pedro-git-projects/go-chess/src/board"
	"github.com/pedro-git-projects/go-chess/src/utils"
)

// TODO: write more tests to check legal moves
func TestBishopCalculateLegalMoves(t *testing.T) {
	b := board.New()
	bishop := b.PieceAt(utils.NewCoordinate('c', 1))
	bishop.CalculateLegalMoves(b)
	got := bishop.LegalMoves()
	expect := []utils.Coordinate{}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expected empty but got %v", got)
	}
}

// TODO: actually caputre a piece
func TestBishopCapture(t *testing.T) {
	b := board.New()
	b.MovePiece(utils.NewCoordinate('d', 2), utils.NewCoordinate('d', 4))
	b.MovePiece(utils.NewCoordinate('c', 1), utils.NewCoordinate('d', 2))
	bishop := b.PieceAt(utils.NewCoordinate('d', 2))
	bishop.CalculateLegalMoves(b)
	got := bishop.LegalMoves()
	expect := []utils.Coordinate{
		utils.NewCoordinate('e', 3),
		utils.NewCoordinate('f', 4),
		utils.NewCoordinate('g', 5),
		utils.NewCoordinate('h', 6),
		utils.NewCoordinate('c', 3),
		utils.NewCoordinate('b', 4),
		utils.NewCoordinate('a', 5),
		utils.NewCoordinate('c', 1),
	}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expected %v but got %v\n", expect, got)
	}
}
