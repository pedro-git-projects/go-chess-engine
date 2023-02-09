package board_test

import (
	"reflect"
	"testing"

	"github.com/pedro-git-projects/go-chess/src/board"
	"github.com/pedro-git-projects/go-chess/src/utils"
)

func TestFind(t *testing.T) {
	board := board.New()
	got := board.Find(utils.NewCoordinate("a", 3))
	expect := true

	if got != expect {
		t.Errorf("expected %t but got %t", expect, got)
	}

	got = board.Find(utils.NewCoordinate("z", 1))
	expect = false

	if got != expect {
		t.Errorf("expected %t but got %t", expect, got)
	}
}

func TestFilterByCol(t *testing.T) {
	board := board.New()
	got := board.FilterByCol("a")
	expect := []utils.Coordinate{
		utils.NewCoordinate("a", 8),
		utils.NewCoordinate("a", 7),
		utils.NewCoordinate("a", 6),
		utils.NewCoordinate("a", 5),
		utils.NewCoordinate("a", 4),
		utils.NewCoordinate("a", 3),
		utils.NewCoordinate("a", 2),
		utils.NewCoordinate("a", 1),
	}

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expected %v but got %v", expect, got)
	}
}

func TestFowardRightDiagonal(t *testing.T) {
	board := board.New()
	got := board.FowardRightDiagonal(utils.NewCoordinate("d", 1))
	expect := []utils.Coordinate{
		utils.NewCoordinate("e", 2),
		utils.NewCoordinate("f", 3),
		utils.NewCoordinate("g", 4),
		utils.NewCoordinate("h", 5),
	}

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("expected %v but got %v", expect, got)
	}
}
