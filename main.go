package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	SCREEN_WIDTH  = 294
	SCREEN_HEIGHT = 308

	FRAME_OX     = 0
	FRAME_OY     = 0
	FRAME_WIDTH  = 32
	FRAME_HEIGHT = 32

	BOARD_WIDTH  = 288
	BOARD_HEIGHT = 288

	ROWS    = 9
	COLUMNS = 9

	BASE = 3
	SIDE = BASE * BASE

	MESSAGE_MOVES   = "Moves: %d"
	MESSAGE_MISSES  = "Misses: %d"
	MESSAGE_DEFEAT  = "Too many misses! You lose :/"
	MESSAGE_VICTORY = "NICE, You completed the board!"

	STARTING_NUMBERS = 38
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

	X = -1
	Y = -1
)

type Game struct {
	misses  int
	state   int
	message string
	board   *Board
	aux     [][]*Tile
}

// state: 0 = stopped
//        1 = playing
//        2 = won

type Board struct {
	tiles [][]*Tile
}

type Tile struct {
	Img        *ebiten.Image
	Value      int
	isEditable bool
	isRight    bool
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func (g *Game) Update() error {
	if g.state == 1 {
		px, py := handleMouse(g)
		if px != -1 && py != -1 {
			X = px
			Y = py
		}

		handleKeyboard(g, Y, X)
		handleMisses(g)
	}

	options(g)

	return nil
}

func handleKeyboard(g *Game, x int, y int) {
	if x != -1 && y != -1 {
		if inpututil.IsKeyJustReleased(ebiten.Key0) {
			handleBoard(g, x, y, t0)
		} else if inpututil.IsKeyJustReleased(ebiten.Key1) {
			handleBoard(g, x, y, t1)
		} else if inpututil.IsKeyJustPressed(ebiten.Key2) {
			handleBoard(g, x, y, t2)
		} else if inpututil.IsKeyJustPressed(ebiten.Key3) {
			handleBoard(g, x, y, t3)
		} else if inpututil.IsKeyJustPressed(ebiten.Key4) {
			handleBoard(g, x, y, t4)
		} else if inpututil.IsKeyJustPressed(ebiten.Key5) {
			handleBoard(g, x, y, t5)
		} else if inpututil.IsKeyJustPressed(ebiten.Key6) {
			handleBoard(g, x, y, t6)
		} else if inpututil.IsKeyJustPressed(ebiten.Key7) {
			handleBoard(g, x, y, t7)
		} else if inpututil.IsKeyJustPressed(ebiten.Key8) {
			handleBoard(g, x, y, t8)
		} else if inpututil.IsKeyJustPressed(ebiten.Key9) {
			handleBoard(g, x, y, t9)
		}
	}
}

func handleMouse(g *Game) (int, int) {
	px, py := -1, -1
	if inpututil.IsMouseButtonJustPressed(0) {
		x, y := ebiten.CursorPosition()
		px, py = whereWasClicked(x, y)
	}
	return px, py
}

func handleBoard(g *Game, x int, y int, num int) {

	tile, err := ebitenutil.NewImageFromURL("https://github.com/RafaelEtec/go_sudoku/blob/31af9f64952f8e622d8403e2febd2d6bb994a625/assets/images/numbers.png?raw=true")
	if err != nil {
		log.Fatal(err)
	}

	tile_right, err := ebitenutil.NewImageFromURL("https://github.com/RafaelEtec/go_sudoku/blob/31af9f64952f8e622d8403e2febd2d6bb994a625/assets/images/numbers_right.png?raw=true")
	if err != nil {
		log.Fatal(err)
	}

	tile_wrong, err := ebitenutil.NewImageFromURL("https://github.com/RafaelEtec/go_sudoku/blob/31af9f64952f8e622d8403e2febd2d6bb994a625/assets/images/numbers_wrong.png?raw=true")
	if err != nil {
		log.Fatal(err)
	}

	if g.board.tiles[x][y].isEditable {
		if num == 0 {
			g.board.tiles[x][y].Value = 0
			g.board.tiles[x][y].Img = tile
		} else if g.aux[x][y].Value == num {
			g.board.tiles[x][y].Value = num
			g.board.tiles[x][y].Img = tile_right
			g.board.tiles[x][y].isEditable = false
			g.board.tiles[x][y].isRight = true
		} else {
			g.board.tiles[x][y].Value = num
			g.board.tiles[x][y].Img = tile_wrong
			g.misses++
		}

		if handleVictory(g) {
			g.state = 2
			g.message = MESSAGE_VICTORY
		}
	}
}

func highlightRowAndColumn(g *Game, x int, y int) {

}

func hightlightSurroundingNumbers(g *Game, num int) {

}

func handleVictory(g *Game) bool {
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLUMNS; c++ {
			if !g.board.tiles[r][c].isRight {
				return false
			}
		}
	}
	return true
}

