package main

import (
	"strconv"

	"github.com/thigamore/SDL/general"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func SDLAdvancedTimer() {
	const SCREEN_WIDTH = 500
	const SCREEN_HEIGHT = 500
	window, _ := general.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT)
	renderer := general.InitRenderer(window, false)
	instructionS := general.InitTexture(renderer)
	fontColor := sdl.Color{R: 0, G: 0, B: 0, A: 255}
	font, err := ttf.OpenFont("arial.ttf", 28)
	if err != nil {
		panic(sdl.GetError())
	}
	instructionS.SetFont(font)
	instructionP := instructionS.Copy()
	timeText := instructionS.Copy()

	instructionS.LoadText("Press S to Start or Stop the Timer", fontColor)
	instructionP.LoadText("Press P to pause or Unpause the Timer", fontColor)

	running := true
	var e sdl.Event
	timer := general.InitTimer()

	for running {
		e = sdl.PollEvent()
		for e != nil {
			switch t := e.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				if e.GetType() == sdl.KEYDOWN {
					switch t.Keysym.Sym {
					case sdl.K_s:
						timer.Start()
					case sdl.K_p:
						timer.Pause()
					}
				}
			}
			e = sdl.PollEvent()
		}
		renderer.Clear()
		renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)

		timeText.LoadText("The timer has been running for "+strconv.FormatUint(uint64(timer.Run()), 10)+"seconds", fontColor)

		instructionS.Render(0, 0, nil)
		instructionP.Render(0, 20, nil)
		timeText.Render(0, SCREEN_HEIGHT/2, nil)

		renderer.Present()
		sdl.Delay(10)
	}

}
