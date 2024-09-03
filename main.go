package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand/v2"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 256
	ScreenHeight = 240

	ROWS    = 9
	COLUMNS = 9
)

var (
	STARTED = false
	BOARD   = [ROWS][COLUMNS]int{}
)

type Game struct {
	//board_solved *Board
	//board_unsolved *Board
	//boardTile      [][]*BoardTile
}

// type Board struct {
// 	tile [][]int
// }

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	//g.input.Update()
	// if err := g.board.checkMove(&g.input.move); err != nil {
	// 	return err
	// }

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{180, 180, 180, 255})

	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(0, 0)

	// screen.DrawImage(
	// 	g.board.Img.SubImage(
	// 		image.Rect(0, 0, 16, 16),
	// 	).(*ebiten.Image),
	// 	&opts,
	// )
}

func main() {
	ebiten.SetWindowSize(ScreenWidth*2, ScreenHeight*2)
	ebiten.SetWindowTitle("SUDOKU by Rafael")

	// boardTileBlank, _, err := ebitenutil.NewImageFromFile("assets/images/numbers.png")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	if !STARTED {
		//boardFillTest(&Game{})
		printBoard()
		boardShuffleTest()
		printBoard()

		STARTED = true
	}

	game := Game{}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

// func boardFillTest(g *Game) {
// 	if !STARTED {
// 		fmt.Println("1 - Starting Game\n2 - Filling board")
// 		for i := 0; i < ROWS; i++ {
// 			for j := 0; j < COLUMNS; j++ {
// 				BOARD[i][j] = 0
// 			}
// 		}
// 	}
// }

func boardShuffleTest() {
	if !STARTED {
		var num int
		for i := 0; i < ROWS; i++ {
			for j := 0; j < COLUMNS; j++ {
				num = rand.IntN(9) + 1
				for appearsInRow(num, i, j) && appearsInCollumn(num, j, i) {
					num = rand.IntN(9) + 1
				}
				BOARD[i][j] = num
			}
			printBoard()
		}
	}
}

func printBoard() {
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLUMNS; j++ {
			if j == ROWS-1 {
				fmt.Print(BOARD[i][j])
			} else {
				fmt.Print(BOARD[i][j], "-")
			}
			time.Sleep(time.Millisecond * 1)
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func appearsInRow(num int, row int, column int) bool {
	for i := 0; i < column; i++ {
		if BOARD[row][i] == num {
			return true
		}
	}
	return false
}

func appearsInCollumn(num int, column int, row int) bool {
	for i := 0; i < row; i++ {
		if BOARD[i][column] == num {
			return true
		}
	}
	return false
}
