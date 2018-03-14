package main

import (
	"fmt"
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	// Settings
	screenWidth  = 800
	screenHeight = 600
)

var (
	loadedSprite    *ebiten.Image
	leftSprite      *ebiten.Image
	rightSprite     *ebiten.Image
	idleSprite      *ebiten.Image
	backgroundImage *ebiten.Image
)

func init() {
	// Preload images
	var err error
	rightSprite, _, err = ebitenutil.NewImageFromFile("res/sprite.png", ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	leftSprite, _, err = ebitenutil.NewImageFromFile("res/sprite.png", ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	idleSprite, _, err = ebitenutil.NewImageFromFile("res/sprite.png", ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	backgroundImage, _, err = ebitenutil.NewImageFromFile("res/background.png", ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
}

var (
	charX = 50
	charY = 380
)

func update(screen *ebiten.Image) error {
	// Controls
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		// Selects preloaded sprite
		loadedSprite = leftSprite
		// Moves character 3px right
		charX -= 3
	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		// Selects preloaded sprite
		loadedSprite = rightSprite
		// Moves character 3px left
		charX += 3
	} else {
		loadedSprite = idleSprite
	}

	if ebiten.IsRunningSlowly() {
		return nil
	}

	// Draws Background Image
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	screen.DrawImage(backgroundImage, op)

	// Draws selected sprite image
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(float64(charX), float64(charY))
	screen.DrawImage(loadedSprite, op)

	// FPS counter
	fps := fmt.Sprintf("FPS: %f", ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, fps)

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Platformer (Ebiten Demo)"); err != nil {
		panic(err)
	}
}
