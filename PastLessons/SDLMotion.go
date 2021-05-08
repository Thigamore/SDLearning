package main

import (
	"github.com/thigamore/SDL/general"
	"github.com/veandco/go-sdl2/sdl"
)

func SDLMotion() {
	SCREENDIM := general.Dimension{W: 500, H: 500}
	window, _ := general.InitWindow(SCREENDIM.W, SCREENDIM.H)
	renderer := general.InitRenderer(window, false)
	frameMan := general.InitFrameRateManager(60)
	dotTexture := general.InitTexture(renderer)
	dot := general.InitPEntity(
		dotTexture,
		sdl.Point{X: SCREENDIM.W / 2, Y: SCREENDIM.H / 2},
		general.Velocity{X: 0, Y: 0},
		general.Dimension{W: 20, H: 20},
	)

	dotTexture.LoadImage("dot.bmp", nil)

	var e sdl.Event
	running := true

	frameMan.Start()
	for running {
		e = sdl.PollEvent()
		for e != nil {
			switch t := e.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				if e.GetType() == sdl.KEYDOWN {
					dot.HandleEvent(t, 10)
				}
			}
			e = sdl.PollEvent()
		}
		renderer.Clear()
		renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)

		dot.Move(SCREENDIM)
		dot.Render()

		renderer.Present()
		frameMan.Run(true)
	}
}
