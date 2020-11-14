package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Paddle is a struct that holds its image and options to draw with.
type Paddle struct {
	Image   *ebiten.Image
	Options *ebiten.DrawImageOptions
}

const (
	// PaddleWidth is the default width of the paddle.
	PaddleWidth = 100

	// PaddleHeight is the default height of the paddle.
	PaddleHeight = 20

	// PaddleBottomMargin is the default margin of the paddle from the bottom.
	PaddleBottomMargin = 50

	// PaddleSpeed is the default speed of the paddle on the x axis.
	PaddleSpeed = 10
)

// PaddleColor is the default color of the paddle.
var PaddleColor = color.White

// Update updates the paddle position according to the input.
func (p *Paddle) Update() error {
	paddleX := p.GetX()
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && paddleX-PaddleSpeed > 0 {
		p.Options.GeoM.Translate(-PaddleSpeed, 0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		x, _ := ebiten.WindowSize()
		if paddleX+PaddleWidth+PaddleSpeed < float64(x) {
			p.Options.GeoM.Translate(PaddleSpeed, 0)
		}
	}

	return nil
}

// GetX is the current x coordinate of the paddle.
func (p *Paddle) GetX() float64 {
	return p.Options.GeoM.Element(0, 2)
}

// GetY is the current y coordinate of the paddle.
func (p *Paddle) GetY() float64 {
	return p.Options.GeoM.Element(1, 2)
}

// NewPaddle returns a new paddle with the default values based on the window size.
func NewPaddle() *Paddle {
	image := ebiten.NewImage(PaddleWidth, PaddleHeight)
	image.Fill(PaddleColor)

	geom := ebiten.GeoM{}
	geom.Translate(getPaddleStart())

	return &Paddle{
		Image: image,
		Options: &ebiten.DrawImageOptions{
			GeoM: geom,
		},
	}
}

// getPaddleStart is the starting point of the paddle based on the window size.
func getPaddleStart() (tx float64, ty float64) {
	x, y := ebiten.WindowSize()
	return float64(x)/2 - PaddleWidth/2, float64(y) - PaddleBottomMargin
}
