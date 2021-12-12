package main

// ./conways-playground -seed=copper -init=0 -delay=100 -rain=0 -interval=0 -scale=large
// ./conways-playground -seed=gosper -init=0 -delay=100 -rain=0 -interval=0 -scale=small

// ./conways-playground -seed=glider -init=10000 -delay=100 -rain=0 -interval=50 -scale=small -sharp
// ./conways-playground -seed=glider -init=300 -delay=200 -rain=20 -interval=10 -scale=macro
// ./conways-playground -seed=glider -init=10000 -delay=100 -rain=10 -interval=50 -scale=small -sharp
// ./conways-playground -seed=glider -init=53000 -delay=50 -rain=20 -scale=micro -sharp

import (
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Initial constants which define program execution
const (
	// Window viewport size (SxS)
	S = 1004

	// Number of random bits in generation 1
	GEN1 = 7000
)

type Scale struct {
	B int
	tileSize int
	tileMargin int
}

// Runtime variables
var (
	// The matrix
	board [][]int

	// Size of board (BxB)
	B = 200

	// Visible board
	Vi = 50
	Vm = 150

	tileImage *ebiten.Image
	tileSize = 8
	tileMargin = 2
	sharp = false

	// Initial population size
	gen = 0

	err error
	rain = 0
	rainInterval = 1
	delay = 200
    background *ebiten.Image
)

// Initialize game state
func init() {
	// Random number generator needs initialization
	rand.Seed(time.Now().UnixNano())

	// Parse imput flags
	readInputs()

	// Setup board with 0 bits everywhere
	board = initBoard()

	setSeed(seed)
	randomizeSeed(randomSeed)

	// Set each tile image to just be plain white to start
	tileImage = ebiten.NewImage(tileSize, tileSize)
	tileImage.Fill(color.White)

	// Load background image
	background, _, _ = ebitenutil.NewImageFromFile("background.png")

	// Load sound context
	initSound()
}


func getNeighbors(y int, x int) int {
	neighbors := 0
	for i := y - 1; i <= y + 1; i++ {
		for j := x - 1; j <= x + 1; j++ {
			neighbors += board[i][j]
		}
	}
	neighbors -= board[y][x]
	return neighbors
}

func main() {

	game := &Game{}

	ebiten.SetWindowSize(S, S)
	ebiten.SetWindowTitle("Conway's Game of Life Playground")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

