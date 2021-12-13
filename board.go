package main

type Scale struct {
	B          int
	tileSize   int
	tileMargin int
}

var scales = map[string]Scale{
	"macro": {
		B:          50,
		tileSize:   35,
		tileMargin: 5,
	},
	"large": {
		B:          100,
		tileSize:   15,
		tileMargin: 5,
	},
	"default": {
		B:          200,
		tileSize:   7,
		tileMargin: 3,
	},
	"small": {
		B:          333,
		tileSize:   5,
		tileMargin: 1,
	},
	"micro": {
		B:          500,
		tileSize:   3,
		tileMargin: 1,
	},
}

func initBoard() [][]int {
	b := make([][]int, B)
	for i := range b {
		b[i] = make([]int, B)
		for j := range b[i] {
			b[i][j] = 0
		}
	}
	return b
}

func getNeighbors(y int, x int) int {
	neighbors := 0
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			neighbors += board[i][j]
		}
	}
	neighbors -= board[y][x]
	return neighbors
}

func evolve() [][]int {
	nextGen := initBoard()
	for i := 1; i < (B - 1); i++ {
		for j := 1; j < (B - 1); j++ {
			n := getNeighbors(i, j)
			if (board[i][j] == 1) && (n < minNeighbors) {
				nextGen[i][j] = 0
			} else if (board[i][j] == 1) && (n > maxNeighbors) {
				nextGen[i][j] = 0
			} else if (board[i][j] == 0) && (n == spawnsNew) {
				nextGen[i][j] = 1
			} else {
				nextGen[i][j] = board[i][j]
			}
		}
	}
	return nextGen
}
