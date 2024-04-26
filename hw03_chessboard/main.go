package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	height, width, err := promptChessboardParameters()
	if err != nil {
		fmt.Println(err)
		return
	}

	chessBoard, err := generateChessBoard(height, width)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(chessboardToString(chessBoard))
}

type ChessCell rune

const (
	white   ChessCell = ' '
	black   ChessCell = '#'
	invalid ChessCell = '?'
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
		return invalid, fmt.Errorf("invalid chess cell value: %d", value)
	}
}

func generateChessBoard(height int, width int) ([][]ChessCell, error) {
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
	return acc.String()
}

func promptChessboardParameters() (height int, width int, err error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter chess board height: ")
	inputHeight, err := reader.ReadString('\n')
	if err != nil {
		err = fmt.Errorf("invalid chess board height value: %s", inputHeight)
		return
	}

	height, err = strconv.Atoi(strings.TrimSpace(inputHeight))
	if err != nil {
		err = fmt.Errorf("invalid chess board height value: %s", inputHeight)
		return
	}

	fmt.Print("Enter chess board width: ")
	inputWidth, err := reader.ReadString('\n')
	if err != nil {
		err = fmt.Errorf("invalid chess board width value: %s", inputWidth)
		return
	}

	width, err = strconv.Atoi(strings.TrimSpace(inputWidth))
	if err != nil {
		err = fmt.Errorf("invalid chess board width value: %s", inputWidth)
		return
	}

	return
}
