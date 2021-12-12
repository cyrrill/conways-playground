package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Ebiten library Game
type Game struct{}

// Each iteration of the game has its data updated here
// evolve() calls the Conway algorithm for next generation
func (g *Game) Update() error {
	time.Sleep(time.Duration(delay) * time.Millisecond)
	board = evolve()
	return nil
}

func evolve() [][]int {
	nextGen := initBoard()
	for i := 1; i < (B - 1); i++ {
		for j := 1; j < (B - 1); j++ {
			n := getNeighbors(i, j)
			if (board[i][j] == 1) && (n < 2) {
				nextGen[i][j] = 0
			} else if (board[i][j] == 1) && (n > 3) {
				nextGen[i][j] = 0
			} else  if (board[i][j] == 0) && (n == 3) {
				nextGen[i][j] = 1
			} else {
				nextGen[i][j] = board[i][j]
			}
		}
	}
	return nextGen
}

// Each iteration get rendered to screen here using Ebiten
func (g *Game) Draw(screen *ebiten.Image) {

	// Render background
	screen.DrawImage(background, nil)

	gen++
	// ebitenutil.DebugPrint(screen, strconv.Itoa(gen))

	totalDots := 0

	// Add random bit for extra noise
	if rain > 0 && gen % rainInterval == 0 {
		randomizeSeed(rain)
	}

	for j := Vi; j < Vm; j++ {
		for i := Vi; i < Vm; i++ {
			var v int
			if board[j][i] == 1 {
				p := float64(rand.Intn(16) + 1)
				v =  int(math.Pow(2, p))
				totalDots++
			} else {
				v = 0
				continue
			}
			if sharp && getNeighbors(j, i) == 0 {
				continue
			}
			// https://github.com/hajimehoshi/ebiten/blob/main/examples/2048/2048/tile.go
			op := &ebiten.DrawImageOptions{}
			x := (i - Vi) * tileSize + (i - Vi + 1) * tileMargin
			y := (j - Vi) * tileSize + (j - Vi + 1) * tileMargin
			op.GeoM.Translate(float64(x), float64(y))
			r, g, b, a := colorToScale(tileBackgroundColor(v))
			op.ColorM.Scale(r, g, b, a)
			screen.DrawImage(tileImage, op)
		}
	}

	// Play note based on board state
	algoSound(totalDots)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return S, S
}