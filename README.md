# Conway's Game of Life Playground

#### Video Demo:

#### Description:
Go playground for Conway's famous Game of Life

This programis a cellular automata specifically a B3/S23 variant, based on John Conway's Game of Life first published
in Scientific American by Marvin Gardener in 1970. Ever since, this algorithm and its derivatives have set the stage
for continual development in the field.

To this day thousands of people world wide dedicate large parts of their lives researching new patterns and objects
that emerge in this system and similar ones.

###### Basic setup
The domain or world for this program is a 2D surface which enumerates any number of rows and columns.
Each cell present in the world has its own location which can be addressed by its x, y coordinates.
Every cell can be evaluated to find its "neighbors", or proximally adjacent cells as per the
[Moore neighborhood algorith](https://en.wikipedia.org/wiki/Moore_neighborhood).
Each iteration of the game requires the calculation for each cell, which determines the rendering for the next generation.
How each cell behaves is entirely algorithmic and described in the following rules
<img src="https://user-images.githubusercontent.com/2376084/145755988-9806cfa4-724d-40ce-aaff-ef750ba5864b.png" width="200" />

###### Rules (B2/S23)

1. Any cell that has less than 2 neighbors dies
2. Any live cell having 2 or 3 neighbors survives to the next generation
3. Any cell having more than 3 neighbors dies from overpopulation
4. Any dead cell which has exactly 3 neighbors will come back online

###### Program-specfic variations

This program implementation has some unique features which differntiate it from the classic version.

- Rain
Users can optionally choose to add some stochastic progression, by defining the amount and interval for particles added to the board. 
For example:  `-rain=20 -interval=10` will add 20 random cells every 10 generation frames.

- Initial seed
Users can optionally select to begin the program using a predefined seed

- Initial random cells

Here are the full list of runtime options which can be set:
```
$ ./conways-playground -h
Usage of ./conways-playground:
  -delay int
    	Milliseconds per frame delay (default 200)
  -init int
    	Initial random population (default 7000)
  -interval int
    	Generate rain every Cycle % X == 0 frames (default 1)
  -rain int
    	Generate random rain each iteration, set number of drops here
  -scale string
    	Scale for board: macro, large, default, small, micro (default "default")
  -seed string
    	A known seed: blinker, toad, glider
  -sharp
    	Sharp mode disables rendering of tiles with less than 2 neighbors
```

#### Example commands
```
./conways-playground -seed=copper -init=0 -delay=100 -rain=0 -interval=0 -scale=large
./conways-playground -seed=gosper -init=0 -delay=100 -rain=0 -interval=0 -scale=small

./conways-playground -seed=glider -init=10000 -delay=100 -rain=0 -interval=50 -scale=small -sharp
./conways-playground -seed=glider -init=300 -delay=200 -rain=20 -interval=10 -scale=macro
./conways-playground -seed=glider -init=10000 -delay=100 -rain=10 -interval=50 -scale=small -sharp
./conways-playground -seed=glider -init=53000 -delay=50 -rain=20 -scale=micro -sharp
```

#### Screenshots
![conway](https://user-images.githubusercontent.com/2376084/145700265-78212588-cbfe-4f20-863a-1c0d54d3aab6.png)
![conway2](https://user-images.githubusercontent.com/2376084/145700266-db68d42b-82a6-489b-95aa-9ede801b6c62.png)
![conway3](https://user-images.githubusercontent.com/2376084/145700267-abcc4433-e47e-4a04-a496-b25ed4f86034.png)
