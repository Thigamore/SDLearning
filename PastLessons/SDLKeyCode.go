package main

import (
	"github.com/thigamore/SDL/general"
	"github.com/veandco/go-sdl2/sdl"
)

func SDLKeyCode() {
	window, _ := general.InitWindow(500, 500)
	renderer := general.InitRenderer(window, true)

	up := general.InitTexture(renderer)
	up.LoadImage("up.png", nil)
	down := general.InitTexture(renderer)
	down.LoadImage("down.png", nil)
	left := general.InitTexture(renderer)
	left.LoadImage("left.png", nil)
	right := general.InitTexture(renderer)
	right.LoadImage("right.png", nil)
	press := general.InitTexture(renderer)
	press.LoadImage("press.png", nil)

	running := true
	var e sdl.Event
	var currentTexture *general.ATexture = nil
	for running {
		e = sdl.PollEvent()
		for e != nil {
			if e.GetType() == sdl.QUIT {
				running = false
			} else {
				keyStates := sdl.GetKeyboardState()
				if keyStates[sdl.SCANCODE_UP] == 1 {
					currentTexture = up
				} else if keyStates[sdl.SCANCODE_DOWN] == 1 {
					currentTexture = down
				} else if keyStates[sdl.SCANCODE_LEFT] == 1 {
					currentTexture = left
				} else if keyStates[sdl.SCANCODE_RIGHT] == 1 {
					currentTexture = right
				} else {
					currentTexture = press
				}
			}
			e = sdl.PollEvent()
		}
		renderer.Clear()
		renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)

		currentTexture.Render(0, 0, nil)

		renderer.Present()
	}
}
