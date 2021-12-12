package main

var scales = map[string]Scale{
	"macro": {
		B: 50,
		tileSize: 35,
		tileMargin: 5,
	},
	"large": {
		B: 100,
		tileSize: 15,
		tileMargin: 5,
	},
	"default": {
		B: 200,
		tileSize: 7,
		tileMargin: 3,
	},
	"small": {
		B: 333,
		tileSize: 5,
		tileMargin: 1,
	},
	"micro": {
		B: 500,
		tileSize: 3,
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