package main

import (
	"fmt"
	"net/http"

	"github.com/pedro-git-projects/go-chess/src/game"
)

func main() {
	g := game.New()
	g.PrintBoard()
	mux := http.NewServeMux()
	mux.HandleFunc("/board", HandleBoard)
	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", mux)
}
