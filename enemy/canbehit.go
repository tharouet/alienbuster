package enemy

import (
	"github.com/alienbuster/animator"
	"github.com/alienbuster/element"
	"github.com/alienbuster/score"
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
		score.Board.EnemyCounter--
		c.Animator.Container.Active = false
	}
	return nil
}

func (c *CanBeHit) OnCollition(other *element.Element) error {
	if c.Container.Name == "enemy" {
		c.Animator.SetSequnce("destroy")
	}
	if c.Container.Name == "player" {
		c.Animator.SetSequnce("destroy")
	}
	return nil
}
