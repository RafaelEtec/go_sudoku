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

	board [ROWS][COLUMNS]int
)

type Game struct {
	board_solved *Board
	//board_unsolved *Board
	//boardTile      [][]*BoardTile
}

type Board struct {
	tile [][]int
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	//LaunchGame(g)
	if !STARTED {
		board_createTEST(g)
		board_shuffleTEST(g)

		STARTED = true
	}
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

	game := Game{
		board_solved: &Board{},
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

func LaunchGame(g *Game) {
	if !STARTED {
		for i := 0; i < ROWS; i++ {
			for j := 0; j < COLUMNS; j++ {
				g.board_solved.tile[i][j] = 0
				fmt.Print(g.board_solved.tile[i][j], "-")
				time.Sleep(time.Millisecond * 10)
			}
			fmt.Println("")
		}
		fmt.Println("Game Started")
		STARTED = true
	}
}

func board_createTEST(g *Game) {
	if !STARTED {
		fmt.Println("1 - Starting Game")
		for i := 0; i < ROWS; i++ {
			for j := 0; j < COLUMNS; j++ {
				board[i][j] = 0
			}
		}
		printBoard()
	}
}

func board_shuffleTEST(g *Game) {
	if !STARTED {
		for i := 0; i < ROWS; i++ {
			chosen := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
			for j := 0; j < COLUMNS; j++ {
				num := rand.IntN(9) + 1
				for wasChosen(num, chosen[:]) {
					num = rand.IntN(9) + 1
				}
				chosen[i] = num
			}

			for k := 0; k < COLUMNS; k++ {
				board[i][k] = chosen[k]
			}
		}
		printBoard()
	}
}

func printBoard() {
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLUMNS; j++ {
			board[i][j] = 0
			if j == ROWS-1 {
				fmt.Print(board[i][j])
			} else {
				fmt.Print(board[i][j], "-")
			}
			time.Sleep(time.Millisecond * 10)
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func wasChosen(num int, chosen []int) bool {
	for i := 0; i < len(chosen); i++ {
		if chosen[i] == num {
			return true
		}
	}
	return false
}
