package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 0xff, A: 0xff})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Fill")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
