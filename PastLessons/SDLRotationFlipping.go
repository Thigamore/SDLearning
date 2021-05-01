package main

import (
	"github.com/thigamore/SDL/general"
	"github.com/veandco/go-sdl2/sdl"
)

func SDLRotationFlipping() {
	window, _ := general.InitWindow(500, 500)
	renderer := general.InitRenderer(window, false)
	texture := general.InitTexture(renderer)
	const SCREEN_WIDTH = 500
	const SCREEN_HEIGHT = 500
	texture.LoadImage("arrow.png", false)

	running := true
	var e sdl.Event
	degrees := float64(0)
	flipType := sdl.FLIP_NONE

	for running {
		e = sdl.PollEvent()
		for e != nil {
			switch t := e.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				switch t.Keysym.Sym {
				case sdl.K_a:
					degrees -= 60
				case sdl.K_d:
					degrees += 60
				case sdl.K_q:
					flipType = sdl.FLIP_HORIZONTAL
				case sdl.K_w:
					flipType = sdl.FLIP_NONE
				case sdl.K_e:
					flipType = sdl.FLIP_VERTICAL
				}
			}
			e = sdl.PollEvent()
		}
		renderer.Clear()
		renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)

		texture.RenderFlip((SCREEN_WIDTH-texture.Width)/2, (SCREEN_HEIGHT-texture.Height)/2, nil, degrees, nil, flipType)

		renderer.Present()
	}
}
