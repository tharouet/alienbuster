package enemy

import (
	"github.com/sandbox/animator"
	"github.com/sandbox/config"
	"github.com/sandbox/element"
	"github.com/sandbox/score"
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
	if c.Container.Position.Y > config.ScreenHight+200 {
		c.Container.Active = false
		score.Board.EnemyCounter--
	}
	return nil
}

func (c *CanBeHit) OnCollition(other *element.Element) error {
	if c.Container.Name == "enemy" {
		c.Animator.SetSequnce("destroy")
	}
	return nil
}
