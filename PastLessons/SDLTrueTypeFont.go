package main

import (
	"github.com/thigamore/SDL/general"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func SDLTrueTypeFont() {
	window, _ := general.InitWindow(700, 300)
	renderer := general.InitRenderer(window, true)
	texture := general.InitTexture(renderer)
	const SCREEN_WIDTH = 700
	const SCREEN_HEIGHT = 300
	font, err := ttf.OpenFont("arial.ttf", 28)
	if err != nil {
		panic(sdl.GetError())
	}

	color := sdl.Color{R: 0, G: 0, B: 0, A: 255}

	texture.SetFont(font)
	texture.LoadText("The quick brown fox jumps over the lazy dog.", color)

	running := true
	var e sdl.Event
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

		texture.Render((SCREEN_WIDTH-texture.Width)/2, (SCREEN_HEIGHT-texture.Height)/2, nil)

		renderer.Present()
	}
	toClose := []general.Freeable{texture}
	general.CloseAll(toClose, renderer, window)

}