func handleMisses(g *Game) {
	if g.misses == 10 {
		g.state = 0
		g.message = MESSAGE_DEFEAT
	}
}

func options(g *Game) {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		restart(g)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		showSolution(g)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		cheat(g)
	}
}

func cheat(g *Game) {
	tile_right, err := ebitenutil.NewImageFromURL("https://github.com/RafaelEtec/go_sudoku/blob/31af9f64952f8e622d8403e2febd2d6bb994a625/assets/images/numbers_right.png?raw=true")
	if err != nil {
		log.Fatal(err)
	}
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLUMNS; c++ {
			if g.board.tiles[r][c].Value == 0 {
				g.board.tiles[r][c].Value = g.aux[r][c].Value
				g.board.tiles[r][c].Img = tile_right
				g.board.tiles[r][c].isEditable = false
				g.board.tiles[r][c].isRight = true
			}
		}
	}

	g.state = 2
	g.message = MESSAGE_VICTORY
}

func restart(g *Game) {
	g.misses = 0
	g.state = 1
	g.message = ""
	createAuxBoard(g)
	createBoard(g)

	fillAuxBoard(g)
	fillBoard(g)

	removeSome(g)
	addStats(g)
}

func showSolution(g *Game) {
	printAuxBoard(g)
}

func whereWasClicked(x int, y int) (int, int) {
	x -= 2
	y -= 2
	if x >= 0 && x <= BOARD_WIDTH && y >= 0 && y <= BOARD_HEIGHT {
		if x >= 0 && x <= FRAME_WIDTH*t1 {
			if y >= 0 && y <= FRAME_HEIGHT*t1 {
				return 0, 0
			} else if y >= FRAME_HEIGHT*t1 && y <= FRAME_HEIGHT*t2 {
				return 0, 1
			} else if y >= FRAME_HEIGHT*t2 && y <= FRAME_HEIGHT*t3 {
				return 0, 2
			} else if y >= FRAME_HEIGHT*t3 && y <= FRAME_HEIGHT*t4 {
				return 0, 3
			} else if y >= FRAME_HEIGHT*t4 && y <= FRAME_HEIGHT*t5 {
				return 0, 4
			} else if y >= FRAME_HEIGHT*t5 && y <= FRAME_HEIGHT*t6 {
				return 0, 5
			} else if y >= FRAME_HEIGHT*t6 && y <= FRAME_HEIGHT*t7 {
				return 0, 6
			} else if y >= FRAME_HEIGHT*t7 && y <= FRAME_HEIGHT*t8 {
				return 0, 7
			} else if y >= FRAME_HEIGHT*t8 && y <= FRAME_HEIGHT*t9 {
				return 0, 8
			}
		}
		if x >= FRAME_WIDTH && x <= FRAME_WIDTH*t2 {
			if y >= 0 && y <= FRAME_HEIGHT*t1 {
				return 1, 0
			} else if y >= FRAME_HEIGHT*t1 && y <= FRAME_HEIGHT*t2 {
				return 1, 1
			} else if y >= FRAME_HEIGHT*t2 && y <= FRAME_HEIGHT*t3 {
				return 1, 2
			} else if y >= FRAME_HEIGHT*t3 && y <= FRAME_HEIGHT*t4 {
				return 1, 3
			} else if y >= FRAME_HEIGHT*t4 && y <= FRAME_HEIGHT*t5 {
				return 1, 4
			} else if y >= FRAME_HEIGHT*t5 && y <= FRAME_HEIGHT*t6 {
				return 1, 5
			} else if y >= FRAME_HEIGHT*t6 && y <= FRAME_HEIGHT*t7 {
				return 1, 6
			} else if y >= FRAME_HEIGHT*t7 && y <= FRAME_HEIGHT*t8 {
				return 1, 7
			} else if y >= FRAME_HEIGHT*t8 && y <= FRAME_HEIGHT*t9 {
				return 1, 8
			}
		}
		if x >= FRAME_WIDTH*t2 && x <= FRAME_WIDTH*t3 {
			if y >= 0 && y <= FRAME_HEIGHT*t1 {
				return 2, 0
			} else if y >= FRAME_HEIGHT*t1 && y <= FRAME_HEIGHT*t2 {
				return 2, 1
			} else if y >= FRAME_HEIGHT*t2 && y <= FRAME_HEIGHT*t3 {
				return 2, 2
			} else if y >= FRAME_HEIGHT*t3 && y <= FRAME_HEIGHT*t4 {
				return 2, 3
			} else if y >= FRAME_HEIGHT*t4 && y <= FRAME_HEIGHT*t5 {
				return 2, 4
			} else if y >= FRAME_HEIGHT*t5 && y <= FRAME_HEIGHT*t6 {
				return 2, 5
			} else if y >= FRAME_HEIGHT*t6 && y <= FRAME_HEIGHT*t7 {
				return 2, 6
			} else if y >= FRAME_HEIGHT*t7 && y <= FRAME_HEIGHT*t8 {
				return 2, 7
			} else if y >= FRAME_HEIGHT*t8 && y <= FRAME_HEIGHT*t9 {
				return 2, 8
			}
		}
		if x >= FRAME_WIDTH*t3 && x <= FRAME_WIDTH*t4 {
			if y >= 0 && y <= FRAME_HEIGHT*t1 {
				return 3, 0
			} else if y >= FRAME_HEIGHT*t1 && y <= FRAME_HEIGHT*t2 {
				return 3, 1
			} else if y >= FRAME_HEIGHT*t2 && y <= FRAME_HEIGHT*t3 {
				return 3, 2
			} else if y >= FRAME_HEIGHT*t3 && y <= FRAME_HEIGHT*t4 {
				return 3, 3
			} else if y >= FRAME_HEIGHT*t4 && y <= FRAME_HEIGHT*t5 {
				return 3, 4
			} else if y >= FRAME_HEIGHT*t5 && y <= FRAME_HEIGHT*t6 {
				return 3, 5
			} else if y >= FRAME_HEIGHT*t6 && y <= FRAME_HEIGHT*t7 {
				return 3, 6
			} else if y >= FRAME_HEIGHT*t7 && y <= FRAME_HEIGHT*t8 {
				return 3, 7
			} else if y >= FRAME_HEIGHT*t8 && y <= FRAME_HEIGHT*t9 {
				return 3, 8
			}
		}
		if x >= FRAME_WIDTH*t4 && x <= FRAME_WIDTH*t5 {
			if y >= 0 && y <= FRAME_HEIGHT*t1 {
				return 4, 0
			} else if y >= FRAME_HEIGHT*t1 && y <= FRAME_HEIGHT*t2 {
				return 4, 1
			} else if y >= FRAME_HEIGHT*t2 && y <= FRAME_HEIGHT*t3 {
				return 4, 2
			} else if y >= FRAME_HEIGHT*t3 && y <= FRAME_HEIGHT*t4 {
				return 4, 3
			} else if y >= FRAME_HEIGHT*t4 && y <= FRAME_HEIGHT*t5 {
				return 4, 4
			} else if y >= FRAME_HEIGHT*t5 && y <= FRAME_HEIGHT*t6 {
				return 4, 5
			} else if y >= FRAME_HEIGHT*t6 && y <= FRAME_HEIGHT*t7 {
				return 4, 6
			} else if y >= FRAME_HEIGHT*t7 && y <= FRAME_HEIGHT*t8 {
				return 4, 7
			} else if y >= FRAME_HEIGHT*t8 && y <= FRAME_HEIGHT*t9 {
				return 4, 8
			}
		}
		if x >= FRAME_WIDTH*t5 && x <= FRAME_WIDTH*t6 {
			if y >= 0 && y <= FRAME_HEIGHT*t1 {
				return 5, 0
			} else if y >= FRAME_HEIGHT*t1 && y <= FRAME_HEIGHT*t2 {
				return 5, 1
			} else if y >= FRAME_HEIGHT*t2 && y <= FRAME_HEIGHT*t3 {
				return 5, 2
			} else if y >= FRAME_HEIGHT*t3 && y <= FRAME_HEIGHT*t4 {
				return 5, 3
			} else if y >= FRAME_HEIGHT*t4 && y <= FRAME_HEIGHT*t5 {
				return 5, 4
			} else if y >= FRAME_HEIGHT*t5 && y <= FRAME_HEIGHT*t6 {
				return 5, 5
			} else if y >= FRAME_HEIGHT*t6 && y <= FRAME_HEIGHT*t7 {
				return 5, 6
			} else if y >= FRAME_HEIGHT*t7 && y <= FRAME_HEIGHT*t8 {
				return 5, 7
			} else if y >= FRAME_HEIGHT*t8 && y <= FRAME_HEIGHT*t9 {
				return 5, 8
			}
		}
		if x >= FRAME_WIDTH*t6 && x <= FRAME_WIDTH*t7 {
			if y >= 0 && y <= FRAME_HEIGHT*t1 {
				return 6, 0
			} else if y >= FRAME_HEIGHT*t1 && y <= FRAME_HEIGHT*t2 {
				return 6, 1
			} else if y >= FRAME_HEIGHT*t2 && y <= FRAME_HEIGHT*t3 {
				return 6, 2
			} else if y >= FRAME_HEIGHT*t3 && y <= FRAME_HEIGHT*t4 {
				return 6, 3
			} else if y >= FRAME_HEIGHT*t4 && y <= FRAME_HEIGHT*t5 {
				return 6, 4
			} else if y >= FRAME_HEIGHT*t5 && y <= FRAME_HEIGHT*t6 {
				return 6, 5
			} else if y >= FRAME_HEIGHT*t6 && y <= FRAME_HEIGHT*t7 {
				return 6, 6
			} else if y >= FRAME_HEIGHT*t7 && y <= FRAME_HEIGHT*t8 {
				return 6, 7
			} else if y >= FRAME_HEIGHT*t8 && y <= FRAME_HEIGHT*t9 {
				return 6, 8
			}
		}
		if x >= FRAME_WIDTH*t7 && x <= FRAME_WIDTH*t8 {
			if y >= 0 && y <= FRAME_HEIGHT*t1 {
				return 7, 0
			} else if y >= FRAME_HEIGHT*t1 && y <= FRAME_HEIGHT*t2 {
				return 7, 1
			} else if y >= FRAME_HEIGHT*t2 && y <= FRAME_HEIGHT*t3 {
				return 7, 2
			} else if y >= FRAME_HEIGHT*t3 && y <= FRAME_HEIGHT*t4 {
				return 7, 3
			} else if y >= FRAME_HEIGHT*t4 && y <= FRAME_HEIGHT*t5 {
				return 7, 4
			} else if y >= FRAME_HEIGHT*t5 && y <= FRAME_HEIGHT*t6 {
				return 7, 5
			} else if y >= FRAME_HEIGHT*t6 && y <= FRAME_HEIGHT*t7 {
				return 7, 6
			} else if y >= FRAME_HEIGHT*t7 && y <= FRAME_HEIGHT*t8 {
				return 7, 7
			} else if y >= FRAME_HEIGHT*t8 && y <= FRAME_HEIGHT*t9 {
				return 7, 8
			}
		}
		if x >= FRAME_WIDTH*t8 && x <= FRAME_WIDTH*t9 {
			if y >= 0 && y <= FRAME_HEIGHT*t1 {
				return 8, 0
			} else if y >= FRAME_HEIGHT*t1 && y <= FRAME_HEIGHT*t2 {
				return 8, 1
			} else if y >= FRAME_HEIGHT*t2 && y <= FRAME_HEIGHT*t3 {
				return 8, 2
			} else if y >= FRAME_HEIGHT*t3 && y <= FRAME_HEIGHT*t4 {
				return 8, 3
			} else if y >= FRAME_HEIGHT*t4 && y <= FRAME_HEIGHT*t5 {
				return 8, 4
			} else if y >= FRAME_HEIGHT*t5 && y <= FRAME_HEIGHT*t6 {
				return 8, 5
			} else if y >= FRAME_HEIGHT*t6 && y <= FRAME_HEIGHT*t7 {
				return 8, 6
			} else if y >= FRAME_HEIGHT*t7 && y <= FRAME_HEIGHT*t8 {
				return 8, 7
			} else if y >= FRAME_HEIGHT*t8 && y <= FRAME_HEIGHT*t9 {
				return 8, 8
			}
		}
	}
	return -1, -1
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{180, 180, 180, 255})

	opts := ebiten.DrawImageOptions{}

	drawTiles(g, opts, screen)
	drawStats(g, screen)
}

