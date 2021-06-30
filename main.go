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
)

func main() {
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

	// add another player unit
	plr2 := player.New(Renderer)
	plr2.Position.X = plr2.Position.X + 50
	element.Elements = append(element.Elements, plr2)

	// add another player unit
	plr3 := player.New(Renderer)
	plr3.Position.X = plr3.Position.X + 25
	plr3.Position.Y = plr3.Position.Y - 50
	element.Elements = append(element.Elements, plr3)

	bullet.InitBulletPoll(Renderer)

	for {

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

		Renderer.Present()
		config.Delta = time.Since(framStartTime).Seconds() * config.TargetTicksPerSecond
	}
}
