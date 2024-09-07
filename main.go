package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	SCREEN_WIDTH  = 256
	SCREEN_HEIGHT = 240

	FRAME_OX     = 0
	FRAME_OY     = 0
	FRAME_WIDTH  = 32
	FRAME_HEIGHT = 32

	ROWS    = 9
	COLUMNS = 9
)

const (
	t0 = 0
	t1 = iota
	t2
	t3
	t4
	t5
	t6
	t7
	t8
	t9
)

var (
	STARTED = false
)

type Game struct {
	moves  int
	misses int
	state  int
	board  *Board
}

type Board struct {
	tiles [][]*Tile
}

type Tile struct {
	Img   *ebiten.Image
	Value int
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func (g *Game) Update() error {
	// TODO
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{180, 180, 180, 255})

	opts := ebiten.DrawImageOptions{}

	drawTiles(g, opts, screen)
	//drawStats(g, screen)
}

func drawTiles(g *Game, opts ebiten.DrawImageOptions, screen *ebiten.Image) {
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLUMNS; c++ {
			tile := g.board.tiles[r][c]

			fox, foy, fw, fh := FRAME_OX, FRAME_OY, FRAME_WIDTH, FRAME_HEIGHT*2
			switch tile.Value {
			case 1:
				foy += 64
				fw *= 2
				fh *= 2
			case 2:
				foy += 64
				fw *= 2
				fh *= 2
			default:
			}

			opts.GeoM.Translate(float64(r)*32, float64(c)*32)
			screen.DrawImage(
				tile.Img.SubImage(
					image.Rect(fox, foy, fw, fh),
				).(*ebiten.Image),
				&opts,
			)
		}
	}
}

func main() {

	tile, _, err := ebitenutil.NewImageFromFile("assets/images/numbers.png")
	if err != nil {
		log.Fatal(err)
	}

	game := Game{
		moves:  0,
		misses: 0,
		state:  0,
		board: &Board{
			tiles: [][]*Tile{
				{
					{
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					},
				}, {
					{
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					},
				}, {
					{
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					},
				}, {
					{
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					},
				}, {
					{
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					},
				}, {
					{
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					},
				}, {
					{
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					},
				}, {
					{
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					},
				}, {
					{
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					}, {
						Img:   tile,
						Value: 0,
					},
				},
			},
		},
	}

	ebiten.SetWindowSize(SCREEN_WIDTH*2, SCREEN_HEIGHT*2)
	ebiten.SetWindowTitle("SUDOKU by Rafael Goulart")
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

// func boardShuffleTest() {
// 	if !STARTED {
// 		var num int
// 		for i := 0; i < ROWS; i++ {
// 			for j := 0; j < COLUMNS; j++ {
// 				num = rand.IntN(9) + 1
// 				for appearsInRow(num, i, j) || appearsInCollumn(num, j, i) {
// 					num = rand.IntN(9) + 1
// 				}
// 				BOARD[i][j] = num
// 			}
// 		}
// 	}
// }

// func printBoard() {
// 	for i := 0; i < ROWS; i++ {
// 		for j := 0; j < COLUMNS; j++ {
// 			if j == ROWS-1 {
// 				fmt.Print(BOARD[i][j])
// 			} else {
// 				fmt.Print(BOARD[i][j], "-")
// 			}
// 			time.Sleep(time.Millisecond * 1)
// 		}
// 		fmt.Println("")
// 	}
// 	fmt.Println("")
// }

// func appearsInRow(num int, row int, column int) bool {
// 	for i := 0; i < column; i++ {
// 		if BOARD[row][i] == num {
// 			return true
// 		}
// 	}
// 	return false
// }

// func appearsInCollumn(num int, column int, row int) bool {
// 	for i := 0; i < row; i++ {
// 		if BOARD[i][column] == num {
// 			return true
// 		}
// 	}
// 	return false
// }
