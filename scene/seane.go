package scene

import (
	"github.com/alienbuster/element"
	"github.com/alienbuster/spriterenderer"
	"github.com/veandco/go-sdl2/sdl"
)

func New(Renderer *sdl.Renderer) *element.Element {
	scene := &element.Element{}
	scene.Position = element.Vector{
		X: 300,
		Y: 400}
	scene.Active = true
	sr := spriterenderer.New(scene, Renderer, "sprites/bkl1.bmp")
	scene.AddComponent(sr)

	scene.Name = "scene"
	return scene
}
