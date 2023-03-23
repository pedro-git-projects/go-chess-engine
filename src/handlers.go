package main

import (
	"fmt"
	"net/http"

	"github.com/pedro-git-projects/go-chess/src/board"
)

func HandleBoard(w http.ResponseWriter, r *http.Request) {
	b := board.New()
	w.Header().Add("Content-Type", "application/json")
	m := b.Marshal()
	fmt.Fprintln(w, m)
}
