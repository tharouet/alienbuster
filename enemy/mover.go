package enemy

import (
	"github.com/sandbox/config"
	"github.com/sandbox/element"
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

	m.Container.Position.X += m.Speed * config.Delta
	m.Container.Position.Y += m.Speed / 20 * config.Delta
	if m.Container.Position.X > config.ScreenWidth+200 {
		m.Container.Position.X = -200
	}

	return nil
}

func (m *Mover) OnUpdate() error {

	m.Container.Collisions[0].X = m.Container.Position.X
	m.Container.Collisions[0].Y = m.Container.Position.Y
	return nil
}

func (m *Mover) OnCollition(other *element.Element) error {

	return nil
}
