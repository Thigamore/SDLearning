package main

import (
	"github.com/thigamore/SDL/general"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

func SDLMusicIsh() {
	sdl.Init(sdl.INIT_AUDIO | sdl.INIT_EVENTS | sdl.INIT_VIDEO)
	mix.Init(mix.INIT_MP3)
	img.Init(img.INIT_PNG)
	err := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 2048)
	if err != nil {
		panic(sdl.GetError())
	}

	const SCREEN_WIDTH = 500
	const SCREEN_HEIGHT = 500
	window, _ := general.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT)
	renderer := general.InitRenderer(window, true)
	prompt := general.InitTexture(renderer)
	music, err := mix.LoadMUS("beat.wav")
	if err != nil {
		panic(sdl.GetError())
	}
	low, err := mix.LoadWAV("low.wav")
	if err != nil {
		panic(sdl.GetError())
	}
	medium, err := mix.LoadWAV("medium.wav")
	if err != nil {
		panic(sdl.GetError())
	}
	high, err := mix.LoadWAV("high.wav")
	if err != nil {
		panic(sdl.GetError())
	}
	scratch, err := mix.LoadWAV("scratch.wav")
	if err != nil {
		panic(sdl.GetError())
	}

	prompt.LoadImage("prompt.png", nil)

	running := true
	var e sdl.Event

	for running {
		e = sdl.PollEvent()
		for e != nil {
			switch t := e.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				switch t.Keysym.Sym {
				case sdl.K_1:
					high.Play(-1, 0)
				case sdl.K_2:
					medium.Play(-1, 0)
				case sdl.K_3:
					low.Play(-1, 0)
				case sdl.K_4:
					scratch.Play(-1, 0)
				case sdl.K_9:
					if !(mix.PlayingMusic()) {
						music.Play(-1)
					} else {
						if mix.PausedMusic() {
							mix.ResumeMusic()
						} else {
							mix.PauseMusic()
						}
					}
				case sdl.K_0:
					mix.HaltMusic()
				}
			}
			e = sdl.PollEvent()
		}
		renderer.Clear()
		renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)

		prompt.Render(0, 0, nil)

		renderer.Present()
	}
}
