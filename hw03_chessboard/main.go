package main

import (
	"fmt"
	"strings"
)

func main() {
	chessBoard, err := generateChessBoard(8, 8)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(chessboardToString(chessBoard))
}

type ChessCell string

const (
	white ChessCell = "  "
	black ChessCell = "#"
)

func (c ChessCell) String() string {
	return string(c)
}

func getChessCellValue(value int) (ChessCell, error) {
	switch value {
	case 0:
		return white, nil
	case 1:
		return black, nil
	default:
		return "", fmt.Errorf("invalid chess cell value: %d", value)
	}
}

func generateChessBoard(width int, height int) ([][]ChessCell, error) {
	chessBoard := make([][]ChessCell, height)
	for i := range chessBoard {
		chessBoard[i] = make([]ChessCell, width)
	}

	for rowIdx := range chessBoard {
		for cellIdx := range chessBoard[rowIdx] {
			cell, err := getChessCellValue((rowIdx + cellIdx) % 2)
			if err != nil {
				return nil, err
			}
			chessBoard[rowIdx][cellIdx] = cell
		}
	}

	return chessBoard, nil
}

func chessboardToString(board [][]ChessCell) string {
	var acc strings.Builder
	for _, row := range board {
		for _, cell := range row {
			acc.WriteString(cell.String())
		}
		acc.WriteString("\n")
	}
	return fmt.Sprint(acc.String())
}
