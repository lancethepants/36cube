package main

import (
	"fmt"
)

var board [6][6]Tower

var freetowers []Tower = initializeFreetowers()

/*
var board_height = [6][6]int{
	{0, 3, 4, 2, 1, 5},
	{2, 1, 5, 0, 3, 4},
	{5, 4, 2, 3, 0, 1},
	{4, 0, 3, 1, 5, 2},
	{3, 5, 1, 4, 2, 0},
	{1, 2, 0, 5, 4, 3},
}
*/

// tricky tricky puzzle maker

var board_height = [6][6]int{
	{0, 3, 4, 2, 1, 5},
	{2, 1, 5, 0, 3, 4},
	{5, 4, 2, 3, 0, 1},
	{4, 1, 3, 0, 5, 2},
	{3, 5, 1, 4, 2, 0},
	{1, 2, 0, 5, 4, 3},
}

func main() {

	var position = Position{0, 5}

	for len(freetowers) > -1 {

		if position.row == -1 || position.column == -1 {
			break
		}

		position = evaluate_cube(position)

		if len(freetowers) == 0 {
			printBoard()
		}
	}
}

func evaluate_cube(p Position) Position {

	//is already purple
	if board[p.row][p.column].inUse && board[p.row][p.column].color == Color(5) {
		board[p.row][p.column].inUse = false
		freetowers = append(freetowers, board[p.row][p.column])
		return backPosition(p)
	}

	i := 0
	if board[p.row][p.column].inUse {
		i = int(board[p.row][p.column].color) + 1
	}
	// search for available tower
	for ; i < 6; i++ {
		// if found available tower
		if index := findAvailableTower(Tower{6 - board_height[p.row][p.column], Color(i), false}); index > -1 && colorIsFree(p, Color(i)) {

			if board[p.row][p.column].inUse {
				freetowers = append(freetowers, board[p.row][p.column])
			}
			board[p.row][p.column] = Tower{6 - board_height[p.row][p.column], Color(i), true}
			freetowers[index] = freetowers[len(freetowers)-1]
			freetowers = freetowers[:len(freetowers)-1]
			return advancePosition(p)
		}
	}
	// if no available tower found
	if board[p.row][p.column].inUse {
		board[p.row][p.column].inUse = false
		freetowers = append(freetowers, board[p.row][p.column])
	}

	return backPosition(p)
}

func printBoard() {
	// add if statement to check our two special pieces are in their places
	if board[3][1].color == Color(2) && board[3][3].color == Color(1) {
		for i := 0; i < 6; i++ {
			for j := 0; j < 6; j++ {
				if board[i][j].inUse {
					fmt.Printf("%v%d ", board[i][j].color, board[i][j].height)
				} else {
					fmt.Printf("*  ")
				}
			}
			fmt.Println("")
		}
		fmt.Println("")
	}
}

func advancePosition(p Position) Position {

	if p.row == 5 && p.column == 0 {
		return backPosition(p)
	}

	if column := p.column - 1; column < 0 {
		return Position{p.row + 1, 5}
	}
	return Position{p.row, p.column - 1}

}

func backPosition(p Position) Position {

	if column := p.column + 1; column > 5 {
		return Position{p.row - 1, 0}
	}
	return Position{p.row, p.column + 1}
}

func findAvailableTower(t Tower) int {

	for i := 0; i < len(freetowers); i++ {

		if freetowers[i].height == t.height && freetowers[i].color == t.color {
			return i
		}
	}
	return -1
}

type Color int

const (
	Red Color = iota
	Orange
	Yellow
	Green
	Blue
	Purple
)

func (c Color) String() string {
	return [...]string{"R", "O", "Y", "G", "B", "P"}[c]
}

type Tower struct {
	height int
	color  Color
	inUse  bool
}

type Position struct {
	row    int
	column int
}

func initializeFreetowers() []Tower {

	var towers []Tower

	for i := 1; i < 7; i++ {
		for j := 0; j < 6; j++ {
			towers = append(towers, Tower{i, Color(j), false})
		}
	}
	return towers
}

func colorIsFree(p Position, color Color) bool {

	for i := 0; i < 6; i++ {
		if board[p.row][i].color == color && board[p.row][i].inUse == true {
			return false
		}
	}

	for i := 0; i < 6; i++ {
		if board[i][p.column].color == color && board[i][p.column].inUse == true {
			return false
		}
	}
	return true
}
