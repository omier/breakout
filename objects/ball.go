package objects

import (
	"image/color"
	"math"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

// Ball is a struct that holds its image and options to draw with.
type Ball struct {
	Image   *ebiten.Image
	Options *ebiten.DrawImageOptions
}

// Ball constants
const (
	BallRadius   = 10
	BallDiameter = BallRadius * 2
)

// BallColor stores ball's color
var BallColor = color.White

// NewBall creates a new ball in the center of the screen
func NewBall() *Ball {
	geom := ebiten.GeoM{}
	geom.Translate(getBallStart())

	image := ebiten.NewImage(BallDiameter, BallDiameter)

	// Draw center
	image.Set(BallRadius, BallRadius, BallColor)

	wg := sync.WaitGroup{}
	// Draw inner rings
	for radius := float64(BallRadius); radius > 0; radius-- {
		wg.Add(1)
		go func(r float64) {
			defer wg.Done()

			minAngle := math.Acos(1.0 - 1/float64(r))
			for ang := 0.0; ang < 180; ang += minAngle {

				x := r * math.Cos(ang)
				y := r * math.Sin(ang)

				image.Set(int(x)+BallRadius, int(y)+BallRadius, BallColor)
			}
		}(radius)
	}

	wg.Wait()

	return &Ball{
		Image: image,
		Options: &ebiten.DrawImageOptions{
			GeoM: geom,
		},
	}
}

// getBallStart returns the initial position of the ball
func getBallStart() (tx float64, ty float64) {
	x, y := ebiten.WindowSize()
	return float64(x)/2 - BallRadius/2, float64(y)/2 - BallRadius/2
}
