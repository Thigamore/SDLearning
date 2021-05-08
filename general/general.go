package general

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Freeable interface {
	Free()
}

//----------------------------------Basic Functions---------------------------

//Prepares Window, Window Surface
func InitWindow(screenWidth int32, screenHeight int32) (*sdl.Window, *sdl.Surface) {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic("Error initaializing sdl")
	}
	err = ttf.Init()
	if err != nil {
		panic("Error initaializing ttf")
	}
	err = img.Init(img.INIT_PNG | img.INIT_JPG)
	if err != nil {
		panic("Error initaializing img")
	}

	window, err := sdl.CreateWindow("SDL", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, 0)
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
	loadSurface, err := img.Load(path)
	if err != nil {
		panic("Error loading image")
	}
	loadSurface.SetBlendMode(sdl.BLENDMODE_BLEND)
	return loadSurface
}

//Frees space allocated
func Close(toFree ...Freeable) {
	for _, i := range toFree {
		i.Free()
	}
}

func CloseAll(toFree []Freeable, renderer *sdl.Renderer, window *sdl.Window) {
	for _, i := range toFree {
		i.Free()
	}

	renderer.Destroy()
	window.Destroy()

	ttf.Quit()
	img.Quit()
	sdl.Quit()
}

//-----------------------------Frame Rate Manager------------------------------

//Initializes a frame rate manager
func InitFrameRateManager(frameCap uint32) *frameRateManager {
	return &frameRateManager{frameCap: frameCap, timer: InitTimer(), lastTic: 0}
}

//A struct that is used to cap framerate to make sure that it doesn't depass a certain limit
type frameRateManager struct {
	frameCap uint32
	timer    *timer
	lastTic  uint32
}

//Starts the manager
func (frameMan *frameRateManager) Start() {
	frameMan.timer.Start()
	frameMan.lastTic = frameMan.timer.Run()
}

//Runs the manager and delays the program if it is running too fast
func (frameMan *frameRateManager) Run(toDisplay bool) {
	currentTic := frameMan.timer.Run()
	if (currentTic - frameMan.lastTic) <= (1000 / frameMan.frameCap) {
		sdl.Delay((1000 / frameMan.frameCap) - (currentTic - frameMan.lastTic))
	}

	if toDisplay {
		fmt.Println(1000 / float64((frameMan.timer.Run() - frameMan.lastTic)))
	}

	frameMan.lastTic = frameMan.timer.Run()
}

//-----------------------------Dimensions--------------------------------------

//Struct that holds the width and height of anything
type Dimension struct {
	W, H int32
}

//-----------------------------Texture Wrapper---------------------------------

//Initializes the texture
func InitTexture(renderer *sdl.Renderer) *ATexture {
	var texture ATexture
	texture.Renderer = renderer
	return &texture
}

//Texture Wrapper
type ATexture struct {
	Height   int32
	Width    int32
	Texture  *sdl.Texture
	Renderer *sdl.Renderer
	font     *ttf.Font
	color    *sdl.Color
}

