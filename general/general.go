package general

import (
	sdlImg "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type freeable interface {
	Free()
}

//Prepares Window, Window Surface
func InitWindow() (*sdl.Window, *sdl.Surface) {
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
func InitRenderer(window *sdl.Window) *sdl.Renderer {
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic("Error loading the renderer")
	}
	return renderer
}

//Loads an image
func LoadMedia(path string) *sdl.Surface {
	loadSurface, err := sdlImg.Load(path)
	if err != nil {
		panic("Error loading image")
	}
	return loadSurface
}

//Frees space allocated
func Close(toFree ...freeable) {
	for _, i := range toFree {
		i.Free()
	}
}

//Initializes the texture
func initTexture(renderer *sdl.Renderer) aTexture {
	var texture aTexture
	texture.renderer = renderer
	return texture
}

//Texture Wrapper
type aTexture struct {
	height   int32
	width    int32
	texture  *sdl.Texture
	renderer *sdl.Renderer
	//
}

//Loads image into texture from specific path
func (texture *aTexture) loadImage(path string) {
	imgSurface := LoadMedia(path)
	var err error
	texture.texture, err = texture.renderer.CreateTextureFromSurface(imgSurface)
	if err != nil {
		panic("Error creating texture from surface")
	}
	texture.height = imgSurface.H
	texture.width = imgSurface.W
	imgSurface.Free()
}

//Renders texture
func (texture *aTexture) render(x int32, y int32, clip sdl.Rect) {
	renderQuad := sdl.Rect{X: x, Y: y, W: texture.width, H: texture.height}
	texture.renderer.Copy(texture.texture, nil, &renderQuad)
}
