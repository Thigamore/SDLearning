package main

import (
	"github.com/thigamore/SDL/general"
	"github.com/veandco/go-sdl2/sdl"
)

func SDLWindow() {
	window, _ := general.InitWindow()
	renderer := general.InitRenderer(window)
	screenWidth := int32(500)
	screenHeight := int32(500)
	imgSurface := general.LoadMedia("viewport.png")
	texture, _ := renderer.CreateTextureFromSurface(imgSurface)
	topLeftViewport := sdl.Rect{X: 0, Y: 0, W: screenWidth / 2, H: screenHeight / 2}

	renderer.SetViewport(&topLeftViewport)
	renderer.Copy(texture, nil, nil)
	renderer.Present()

	sdl.Delay(2000)

}
