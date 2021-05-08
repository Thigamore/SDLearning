package main

import (
	"github.com/thigamore/SDL/general"
	"github.com/veandco/go-sdl2/sdl"
)

func SDLAlphaBlending() {
	window, _ := general.InitWindow(500, 500)
	renderer := general.InitRenderer(window, false)
	foreground := general.InitTexture(renderer)
	background := general.InitTexture(renderer)
	foreground.LoadImage("fadeout.png", nil)
	background.LoadImage("fadein.png", nil)
	alpha := uint8(255)

	notQuit := true
	var e sdl.Event
	for notQuit {
		e = sdl.PollEvent()
		for e != nil {
			switch t := e.(type) {
			case *sdl.QuitEvent:
				notQuit = false
			case *sdl.KeyboardEvent:
				switch t.Keysym.Sym {
				case sdl.K_w:
					if (alpha + 32) >= 255 {
						alpha = 255
					} else {
						alpha += 32
					}
				case sdl.K_s:
					if (alpha - 32) <= 0 {
						alpha = 0
					} else {
						alpha -= 32
					}

				}

			}
			e = sdl.PollEvent()
		}
		renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)
		renderer.Clear()

		background.Render(0, 0, nil)

		foreground.SetAlpha(alpha)
		foreground.Render(0, 0, nil)

		renderer.Present()

	}
}
