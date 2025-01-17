package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// game struct
type Game struct {
}

// update loop
func (g *Game) Update() error {

	return nil

}

// draw method

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 0, G: 255, B: 0, A: 255})
}

// layout method
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}
func main() {
	game := &Game{}

	err := ebiten.RunGame(game)
	if err != nil {
		fmt.Println("Yo couldnt run game foo")
		panic(err) // game aint runnin shut it down!

	}
}
