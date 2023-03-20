package piece_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/pedro-git-projects/go-chess/src/board"
	"github.com/pedro-git-projects/go-chess/src/utils"
)

func TestQueenCalculateLegalMovements(t *testing.T) {
	b := board.New()
	b.MovePiece(utils.NewCoordinate('d', 2), utils.NewCoordinate('d', 4))
	b.MovePiece(utils.NewCoordinate('d', 7), utils.NewCoordinate('d', 5))
	queen := b.PieceAt(utils.NewCoordinate('d', 1))
	queen.CalculateLegalMoves(b)
	got := queen.LegalMoves()
	want := []utils.Coordinate{
		utils.NewCoordinate('d', 2),
		utils.NewCoordinate('d', 3),
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v\n", got, want)
	}
	b.MovePiece(utils.NewCoordinate('d', 1), utils.NewCoordinate('d', 3))
	queen = b.PieceAt(utils.NewCoordinate('d', 3))
	queen.CalculateLegalMoves(b)
	fmt.Println(queen.LegalMoves())
}
