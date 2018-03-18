package main

import (
	"fmt"
	"image"
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
	spriteFps       int
	m               map[string]*ebitenSprite
)

type ebitenSprite struct {
	listSprites    []ebitenSpriteDetails
	isCurrent      bool
	spritePosition int
}

type ebitenSpriteDetails struct {
	rec   image.Rectangle
	frame int
}

type ebitenSpriteMove struct {
	jump  bool
	right bool
	left  bool
}

func init() {
	spriteFps = 0
	//
	//s := sprite{spriteRect: image.Rect(0, 0, 100, 100), second: 50}
	makeSprite()

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

func makeSprite() {

	m = make(map[string]*ebitenSprite)

	var arrEbitenSpriteDetails []ebitenSpriteDetails

	idleSpriteValue := [][]int{{4, 15, 45, 48, 60}, {52, 15, 93, 48, 5}, {104, 15, 145, 48, 5}, {152, 15, 193, 48, 120}, {199, 15, 240, 48, 8}, {250, 15, 291, 48, 8}, {299, 15, 340, 48, 8}, {349, 15, 390, 48, 8}}

	for _, v := range idleSpriteValue {
		r := image.Rect(v[0], v[1], v[2], v[3])

		s := ebitenSpriteDetails{rec: r, frame: v[4]}
		arrEbitenSpriteDetails = append(arrEbitenSpriteDetails, s)
	}

	m["idle"] = &ebitenSprite{listSprites: arrEbitenSpriteDetails, isCurrent: false, spritePosition: 0}

}

func chooseSprite(mapKey string) image.Rectangle {

	spriteFps++

	p := &m[mapKey].spritePosition
	frameDetails := m[mapKey].listSprites[*p]

	/*fmt.Println("********RESUME*************")
	fmt.Println("spriteFPS ", spriteFps)
	fmt.Println("frameDetails.frame", frameDetails.frame)
	fmt.Println("*p+1 ", *p+1)
	fmt.Println("len(m[mapKey].listSprites)", len(m[mapKey].listSprites))
	fmt.Println("***************************")*/

	//if end of frame number then return the next rectangle
	if spriteFps > frameDetails.frame {

		//if end of array sprites
		if *p+1 >= len(m[mapKey].listSprites) {
			fmt.Println("end")
			*p = 0
			frameDetails = m[mapKey].listSprites[*p]
			return frameDetails.rec
		}
		spriteFps = 0
		*p++
		frameDetails = m[mapKey].listSprites[*p]
		return frameDetails.rec
	}

	return frameDetails.rec

}

var (
	charX = 80
	charY = 100
)

func getMoveFromKeyboard() {

	eMove := ebitenSpriteMove{jump: false, right: false, left: false}
	// Controls

	loadedSprite = idleSprite

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		eMove.left = true
		charX -= 3
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		eMove.right = true
		charX += 3
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		eMove.jump = true
		charY += 3
	}

}

func update(screen *ebiten.Image) error {

	if ebiten.IsRunningSlowly() {
		return nil
	}

	getMoveFromKeyboard()

	// Draws Background Image
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	screen.DrawImage(backgroundImage, op)

	// Draws selected sprite image
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(8, 8)
	op.GeoM.Translate(float64(charX), float64(charY))

	r := chooseSprite("idle")
	//r := m["idle"].listSprites[2]
	op.SourceRect = &r
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
