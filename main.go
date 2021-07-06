package main

import (
	"fmt"
	"log"
	"time"

	"github.com/sandbox/bullet"
	"github.com/sandbox/config"
	"github.com/sandbox/element"
	"github.com/sandbox/enemy"
	"github.com/sandbox/player"
	"github.com/sandbox/scene"
	"github.com/sandbox/score"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	if err := ttf.Init(); err != nil {
		log.Println("Initializing Error:", err)
		return
	}
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Println("Initializing Error:", err)
		return
	}
	window, err := sdl.CreateWindow(
		"Alien Buster",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		config.ScreenWidth, config.ScreenHight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		log.Println("Window Creation Error:", err)
		return
	}

	defer window.Destroy()

	Renderer, _ := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	defer Renderer.Destroy()

	scene := scene.New(Renderer)
	element.Elements = append(element.Elements, scene)

	plr := player.New(Renderer)
	element.Elements = append(element.Elements, plr)

	bullet.InitBulletPoll(Renderer)
	for {

		//WriteChoices(Renderer)

		if score.Board.Lives > 0 && !plr.Active {
			plr = player.New(Renderer)
			element.Elements = append(element.Elements, plr)
		}

		// add enemy and level up
		if score.Board.EnemyCounter == 0 {
			enemy.Add(Renderer)
			score.Board.Level++
		}

		framStartTime := time.Now()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		Renderer.SetDrawColor(0, 0, 0, 255)
		Renderer.Clear()
		for _, elm := range element.Elements {

			if elm.Active == true {
				err = elm.Update()
				if err != nil {
					fmt.Println("Update Error: ", err)
					return
				}
				err = elm.Draw(Renderer)
				if err != nil {
					fmt.Println("Draw Error: ", err)
					return
				}
			}

		}

		if err := element.CheckCollisions(); err != nil {
			fmt.Println("Collision Error: ", err)
			return
		}
		WriteChoices(Renderer)
		Renderer.Present()
		config.Delta = time.Since(framStartTime).Seconds() * config.TargetTicksPerSecond

	}
}

func WriteChoices(renderer *sdl.Renderer) {
	font, _ := ttf.OpenFont("data.ttf", 18)
	font.SetOutline(0)
	Score := fmt.Sprintf("Level: %v   | Lives:   %v   | Enemies Remaining: %v", score.Board.Level, score.Board.Lives, score.Board.EnemyCounter)
	surface, _ := font.RenderUTF8Solid(Score, sdl.Color{11, 156, 49, 255})
	texture, _ := renderer.CreateTextureFromSurface(surface)
	renderer.Copy(texture, nil, &sdl.Rect{W: surface.W, H: surface.H})
	font.Close()
	surface.Free()
	texture.Destroy()
}
