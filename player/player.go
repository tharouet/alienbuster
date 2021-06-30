package player

import (
	"time"

	"github.com/sandbox/animator"
	"github.com/sandbox/config"
	"github.com/sandbox/element"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed        = 8
	playerSize         = 50
	playerShotCoolDown = time.Millisecond * 200
)

func New(Renderer *sdl.Renderer) *element.Element {
	player := &element.Element{}
	player.Position = element.Vector{
		X: config.ScreenHight/2 - 100,
		Y: config.ScreenWidth - 50}
	player.Active = true

	idleSequence, err := animator.NewSequence("sprites/player/idle", 20, true, Renderer)
	if err != nil {
		panic(err)
	}
	destroySequence, err := animator.NewSequence("sprites/enemy/explosion", 50, false, Renderer)
	if err != nil {
		panic(err)
	}

	mover := NewKeybordMover(player, playerSpeed)
	player.AddComponent(mover)

	shooter := NewKeyboardShooter(player, playerShotCoolDown)
	player.AddComponent(shooter)

	seqs := map[string]*animator.Sequence{
		"idle":    idleSequence,
		"destroy": destroySequence}

	animator := animator.NewAnimator(player, seqs, "idle")
	player.AddComponent(animator)

	col := element.Circle{
		X:      player.Position.X - 20,
		Y:      player.Position.Y - 20,
		Radius: 20}

	player.Collisions = append(player.Collisions, col)
	canbehit := NewCanBeHit(player)
	player.AddComponent(canbehit)

	player.Name = "player"
	return player
}
