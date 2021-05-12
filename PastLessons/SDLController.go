package main

import (
	"fmt"
	"math"

	"github.com/thigamore/SDL/general"
	"github.com/veandco/go-sdl2/sdl"
)

func SDLController() {
	const SCREEN_WIDTH = 500
	const SCREEN_HEIGHT = 500
	window, _ := general.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT)
	defer window.Destroy()
	renderer := general.InitRenderer(window, true)
	defer renderer.Destroy()
	var gameController *sdl.Joystick = nil
	const JOYSTICK_DEAD_ZONE = 8000
	arrow := general.InitTexture(renderer)
	arrow.LoadImage("arrow.png", nil, nil)

	if !(sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")) {
		fmt.Println("Warning: Linear texture filtering not enabled!")
	}
	if sdl.NumJoysticks() < 1 {
		fmt.Println(sdl.NumJoysticks())
		panic("No controller connected")
	} else {
		gameController = sdl.JoystickOpen(0)
		defer gameController.Close()
		if gameController == nil {
			panic("Error connecting to controller")
		}
	}

	running := true
	var e sdl.Event
	xDir := 0
	yDir := 0

	for running {
		e = sdl.PollEvent()
		for e != nil {
			switch t := e.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.JoyAxisEvent:
				if t.Which == gameController.InstanceID() {
					if t.Axis == 0 {
						if t.Value < (-1 * JOYSTICK_DEAD_ZONE) {
							xDir = -1
						} else if t.Value > JOYSTICK_DEAD_ZONE {
							xDir = 1
						} else {
							xDir = 0
						}
					} else if t.Axis == 1 {
						if t.Value < (-1 * JOYSTICK_DEAD_ZONE) {
							yDir = -1
						} else if t.Value > JOYSTICK_DEAD_ZONE {
							yDir = 1
						} else {
							xDir = 0
						}
					}
				}
			}
			e = sdl.PollEvent()
		}
		renderer.Clear()
		renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)

		joystickAngle := math.Atan2(float64(yDir), float64(xDir)) * (180.0 / math.Pi)

		if xDir == 0 && yDir == 0 {
			joystickAngle = 0
		}

		arrow.RenderFlip((SCREEN_WIDTH-arrow.Width)/2, (SCREEN_HEIGHT-arrow.Height)/2, nil, joystickAngle, nil, 0)

		renderer.Present()

	}

}
