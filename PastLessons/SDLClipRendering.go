package main

import (
	"github.com/thigamore/SDL/general"
	"github.com/veandco/go-sdl2/sdl"
)

func SDLClipRendering() {
	window, _ := general.InitWindow(500, 500)
	screenWidth := int32(500)
	screenHeight := int32(500)
	renderer := general.InitRenderer(window, false)
	SpriteClips := make([]sdl.Rect, 4)
	spriteSheetTexture := general.InitTexture(renderer)

	spriteSheetTexture.LoadImage("dots.png", &sdl.Color{R: 0x00, G: 0xFF, B: 0xFF})

	//Set top left sprite
	SpriteClips[0].X = 0
	SpriteClips[0].Y = 0
	SpriteClips[0].W = 100
	SpriteClips[0].H = 100

	//Set top right sprite
	SpriteClips[1].X = 100
	SpriteClips[1].Y = 0
	SpriteClips[1].W = 100
	SpriteClips[1].H = 100

	//Set bottom left sprite
	SpriteClips[2].X = 0
	SpriteClips[2].Y = 100
	SpriteClips[2].W = 100
	SpriteClips[2].H = 100

	//Set bottom right sprite
	SpriteClips[3].X = 100
	SpriteClips[3].Y = 100
	SpriteClips[3].W = 100
	SpriteClips[3].H = 100

	notQuit := true
	var e sdl.Event
	for notQuit {
		e = sdl.PollEvent()
		for e != nil {
			if e.GetType() == sdl.QUIT {
				notQuit = false
				break
			}
			e = sdl.PollEvent()
		}
		renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)
		renderer.Clear()

		spriteSheetTexture.Render(0, 0, &SpriteClips[0])

		spriteSheetTexture.Render(screenWidth-SpriteClips[1].W, 0, &SpriteClips[1])

		spriteSheetTexture.Render(0, screenHeight-SpriteClips[2].H, &SpriteClips[2])

		spriteSheetTexture.Render(screenWidth-SpriteClips[3].W, screenHeight-SpriteClips[3].H, &SpriteClips[3])

		renderer.Present()

	}
}
