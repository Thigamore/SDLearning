package main

import (
	"github.com/thigamore/SDL/general"
)

func main() {
	window, _ := general.InitWindow(500, 500)
	renderer := general.InitRenderer(window, true)

}
