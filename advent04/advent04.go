package advent04

import (
	"2021-advent/util"
	"fmt"
	"path/filepath"
	"runtime"
)

func RunPuzzle() {
	fmt.Println("ho ho ho")
	_, pwd, _, _ := runtime.Caller(0)
	println(pwd)
	puzzlePath := filepath.Join(pwd, "..", "04_puzzle.txt")
	fmt.Println(puzzlePath)
	game := util.ReadFile(puzzlePath)
	winner_index, score := RunGame(game)
	fmt.Printf("Winner Card %d with Score %d\n", winner_index, score)
	winner_index, score = FindLoser(game)
	fmt.Printf("Loser Card %d with Score %d\n", winner_index, score)
}
