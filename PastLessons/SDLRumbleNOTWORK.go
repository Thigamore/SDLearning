package main

import (
	"fmt"

	"github.com/thigamore/SDL/general"
	"github.com/veandco/go-sdl2/sdl"
)

func SDLRumbleNOTWORK() {
	sdl.Init(sdl.INIT_HAPTIC)
	const SCREEN_HEIGHT = 500
	const SCREEN_WIDTH = 500
	window, _ := general.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT)
	renderer := general.InitRenderer(window, true)
	splash := general.InitTexture(renderer)
	splash.LoadImage("splash.png", nil)
	var gameController *sdl.Joystick
	var controllerHaptic *sdl.Haptic

	if sdl.NumJoysticks() <= 0 {
		panic("No controller connected")
	}
	gameController = sdl.JoystickOpen(0)
	if gameController == nil {
		panic("Error connecting to controller")
	}
	controllerHaptic, err := sdl.HapticOpenFromJoystick(gameController)
	if err != nil {
		panic(sdl.GetError())
	}
	if controllerHaptic.RumbleInit() != nil {
		fmt.Println("Warning: Unable to initialize rumble!")
	}

	running := true
	var e sdl.Event

	for running {
		e = sdl.PollEvent()
		for e != nil {
			switch e.GetType() {
			case sdl.QUIT:
				running = false
			case sdl.JOYBUTTONDOWN:
				err := controllerHaptic.RumblePlay(0.75, 500)
				if err != nil {
					fmt.Println("Warning: Unable to play rumble!")
				}
			}
		}
		renderer.Clear()
		renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)

		splash.Render(0, 0, nil)

		renderer.Present()
	}

	controllerHaptic.Close()
	gameController.Close()

}
