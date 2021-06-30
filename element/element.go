package element

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type Vector struct {
	X, Y float64
}

type Component interface {
	OnUpdate() error
	OnDraw(renderder *sdl.Renderer) error
	OnCollition(other *Element) error
}

type Element struct {
	Position   Vector
	Rotation   float64
	Active     bool
	Collisions []Circle
	Components []Component
	Name       string
}

func (e *Element) Draw(R *sdl.Renderer) error {
	for _, comp := range e.Components {
		err := comp.OnDraw(R)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Element) Update() error {
	for _, comp := range e.Components {
		err := comp.OnUpdate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Element) AddComponent(new Component) {
	for _, existing := range e.Components {
		if reflect.TypeOf(new) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf(
				"attempt to add new component with existing type of %v",
				reflect.TypeOf(new)))
		}
	}
	e.Components = append(e.Components, new)
}

func (e *Element) GetComponent(withType Component) Component {
	typ := reflect.TypeOf(withType)
	for _, comp := range e.Components {
		if reflect.TypeOf(comp) == typ {
			return comp
		}
	}
	panic(fmt.Sprintf(
		"attempt to get component with existing type of %v has failed",
		typ))

}

var Elements []*Element

func (e *Element) Collision(other *Element) error {
	for _, comp := range e.Components {
		err := comp.OnCollition(other)
		if err != nil {
			return err
		}
	}
	return nil
}
