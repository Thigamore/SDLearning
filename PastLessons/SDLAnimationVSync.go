package main

import (
	"github.com/thigamore/SDL/general"
	"github.com/veandco/go-sdl2/sdl"
)

func SDLAnimationVSync() {
	window, _ := general.InitWindow(500, 500)
	renderer := general.InitRenderer(window, true)
	stickManSheet := general.InitTexture(renderer)
	const WALKING_ANIMATION_FRAMES = 4
	stickManClips := make([]sdl.Rect, WALKING_ANIMATION_FRAMES)
	const SCREEN_WIDTH = 500
	const SCREEN_HEIGHT = 500

	stickManSheet.LoadImage("foo.png", true)
	//Set sprite clips
	stickManClips[0].X = 0
	stickManClips[0].Y = 0
	stickManClips[0].W = 64
	stickManClips[0].H = 205

	stickManClips[1].X = 64
	stickManClips[1].Y = 0
	stickManClips[1].W = 64
	stickManClips[1].H = 205

	stickManClips[2].X = 128
	stickManClips[2].Y = 0
	stickManClips[2].W = 64
	stickManClips[2].H = 205

	stickManClips[3].X = 196
	stickManClips[3].Y = 0
	stickManClips[3].W = 64
	stickManClips[3].H = 205

	running := true
	var e sdl.Event
	frame := 0
	var currentClip *sdl.Rect
	for running {
		e = sdl.PollEvent()
		for e != nil {
			if e.GetType() == sdl.QUIT {
				running = false
			}
			e = sdl.PollEvent()
		}

		renderer.Clear()
		renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)

		currentClip = &stickManClips[frame%4]
		stickManSheet.Render((SCREEN_WIDTH-currentClip.W)/2, (SCREEN_HEIGHT-currentClip.H)/2, currentClip)

		renderer.Present()
		frame += 1
	}
}
