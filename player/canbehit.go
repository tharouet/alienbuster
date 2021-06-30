package player

import (
	"github.com/sandbox/animator"
	"github.com/sandbox/element"
	"github.com/veandco/go-sdl2/sdl"
)

type CanBeHit struct {
	Container *element.Element
	Animator  *animator.Animator
}

func NewCanBeHit(Container *element.Element) *CanBeHit {
	return &CanBeHit{
		Container: Container,
		Animator:  Container.GetComponent(&animator.Animator{}).(*animator.Animator),
	}

}

func (c *CanBeHit) OnDraw(R *sdl.Renderer) error {
	return nil
}

func (c *CanBeHit) OnUpdate() error {
	if c.Animator.Finished && c.Animator.Current == "destroy" {
		c.Animator.Container.Active = false
	}
	return nil
}

func (c *CanBeHit) OnCollition(other *element.Element) error {
	if c.Container.Name == "player" {
		c.Animator.SetSequnce("destroy")
	}
	return nil
}
