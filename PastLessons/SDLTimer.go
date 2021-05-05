package main

import (
	"strconv"

	"github.com/thigamore/SDL/general"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func SDLTimer() {
	sdl.Init(sdl.INIT_EVERYTHING)
	ttf.Init()

	const SCREEN_WIDTH = 500
	const SCREEN_HEIGHT = 500
	window, _ := general.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT)
	renderer := general.InitRenderer(window, true)
	resetText := general.InitTexture(renderer)
	fontColor := sdl.Color{R: 0, G: 0, B: 0, A: 255}
	font, err := ttf.OpenFont("arial.ttf", 28)

	if err != nil {
		panic(sdl.GetError())

	}
	resetText.SetFont(font)
	timeText := resetText.Copy()

	resetText.LoadText("Press Enter to Reset Start Time.", fontColor)

	running := true
	var e sdl.Event
	startTime := uint32(0)

	for running {
		e = sdl.PollEvent()
		for e != nil {
			switch t := e.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				if t.Keysym.Sym == sdl.K_RETURN {
					startTime = sdl.GetTicks()
				}
			}
			e = sdl.PollEvent()
		}
		renderer.Clear()
		renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)

		resetText.Render(0, 0, nil)

		timeText.LoadText(strconv.FormatUint(uint64(sdl.GetTicks()-startTime), 10), fontColor)
		timeText.Render(0, SCREEN_HEIGHT/2, nil)

		renderer.Present()
	}
}