func drawTiles(g *Game, opts ebiten.DrawImageOptions, screen *ebiten.Image) {
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLUMNS; c++ {
			tile := g.board.tiles[r][c]

			fox, foy, fw, fh := FRAME_OX, FRAME_OY, FRAME_WIDTH, FRAME_HEIGHT
			switch tile.Value {
			case 1:
				foy += 32
				fh *= 2
			case 2:
				foy += 32 * 2
				fh *= 3
			case 3:
				foy += 32 * 3
				fh *= 4
			case 4:
				foy += 32 * 4
				fh *= 5
			case 5:
				foy += 32 * 5
				fh *= 6
			case 6:
				foy += 32 * 6
				fh *= 7
			case 7:
				foy += 32 * 7
				fh *= 8
			case 8:
				foy += 32 * 8
				fh *= 9
			case 9:
				foy += 32 * 9
				fh *= 10
			default:
			}

			rspace := r/3 + 2
			cspace := c/3 + 2

			opts.GeoM.Translate(float64(c)*32+float64(cspace), float64(r)*32+float64(rspace))
			screen.DrawImage(
				tile.Img.SubImage(
					image.Rect(fox, foy, fw, fh),
				).(*ebiten.Image),
				&opts,
			)

			opts.GeoM.Reset()
		}
	}
}

