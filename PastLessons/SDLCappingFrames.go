package main

import (
	"strconv"

	"github.com/thigamore/SDL/general"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func SDLCappingFrames() {
	const SCREEN_WIDTH = 500
	const SCREEN_HEIGHT = 500
	const WANTED_FPS = 60
	window, _ := general.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT)
	renderer := general.InitRenderer(window, false)
	timer := general.InitTimer()
	text := general.InitTexture(renderer)
	lastTime := uint32(0)
	color := sdl.Color{R: 0, G: 0, B: 0, A: 255}
	font, err := ttf.OpenFont("arial.ttf", 25)
	if err != nil {
		panic(sdl.GetError())
	}

	text.SetFont(font)

	running := true
	var e sdl.Event

	timer.Start()
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

		if (timer.Run() - lastTime) < (1000 / WANTED_FPS) {
			sdl.Delay((1000 / WANTED_FPS) - (timer.Run() - lastTime))
		}

		text.LoadText("Average FPS (With Cap)"+strconv.FormatFloat((1000/float64((timer.Run()-lastTime))), 'f', 5, 64), color)
		lastTime = timer.Run()

		text.Render(0, SCREEN_HEIGHT/2, nil)

		renderer.Present()

	}
}
