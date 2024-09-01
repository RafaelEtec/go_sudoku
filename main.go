package main

import (
	"fmt"
	"image/color"
	"log"

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
)

type Game struct {
	board_solved   *Board
	board_unsolved *Board
	//boardTile      [][]*BoardTile
}

type Board struct {
	tile          [ROWS][COLUMNS]int
	rows, columns int
}

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

	game := Game{}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) LaunchGame() {
	if !STARTED {
		g.board_solved.rows = ROWS
		g.board_solved.columns = COLUMNS

		for i := 0; i < g.board_solved.rows; i++ {
			for j := 0; j < g.board_solved.columns; j++ {
				g.board_solved.tile[i][j] = 0
			}
		}
		fmt.Println("Game Started")
	}
	STARTED = true
}
