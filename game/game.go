package game

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
	grainWidth   = 10
)

type Game struct {
	gridCols int
	gridRows int
	grid     [][]int
	timer    *Timer
}

func checkIfInBound(ri int, ci int, totalRows int, totalCols int) bool {
	return ri >= 0 && ci >= 0 && ri < totalRows && ci < totalCols
}

func newGrid(gridRows int, gridCols int) [][]int {
	grid := make([][]int, gridRows)

	for i := range grid {
		grid[i] = make([]int, gridCols)
	}

	return grid
}

func (g *Game) Update() error {
	// Check if the left mouse button is pressed
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		// Handle mouse click event here
		mouseX, mouseY := ebiten.CursorPosition()

		row := mouseY / grainWidth
		col := mouseX / grainWidth

		// Within a matrix around row, col as the center
		// create a sand with x probability at any position within
		// the matrix

		matrixWidth := 5
		matrixRadius := matrixWidth / 2

		for i := -matrixRadius; i <= matrixRadius; i++ {
			for j := -matrixRadius; j <= matrixRadius; j++ {
				if rand.Float64() <= 0.10 {
					nRow := row + i
					nCol := col + j
					if checkIfInBound(nRow, nCol, g.gridRows, g.gridCols) {
						g.grid[nRow][nCol] = 1
					}
				}
			}
		}
	}

	// Uncomment to control sand speed

	// g.timer.Update()

	// if !g.timer.isReady() {
	// 	return nil
	// }

	// g.timer.Reset()

	newGrid := newGrid(g.gridRows, g.gridCols)

	for i, r := range g.grid {
		for j, c := range r {
			if c == 1 {
				dir := 1

				// randomly fall left or right
				if rand.Float64() < 0.50 {
					dir = -1
				}

				nextRow := i + 1

				nextColA := j + dir
				nextColB := j - dir
				nextColC := j

				if checkIfInBound(nextRow, nextColC, g.gridRows, g.gridCols) && g.grid[nextRow][nextColC] != 1 {
					// just below
					newGrid[nextRow][nextColC] = 1
				} else if checkIfInBound(nextRow, nextColA, g.gridRows, g.gridCols) && g.grid[nextRow][nextColA] != 1 {
					// below left or right
					newGrid[nextRow][nextColA] = 1
				} else if checkIfInBound(nextRow, nextColB, g.gridRows, g.gridCols) && g.grid[nextRow][nextColB] != 1 {
					// below left or right
					newGrid[nextRow][nextColB] = 1
				} else {
					newGrid[i][j] = 1
				}
			}
		}
	}
	g.grid = newGrid

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i, r := range g.grid {
		for j, c := range r {
			if c == 1 {
				ebitenutil.DrawRect(
					screen, float64(grainWidth*j), float64(grainWidth*i), grainWidth, grainWidth, color.RGBA{R: 194, G: 178, B: 128, A: 255})
			}
		}
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() *Game {
	gridRows := ScreenHeight / grainWidth
	gridCols := ScreenWidth / grainWidth

	grid := newGrid(gridRows, gridCols)

	g := &Game{
		grid:     grid,
		gridRows: gridRows,
		gridCols: gridCols,
		timer:    NewTimer(100 * time.Millisecond),
	}

	return g
}
