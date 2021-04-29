package main

import (
	"github.com/thigamore/SDL/general"
	"github.com/veandco/go-sdl2/sdl"
)

func SDLColorKey() {
	window, _ := general.InitWindow()
	renderer := general.InitRenderer(window)

	fooSurface := general.LoadMedia("foo.png")
	backgroundSurface := general.LoadMedia("background.png")

	fooSurface.SetColorKey(true, sdl.MapRGB(fooSurface.Format, 0, 0xFF, 0xFF))
	backgroundSurface.SetColorKey(true, sdl.MapRGB(backgroundSurface.Format, 0, 0xFF, 0xFF))

	fooTexture, _ := renderer.CreateTextureFromSurface(fooSurface)
	backgroundTexture, _ := renderer.CreateTextureFromSurface(backgroundSurface)

	renderQuad := sdl.Rect{X: 0, Y: 0, W: backgroundSurface.W, H: backgroundSurface.H}
	renderer.Copy(backgroundTexture, nil, &renderQuad)
	renderQuad = sdl.Rect{X: 100, Y: 100, W: fooSurface.W, H: fooSurface.H}
	renderer.Copy(fooTexture, nil, &renderQuad)

	general.Close(fooSurface, backgroundSurface)

	renderer.Present()
	sdl.Delay(2000)
}
