package main

import (
	"fmt"

	"github.com/thigamore/SDL/general"
	"github.com/veandco/go-sdl2/sdl"
)

func SDLTextureManipulation() {
	sdl.Init(sdl.INIT_EVENTS | sdl.INIT_VIDEO | sdl.INIT_TIMER)
	SCREENDIM := general.Dimension{W: 500, H: 500}
	window, _ := general.InitWindow(SCREENDIM.W, SCREENDIM.H)
	renderer := general.InitRenderer(window, false)
	foo := general.InitTexture(renderer)

	foo.LoadImage("foo.png", nil, window)

	foo.Lock()
	pixelFormat, err := window.GetPixelFormat()
	if err != nil {
		panic(sdl.GetError())
	}

	mappingFormat, err := sdl.AllocFormat(uint(pixelFormat))
	if err != nil {
		panic(sdl.GetError())
	}

	pixels := *foo.GetPixel()
	pixelCount := (foo.GetPitch()) * int(foo.Height)
	//colorKey := sdl.MapRGB(mappingFormat, 0, 0xFF, 0xFF)
	//transparent := sdl.MapRGBA(mappingFormat, 0xFF, 0xFF, 0xFF, 0)

	fmt.Println(pixels)
	for i := 0; i < pixelCount; i += 4 {
		if (pixels[i] | pixels[i+1] | pixels[i+2]) == byte(0xFF) {
			pixels[i] = 0xFF
			pixels[i+1] = 0xFF
			pixels[i+2] = 0xFF
			pixels[i+3] = 0xFF
		}
	}
	foo.Unlock()
	mappingFormat.Free()
	renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)
	renderer.Copy(foo.Texture, nil, &sdl.Rect{W: 500, H: 500})
	renderer.Present()
	sdl.Delay(2000)
}