func drawStats(g *Game, screen *ebiten.Image) {
	misses := fmt.Sprintf(MESSAGE_MISSES, g.misses)

	ebitenutil.DebugPrintAt(screen, g.message, 1, 291)
	ebitenutil.DebugPrintAt(screen, misses, 216, 291)

}

func main() {

	game := &Game{
		misses: 0,
		state:  1,
		board: &Board{
			tiles: [][]*Tile{
				{{}, {}, {}, {}, {}, {}, {}, {}, {}},
				{{}, {}, {}, {}, {}, {}, {}, {}, {}},
				{{}, {}, {}, {}, {}, {}, {}, {}, {}},
				{{}, {}, {}, {}, {}, {}, {}, {}, {}},
				{{}, {}, {}, {}, {}, {}, {}, {}, {}},
				{{}, {}, {}, {}, {}, {}, {}, {}, {}},
				{{}, {}, {}, {}, {}, {}, {}, {}, {}},
				{{}, {}, {}, {}, {}, {}, {}, {}, {}},
				{{}, {}, {}, {}, {}, {}, {}, {}, {}},
			},
		},
		aux: [][]*Tile{
			{{}, {}, {}, {}, {}, {}, {}, {}, {}},
			{{}, {}, {}, {}, {}, {}, {}, {}, {}},
			{{}, {}, {}, {}, {}, {}, {}, {}, {}},
			{{}, {}, {}, {}, {}, {}, {}, {}, {}},
			{{}, {}, {}, {}, {}, {}, {}, {}, {}},
			{{}, {}, {}, {}, {}, {}, {}, {}, {}},
			{{}, {}, {}, {}, {}, {}, {}, {}, {}},
			{{}, {}, {}, {}, {}, {}, {}, {}, {}},
			{{}, {}, {}, {}, {}, {}, {}, {}, {}},
		},
	}

	createAuxBoard(game)
	createBoard(game)

	fillAuxBoard(game)
	fillBoard(game)

	removeSome(game)
	addStats(game)

	ebiten.SetWindowSize(SCREEN_WIDTH*2, SCREEN_HEIGHT*2)
	ebiten.SetWindowTitle("SUDOKU by Rafael Goulart")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func removeSome(g *Game) {
	squares := SIDE * SIDE
	empties := squares * 3 / 4

	positions := make([]int, squares)
	for i := 0; i < squares; i++ {
		positions[i] = i
	}

	rand.Shuffle(len(positions), func(i, j int) {
		positions[i], positions[j] = positions[j], positions[i]
	})

	for _, p := range positions[:empties] {
		g.board.tiles[p/SIDE][p%SIDE].Value = 0
	}
}

func pattern(r, c int) int {
	return (BASE*(r%BASE) + r/BASE + c) % SIDE
}

func shuffle(s []int) []int {
	shuffled := make([]int, len(s))
	copy(shuffled, s)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled
}

func fillAuxBoard(g *Game) {
	rBASE := make([]int, BASE)
	for i := 0; i < BASE; i++ {
		rBASE[i] = i
	}

	rows := make([]int, 0, SIDE)
	cols := make([]int, 0, SIDE)
	for g := range shuffle(rBASE) {
		for r := range shuffle(rBASE) {
			rows = append(rows, g*BASE+r)
			cols = append(cols, g*BASE+r)
		}
	}

	nums := shuffle(rangeSlice(1, BASE*BASE+1))

	for r := 0; r < SIDE; r++ {
		for c := 0; c < SIDE; c++ {
			g.aux[r][c].Value = nums[pattern(rows[r], cols[c])]
		}
	}
}

func rangeSlice(start, end int) []int {
	slice := make([]int, end-start)
	for i := start; i < end; i++ {
		slice[i-start] = i
	}
	return slice
}

func createAuxBoard(g *Game) {
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLUMNS; c++ {
			g.aux[r][c].Value = 0
		}
	}
}

func createBoard(g *Game) {
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLUMNS; c++ {
			g.board.tiles[r][c].Value = 0
		}
	}
}

func fillBoard(g *Game) {
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLUMNS; j++ {
			g.board.tiles[i][j].Value = g.aux[i][j].Value
		}
	}
}

func addStats(g *Game) {
	tile, err := ebitenutil.NewImageFromURL("https://github.com/RafaelEtec/go_sudoku/blob/31af9f64952f8e622d8403e2febd2d6bb994a625/assets/images/numbers.png?raw=true")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLUMNS; j++ {
			if g.board.tiles[i][j].Value == 0 {
				g.board.tiles[i][j].isEditable = true
				g.board.tiles[i][j].isRight = false
			} else {
				g.board.tiles[i][j].isEditable = false
				g.board.tiles[i][j].isRight = true
			}
			g.board.tiles[i][j].Img = tile
		}
	}
}

func printAuxBoard(g *Game) {
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLUMNS; j++ {
			if j == ROWS-1 {
				fmt.Print(g.aux[i][j].Value)
			} else {
				fmt.Print(g.aux[i][j].Value, "-")
			}
			time.Sleep(time.Millisecond * 1)
		}
		fmt.Println("")
	}
	fmt.Println("")
}
