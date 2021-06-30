package bullet

import (
	"github.com/sandbox/element"
	"github.com/sandbox/spriterenderer"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	BulletSpeed = 20
	BulletSize  = 25
	WBulletSize = 25
	HBulletSize = 50
)

func New(R *sdl.Renderer) *element.Element {
	bullet := &element.Element{Active: false}
	sr := spriterenderer.New(bullet, R, "sprites/missle.bmp")
	bullet.AddComponent(sr)
	mover := NewMover(bullet, BulletSpeed)
	bullet.AddComponent(mover)
	col := element.Circle{
		X:      bullet.Position.X,
		Y:      bullet.Position.Y,
		Radius: 4}
	bullet.Collisions = append(bullet.Collisions, col)
	bullet.Name = "bullet"
	return bullet
}

var BulletPool []*element.Element

func InitBulletPoll(Renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		bul := New(Renderer)
		element.Elements = append(element.Elements, bul)
		BulletPool = append(BulletPool, bul)
	}

}

func BulletFromPool() (*element.Element, bool) {
	for _, bull := range BulletPool {
		if !bull.Active {
			return bull, true
		}
	}
	return nil, false
}
