package main

import (
	"github.com/thigamore/SDL/general"
	"github.com/veandco/go-sdl2/sdl"
)

func SDLButtonsPain() {
	const SCREEN_WIDTH = 600
	const SCREEN_HEIGHT = 400
	window, _ := general.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT)
	renderer := general.InitRenderer(window, true)
	spriteSheet := general.InitTexture(renderer)

	spriteSheet.LoadImage("button.png", false)

	const BUTTON_WIDTH = 300
	const BUTTON_HEIGHT = 200
	const TOTAL_BUTTONS = 4

	spriteRects := map[string]sdl.Rect{
		"out":  {X: 0, Y: 0, W: BUTTON_WIDTH, H: BUTTON_HEIGHT},
		"over": {X: 0, Y: 200, W: BUTTON_WIDTH, H: BUTTON_HEIGHT},
		"down": {X: 0, Y: 400, W: BUTTON_WIDTH, H: BUTTON_HEIGHT},
		"up":   {X: 0, Y: 600, W: BUTTON_WIDTH, H: BUTTON_HEIGHT},
	}

	buttonPositions := []sdl.Point{
		{X: 0, Y: 0},
		{X: 300, Y: 0},
		{X: 0, Y: 200},
		{X: 300, Y: 200},
	}

	buttons := make([]general.Button, TOTAL_BUTTONS)

	for i := 0; i < TOTAL_BUTTONS; i++ {
		buttons[i] = general.InitButton(&buttonPositions[i], spriteSheet, spriteRects["out"], 300, 200)
	}

	running := true
	var e sdl.Event
	for running {
		e = sdl.PollEvent()
		for e != nil {
			switch e.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent, *sdl.MouseButtonEvent:
				for i, _ := range buttons {
					buttons[i].SetSpritePos(spriteRects[buttons[i].HandleEvent(e)])
				}

			}
			e = sdl.PollEvent()
		}
		renderer.Clear()
		renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)
		for _, i := range buttons {
			i.Render()
		}

		renderer.Present()
		sdl.Delay(10)
	}

}
