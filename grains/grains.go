package grains

import "errors"

func createBoard() []uint64 {
	boardSlice := make([]uint64, 64)
	doubler := 1
	for i := 0; i < len(boardSlice); i++ {
		boardSlice[i] = uint64(doubler)
		doubler *= 2
	}
	return boardSlice
}

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("invalid tile number")
	}

	boardSlice := createBoard()
	return boardSlice[number-1], nil
}

func Total() uint64 {
	boardSlice := createBoard()
	sum := uint64(0)
	for _, v := range boardSlice {
		sum += v
	}
	return sum
}
