package bullet

import (
	"math"

	"github.com/alienbuster/config"
	"github.com/alienbuster/element"
	"github.com/veandco/go-sdl2/sdl"
)

type Mover struct {
	Container *element.Element
	Speed     float64
}

func NewMover(Container *element.Element, Speed float64) *Mover {
	return &Mover{
		Container: Container,
		Speed:     Speed,
	}
}

func (m *Mover) OnDraw(Renderer *sdl.Renderer) error {
	return nil
}

func (m *Mover) OnUpdate() error {

	m.Container.Position.X += BulletSpeed * math.Cos(m.Container.Rotation) * config.Delta
	m.Container.Position.Y += BulletSpeed * math.Sin(m.Container.Rotation) * config.Delta

	if m.Container.Position.X > config.ScreenWidth || m.Container.Position.X < 0 ||
		m.Container.Position.Y > config.ScreenHight || m.Container.Position.Y < 0 {
		m.Container.Active = false
	}
	m.Container.Collisions[0].X = m.Container.Position.X
	m.Container.Collisions[0].Y = m.Container.Position.Y
	return nil
}

func (m *Mover) OnCollition(other *element.Element) error {
	m.Container.Active = false
	return nil
}
