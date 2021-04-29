package main

import (
	"github.com/thigamore/SDL/general"
)

func main() {
	window, screenSurface := general.InitWindow()
	renderer := general.InitRenderer(window)
	texture := general.InitTexture(renderer)
}