//Loads image into texture from specific path
func (texture *ATexture) LoadImage(path string, keyColor *sdl.Color) {
	imgSurface := LoadMedia(path)
	if keyColor != nil {
		imgSurface.SetColorKey(true, sdl.MapRGB(imgSurface.Format, keyColor.R, keyColor.G, keyColor.B))
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
func (texture *ATexture) Render(x int32, y int32, clip *sdl.Rect) {
	var renderQuad sdl.Rect
	if clip != nil {
		renderQuad = sdl.Rect{X: x, Y: y, W: clip.W, H: clip.H}
	} else {
		renderQuad = sdl.Rect{X: x, Y: y, W: texture.Width, H: texture.Height}
	}

	texture.Renderer.Copy(texture.Texture, clip, &renderQuad)
}

//Renders texture and flips
func (texture *ATexture) RenderFlip(x int32, y int32, clip *sdl.Rect, angle float64, center *sdl.Point, flip sdl.RendererFlip) {
	var renderQuad sdl.Rect
	if clip != nil {
		renderQuad = sdl.Rect{X: x, Y: y, W: clip.W, H: clip.H}
	} else {
		renderQuad = sdl.Rect{X: x, Y: y, W: texture.Width, H: texture.Height}
	}

	texture.Renderer.CopyEx(texture.Texture, clip, &renderQuad, angle, center, flip)
}

//Modulates texture
func (texture *ATexture) SetColor(red uint8, green, blue uint8) {
	err := texture.Texture.SetColorMod(red, green, blue)
	if err != nil {
		panic("Error setting color modulation")
	}
}

//Sets alpha modulation
func (texture *ATexture) SetAlpha(alpha uint8) {
	texture.Texture.SetAlphaMod(alpha)
}

//Sets blending mode
func (texture *ATexture) SetBlendMode(blending sdl.BlendMode) {
	texture.Texture.SetBlendMode(blending)
}

//Creates image from font string
func (texture *ATexture) LoadText(text string, color sdl.Color) {
	texture.Texture.Destroy()

	textSurface, err := texture.font.RenderUTF8Solid(text, color)
	if err != nil {
		panic("Error creating surface from text")
	}

	newTexture, err := texture.Renderer.CreateTextureFromSurface(textSurface)
	if err != nil {
		panic("Error creating texture from text surface")
	}

	texture.Width = textSurface.W
	texture.Height = textSurface.H

	textSurface.Free()

	texture.Texture = newTexture
}

//Sets font
func (texture *ATexture) SetFont(font *ttf.Font) {
	texture.font = font
}

//Sets color
func (texture *ATexture) SetFontColor(color *sdl.Color) {
	texture.color = color
}

//Copies the object
func (texture *ATexture) Copy() *ATexture {
	newTexture := *texture
	return &newTexture
}

//Destroys and frees everything in texture
func (texture *ATexture) Free() {
	texture.Texture.Destroy()
	texture.font.Close()
}

//------------------------------------Button Class----------------------

//Initializes a new button
func InitButton(position *sdl.Point, sprite *ATexture, spritePos sdl.Rect, width int32, height int32) Button {
	var newButton Button
	newButton.position = position
	newButton.spriteSheet = sprite
	newButton.height = height
	newButton.width = width
	newButton.spritePos = spritePos
	return newButton
}

type Button struct {
	position    *sdl.Point
	spriteSheet *ATexture
	spritePos   sdl.Rect
	width       int32
	height      int32
}

func (button *Button) SetPosition(x int32, y int32) {
	button.position.X = x
	button.position.Y = y
}

func (button *Button) HandleEvent(e sdl.Event) string {
	var output string
	if e.GetType() == sdl.MOUSEMOTION || e.GetType() == sdl.MOUSEBUTTONDOWN || e.GetType() == sdl.MOUSEBUTTONUP {
		x, y, _ := sdl.GetMouseState()
		inside := true
		if x < button.position.X || x > (button.position.X+button.width) || y < button.position.Y || y > (button.position.Y+button.height) {
			inside = false
		}
		if !inside {
			output = "out"
		} else {
			switch e.GetType() {
			case sdl.MOUSEMOTION:
				output = "over"
			case sdl.MOUSEBUTTONDOWN:
				output = "down"
			case sdl.MOUSEBUTTONUP:
				output = "up"
			}
		}
	}
	return output
}

//Sets sprite position
func (button *Button) SetSpritePos(pos sdl.Rect) {
	button.spritePos = pos
}

//Renders button
func (button *Button) Render() {
	button.spriteSheet.Render(button.position.X, button.position.Y, &button.spritePos)
}

//------------------------------Timer Class-----------------------------

//Initializes a new timer
func InitTimer() *timer {
	timer := timer{ticksStart: 0, ticksPaused: 0, paused: false, started: false}
	return &timer
}

//Timer class
type timer struct {
	ticksStart  uint32
	ticksPaused uint32
	paused      bool
	started     bool
}

//Starts the timer
func (timer *timer) Start() {
	if !(timer.started) {
		timer.started = true
		timer.ticksStart = sdl.GetTicks()
	} else {
		timer.started = false
		timer.paused = false
		timer.ticksStart = 0
		timer.ticksPaused = 0
	}

}

//Pauses the timer
func (timer *timer) Pause() {
	if !(timer.paused) && timer.started {
		timer.paused = true
		timer.ticksPaused = sdl.GetTicks()
	} else if timer.paused && timer.started {
		timer.paused = false
		timer.ticksPaused = 0
	}
}

//Gets the amount of time the timer has been started
func (timer *timer) Run() uint32 {
	if timer.paused {
		return timer.ticksPaused - timer.ticksStart
	} else if timer.started {
		return sdl.GetTicks() - (timer.ticksPaused + timer.ticksStart)
	} else {
		return 0
	}
}

//Checks if the timer is running
func (timer *timer) IsStarted() bool {
	return timer.started
}

//Checks if the timer is paused
func (timer *timer) IsPaused() bool {
	return timer.paused
}

//-----------------------------------Velocity-----------------------
type Velocity struct {
	X, Y int32
}

//-----------------------------------pEntity-------------------------

//Initializes an pEntity
func InitPEntity(texture *ATexture, location sdl.Point, initVelocity Velocity, dimension Dimension) *pEntity {
	newpEntity := pEntity{location: location, dimension: dimension, velocity: initVelocity, texture: texture}
	return &newpEntity
}

//An pEntity struct, p stands for player
//Entities are objects of some kind that can be moved by player
type pEntity struct {
	location  sdl.Point
	dimension Dimension
	velocity  Velocity
	texture   *ATexture
}

//Handles a keyboard event
func (pEntity *pEntity) HandleEvent(t *sdl.KeyboardEvent, velocityChange uint32) {
	switch t.Keysym.Sym {
	case sdl.K_w:
		pEntity.velocity.Y -= 10
	case sdl.K_s:
		pEntity.velocity.Y += 10
	case sdl.K_d:
		pEntity.velocity.X += 10
	case sdl.K_a:
		pEntity.velocity.X -= 10
	case sdl.K_SPACE:
		pEntity.velocity.X = 0
		pEntity.velocity.Y = 0
	}
}

//Moves the entity every frame
func (pEntity *pEntity) Move(screenDim Dimension) {
	if ((pEntity.location.X + pEntity.velocity.X) > screenDim.W) || ((pEntity.location.X + pEntity.velocity.X + pEntity.dimension.W) < 0) {
		pEntity.location.X = screenDim.W - pEntity.location.X
	}
	if ((pEntity.location.Y + pEntity.velocity.Y) > screenDim.H) || ((pEntity.location.Y + pEntity.velocity.Y + pEntity.dimension.H) < 0) {
		pEntity.location.Y = screenDim.H - pEntity.location.Y
	}

	pEntity.location.X += pEntity.velocity.X
	pEntity.location.Y += pEntity.velocity.Y
}

//Renders the enitity
func (pEntity *pEntity) Render() {
	pEntity.texture.Render(pEntity.location.X, pEntity.location.Y, nil)
}
