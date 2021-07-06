package player

import (
	"math"
	"time"

	"github.com/sandbox/bullet"
	"github.com/sandbox/config"
	"github.com/sandbox/element"
	"github.com/veandco/go-sdl2/sdl"
)

type KeyboardMover struct {
	Container *element.Element
	Speed     float64
}

func NewKeybordMover(container *element.Element, speed float64) *KeyboardMover {
	return &KeyboardMover{
		Container: container,
		Speed:     speed}
}
func (mover *KeyboardMover) OnUpdate() error {
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_LEFT] == 1 {
		mover.Container.Position.X -= mover.Speed * config.Delta
		mover.Container.Collisions[0].X = mover.Container.Position.X
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		mover.Container.Position.X += mover.Speed * config.Delta
		mover.Container.Collisions[0].X = mover.Container.Position.X
	}

	if keys[sdl.SCANCODE_UP] == 1 {
		mover.Container.Position.Y -= mover.Speed * config.Delta
		mover.Container.Collisions[0].Y = mover.Container.Position.Y
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		mover.Container.Position.Y += mover.Speed * config.Delta
		mover.Container.Collisions[0].Y = mover.Container.Position.Y
	}

	return nil
}

func (mover *KeyboardMover) OnDraw(renderer *sdl.Renderer) error {
	return nil
}

func (m *KeyboardMover) OnCollition(other *element.Element) error {
	m.Container.Active = false
	return nil
}

type KeyboardShooter struct {
	Container    *element.Element
	CooldownTime time.Duration
	LastShot     time.Time
}

func NewKeyboardShooter(Container *element.Element, CooldownTime time.Duration) *KeyboardShooter {
	return &KeyboardShooter{
		Container:    Container,
		CooldownTime: CooldownTime,
	}
}

func (mover *KeyboardShooter) OnUpdate() error {
	keys := sdl.GetKeyboardState()
	pos := mover.Container.Position
	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(mover.LastShot) >= mover.CooldownTime {
			mover.shoot(pos.X, pos.Y-50)
			mover.LastShot = time.Now()
		}
	}
	return nil
}

func (mover *KeyboardShooter) OnDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *KeyboardShooter) shoot(x, y float64) {
	if b, ok := bullet.BulletFromPool(); ok {
		b.Active = true
		b.Position.X = x
		b.Position.Y = y
		b.Rotation = 270 * (math.Pi / 180)
	}
}

func (m *KeyboardShooter) OnCollition(other *element.Element) error {
	return nil
}
