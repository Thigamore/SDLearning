package main

import (
	"fmt"

	sdlImg "github.com/veandco/go-sdl2/img"
	sdl "github.com/veandco/go-sdl2/sdl"
)

type freeable interface {
	Free()
}

//Prepares Window, Window Surface
func initWindow() (*sdl.Window, *sdl.Surface) {
	window, err := sdl.CreateWindow("SDL", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 500, 500, 0)
	if err != nil {
		panic("There was an error creating the window")
	}
	surface, err := window.GetSurface()
	if err != nil {
		panic("Error getting the surface from the window.")
	}

	return window, surface
}

//Prepares Renderer and Screen
func initRenderer(window *sdl.Window) *sdl.Renderer {
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic("Error loading the renderer")
	}
	return renderer
}

//Loads an image
func loadMedia(path string) *sdl.Surface {
	loadSurface, err := sdlImg.Load(path)
	if err != nil {
		panic("Error loading image")
	}
	return loadSurface
}

//Frees space allocated
func close(toFree ...freeable) {
	for _, i := range toFree {
		i.Free()
	}
}

func main() {
	//-------------------------Initializing Variables-----------------------
	window, _ := initWindow()
	renderer := initRenderer(window)
	screenWidth := int32(500)
	screenHeight := int32(500)

	//-------------------------Event Loop ----------------------------------
	notQuit := true
	var e sdl.Event
	for notQuit {
		e = sdl.PollEvent()
		for e != nil {
			if e.GetType() == sdl.QUIT {
				notQuit = false
				break
			}
			e = sdl.PollEvent()
		}

		//Clearing Screen
		renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)
		renderer.Clear()

		//Render red filled quad
		fillRect := sdl.Rect{X: screenWidth / 4, Y: screenHeight / 4, W: screenWidth / 2, H: screenHeight / 2}
		err := renderer.SetDrawColor(0xFF, 0x00, 0x00, 0xFF)
		if err != nil {
			fmt.Println(sdl.GetError())
		}
		renderer.FillRect(&fillRect)
		//renderer.DrawRect for outlined rectangle

		//Updating Window (Frame Update)
		renderer.Present()
	}
	fmt.Println(window)
}
