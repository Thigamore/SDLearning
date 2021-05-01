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
func InitRenderer(window *sdl.Window, vSync bool) *sdl.Renderer {
	var renderer *sdl.Renderer
	var err error
	if vSync {
		renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	} else {
		renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	}
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
	loadSurface.SetBlendMode(sdl.BLENDMODE_BLEND)
	return loadSurface
}

//Frees space allocated
func Close(toFree ...freeable) {
	for _, i := range toFree {
		i.Free()
	}
}

//Initializes the texture
func InitTexture(renderer *sdl.Renderer) aTexture {
	var texture aTexture
	texture.Renderer = renderer
	return texture
}

//Texture Wrapper
type aTexture struct {
	Height   int32
	Width    int32
	Texture  *sdl.Texture
	Renderer *sdl.Renderer
}

//Loads image into texture from specific path
func (texture *aTexture) LoadImage(path string, toKey bool) {
	imgSurface := LoadMedia(path)
	if toKey {
		imgSurface.SetColorKey(true, sdl.MapRGB(imgSurface.Format, 0, 0xFF, 0xFF))
	}
	var err error
	texture.Texture, err = texture.Renderer.CreateTextureFromSurface(imgSurface)
	if err != nil {
		panic("Error creating texture from surface")
	}
	texture.Height = imgSurface.H
	texture.Width = imgSurface.W
	imgSurface.Free()
}

//Renders texture
func (texture *aTexture) Render(x int32, y int32, clip *sdl.Rect) {
	var renderQuad sdl.Rect
	if clip != nil {
		renderQuad = sdl.Rect{X: x, Y: y, W: clip.W, H: clip.H}
	} else {
		renderQuad = sdl.Rect{X: x, Y: y, W: texture.Width, H: texture.Height}
	}
	texture.Renderer.Copy(texture.Texture, clip, &renderQuad)
}

//Modulates texture
func (texture *aTexture) SetColor(red uint8, green, blue uint8) {
	err := texture.Texture.SetColorMod(red, green, blue)
	if err != nil {
		panic("Error setting color modulation")
	}
}

//Sets alpha modulation
func (texture *aTexture) SetAlpha(alpha uint8) {
	texture.Texture.SetAlphaMod(alpha)
}

//Sets blending mode
func (texture *aTexture) SetBlendMode(blending sdl.BlendMode) {
	texture.Texture.SetBlendMode(blending)
}
