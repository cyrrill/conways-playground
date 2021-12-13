package main

// Conway's Game of Life Playground
// ********************************
//  https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life
//
//	Sample executions to try:
//  *************************
// ./conways-playground -init=40000 -delay=100 -rain=0 -scale=micro
// ./conways-playground -init=10000 -delay=0 -rain=0 -scale=small
// ./conways-playground -seed=copper -init=0 -delay=100 -rain=0 -interval=0 -scale=large
// ./conways-playground -seed=gosper -init=0 -delay=100 -rain=0 -interval=0 -scale=small
// ./conways-playground -seed=glider -init=10000 -delay=100 -rain=0 -interval=50 -scale=small -sharp
// ./conways-playground -seed=glider -init=300 -delay=200 -rain=20 -interval=10 -scale=macro
// ./conways-playground -seed=glider -init=10000 -delay=100 -rain=10 -interval=50 -scale=small -sharp
// ./conways-playground -seed=glider -init=53000 -delay=50 -rain=20 -scale=micro -sharp

import (
	"image/color"
	_ "image/png"
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

// Runtime variables
var (
	// Game runtime
	game *ConwaysGame

	// The matrix
	board [][]int

	// Size of board (BxB)
	B = 200

	// Visible board
	Vi = 50
	Vm = 150

	// B3/S23 rule
	minNeighbors = 2
	maxNeighbors = 3
	spawnsNew    = 3

	// Tile rendering
	tileImage  *ebiten.Image
	tileSize   = 8
	tileMargin = 2
	sharp      = false

	// Initial population size
	gen = 0

	// Random bits introduced every per interval frame
	rain         = 0
	rainInterval = 1

	// Speed of rendering
	delay = 200

	// Loads image from file
	background *ebiten.Image
	err        error

	sound = false
)

// Initialize game state
func init() {

	game = new(ConwaysGame)

	// Random number generator needs initialization
	rand.Seed(time.Now().UnixNano())

	// Parse imput flags
	readInputs()

	// Setup board with 0 bits everywhere
	board = initBoard()

	// Load known seed shape at start
	setSeed(seed)

	// Create randomized generation 0
	randomizeSeed(randomSeed)

	// Set each tile image to just be plain white to start
	tileImage = ebiten.NewImage(tileSize, tileSize)
	tileImage.Fill(color.White)

	// Load background image
	background, _, _ = ebitenutil.NewImageFromFile("background.png")

	// Load sound context
	initSound()
}

func main() {
	game.Start()
}
