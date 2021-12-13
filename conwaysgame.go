package main

import (
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game controller
type ConwaysGame struct{}

// Initialize window and run Ebiten game runtime
func (g *ConwaysGame) Start() {
	ebiten.SetWindowSize(S, S)
	ebiten.SetWindowTitle("Conway's Game of Life Playground")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

// Each iteration of the game has its data updated here
// evolve() calls the Conway algorithm for next generation
func (g *ConwaysGame) Update() error {
	time.Sleep(time.Duration(delay) * time.Millisecond)
	board = evolve()
	return nil
}

// Each iteration get rendered to screen here using Ebiten
func (g *ConwaysGame) Draw(screen *ebiten.Image) {

	// Render background
	screen.DrawImage(background, nil)

	// Track generations
	gen++
	ebitenutil.DebugPrint(screen, strconv.Itoa(gen))

	// Population: keeps track of total bits rendered
	totalDots := 0

	// Add random bit for extra noise
	if rain > 0 && gen%rainInterval == 0 {
		randomizeSeed(rain)
	}

	// Iterate through all points, calculate neighbors, color and render to image
	for j := Vi; j < Vm; j++ {
		for i := Vi; i < Vm; i++ {
			var v int
			if board[j][i] == 1 {
				// Set randomized color
				v = int(math.Pow(2, float64(rand.Intn(16)+1)))
				totalDots++
			} else {
				// Skip rendering process for off bits
				v = 0
				continue
			}
			if sharp && getNeighbors(j, i) == 0 {
				continue
			}
			// https://github.com/hajimehoshi/ebiten/blob/main/examples/2048/2048/tile.go
			op := &ebiten.DrawImageOptions{}
			x := (i-Vi)*tileSize + (i-Vi+1)*tileMargin
			y := (j-Vi)*tileSize + (j-Vi+1)*tileMargin
			op.GeoM.Translate(float64(x), float64(y))
			r, g, b, a := colorToScale(tileBackgroundColor(v))
			op.ColorM.Scale(r, g, b, a)
			screen.DrawImage(tileImage, op)
		}
	}

	// Play note based on board state
	algoSound(totalDots)
}

func (g *ConwaysGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return S, S
}
