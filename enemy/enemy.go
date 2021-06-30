package enemy

import (
	"github.com/sandbox/animator"
	"github.com/sandbox/element"
	"github.com/sandbox/score"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	enemySpeed = 5
	enemySize  = 60
	WenemySize = 200
	HenemySize = 120
)

func New(Renderer *sdl.Renderer, postion element.Vector) *element.Element {
	enemy := &element.Element{}
	enemy.Name = "enemy"

	enemy.Position = postion
	enemy.Rotation = 0
	enemy.Active = true
	enemy.Rotation = 20
	// sr := spriterenderer.New(enemy, Renderer, "sprites/enemy.bmp")
	// enemy.AddComponent(sr)
	mover := NewMover(enemy, enemySpeed)
	enemy.AddComponent(mover)

	idleSequence, err := animator.NewSequence("sprites/enemy/delta", 50, true, Renderer)
	if err != nil {
		panic(err)
	}

	destroySequence, err := animator.NewSequence("sprites/enemy/explosion", 50, false, Renderer)
	if err != nil {
		panic(err)
	}
	seqs := map[string]*animator.Sequence{
		"idle":    idleSequence,
		"destroy": destroySequence}

	animator := animator.NewAnimator(enemy, seqs, "idle")
	enemy.AddComponent(animator)

	col := element.Circle{
		X:      enemy.Position.X - 20,
		Y:      enemy.Position.Y - 20,
		Radius: 20}
	enemy.Collisions = append(enemy.Collisions, col)
	canbehit := NewCanBeHit(enemy)
	enemy.AddComponent(canbehit)
	return enemy
}

func Add(Renderer *sdl.Renderer) {
	for i := 1; i < 11; i++ {
		score.Board.EnemyCounter++
		enm := New(Renderer, element.Vector{X: float64(-1220 + (i * 100)), Y: 50})
		element.Elements = append(element.Elements, enm)
	}
	for i := 1; i < 11; i++ {
		score.Board.EnemyCounter++
		enm := New(Renderer, element.Vector{X: float64(-1260 + (i * 100)), Y: 100})
		element.Elements = append(element.Elements, enm)
	}
	for i := 1; i < 11; i++ {
		score.Board.EnemyCounter++
		enm := New(Renderer, element.Vector{X: float64(-1320 + (i * 100)), Y: 150})
		element.Elements = append(element.Elements, enm)
	}
}
