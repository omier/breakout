package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/omier/breakout/objects"
)

// Game implements ebiten.Game interface.
type Game struct {
	Paddle *objects.Paddle
	Ball   *objects.Ball
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	if err := g.Paddle.Update(); err != nil {
		return err
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	screen.Fill(color.Black)
	screen.DrawImage(g.Paddle.Image, g.Paddle.Options)
	screen.DrawImage(g.Ball.Image, g.Ball.Options)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(1600, 900)
	ebiten.SetWindowTitle("Breakout")

	game := &Game{
		Paddle: objects.NewPaddle(),
		Ball:   objects.NewBall(),
	}

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
