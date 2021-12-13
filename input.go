package main

import (
	"flag"
	"fmt"
)

var (
	seed       string
	randomSeed int
)

func readInputs() {

	scaleFlag := flag.String("scale", "default", "Scale for board: macro, large, default, small, micro")

	// Load known shape from seed list
	seedFlag := flag.String("seed", "", "A known seed: blinker, toad, glider")

	// Set how much rain to use
	rainFlag := flag.Int("rain", 0, "Generate random rain each iteration, set number of drops here")
	rainIntervalFlag := flag.Int("interval", 1, "Generate rain every Cycle % X == 0 frames")

	sharpFlag := flag.Bool("sharp", false, "Sharp mode disables rendering of tiles with less than 2 neighbors")

	// Add first generation as random bits
	initFlag := flag.Int("init", GEN1, "Initial random population")

	// Set milliseconds of delay in render
	delayFlag := flag.Int("delay", 200, "Milliseconds per frame delay")

	flag.Parse()
	fmt.Printf("Seed: %s, Rain: %d, Delay: %d, Init: %d\n", *seedFlag, *rainFlag, *delayFlag, *initFlag)

	scale := scales[*scaleFlag]
	B = scale.B
	Bm := B / 4
	Vi = 0 + Bm
	Vm = B - Bm
	tileSize = scale.tileSize
	tileMargin = scale.tileMargin
	fmt.Printf("B: %d, Vi: %d, Vm: %d, tileSize: %d, tileMargin: %d\n", B, Vi, Vm, tileSize, tileMargin)

	randomSeed = *initFlag
	seed = *seedFlag
	rain = *rainFlag
	rainInterval = *rainIntervalFlag
	sharp = *sharpFlag
	delay = *delayFlag
}
