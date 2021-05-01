package main

import (
	"github.com/thigamore/SDL/general"
	"github.com/veandco/go-sdl2/sdl"
)

func SDLColorMod() {
	window, _ := general.InitWindow(500, 500)
	renderer := general.InitRenderer(window, false)
	texture := general.InitTexture(renderer)
	texture.LoadImage("colors.png", false)

	r := uint8(255)
	g := uint8(255)
	b := uint8(255)

	var e sdl.Event
	notQuit := true
	for notQuit {
		e = sdl.PollEvent()
		for e != nil && notQuit {
			switch t := e.(type) {
			case *sdl.QuitEvent:
				notQuit = false
			case *sdl.KeyboardEvent:
				switch t.Keysym.Sym {
				case sdl.K_q:
					r += 32
				case sdl.K_w:
					g += 32
				case sdl.K_e:
					b += 32
				case sdl.K_s:
					r -= 32
				case sdl.K_d:
					g -= 32
				case sdl.K_a:
					b -= 32
				}

			}
			e = sdl.PollEvent()
			renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)
			renderer.Clear()

			texture.SetColor(r, g, b)
			texture.Render(0, 0, &sdl.Rect{X: 0, Y: 0, W: texture.Height, H: texture.Width})

			renderer.Present()
		}
	}
}
