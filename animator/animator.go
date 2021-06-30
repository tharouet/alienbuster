package animator

import (
	"io/ioutil"
	"path"
	"time"

	"github.com/sandbox/builder"
	"github.com/sandbox/element"
	"github.com/veandco/go-sdl2/sdl"
)

type Animator struct {
	Container       *element.Element
	Sequences       map[string]*Sequence
	Current         string
	Lastframechange time.Time
	Finished        bool
}

type Sequence struct {
	Textures   []*sdl.Texture
	Frame      int
	SampleRate float64
	Loop       bool
}

func NewAnimator(
	container *element.Element,
	sequences map[string]*Sequence,
	defaultSequence string) *Animator {
	a := &Animator{
		Container:       container,
		Sequences:       sequences,
		Current:         defaultSequence,
		Lastframechange: time.Now(),
	}

	return a
}

func (A *Animator) OnUpdate() error {
	Sequence := A.Sequences[A.Current]
	frameInterval := float64(time.Second) / Sequence.SampleRate
	if time.Since(A.Lastframechange) >= time.Duration(frameInterval) {
		A.Finished = Sequence.NextFrame()
		A.Lastframechange = time.Now()

	}
	return nil
}
func (A *Animator) OnDraw(renderder *sdl.Renderer) error {
	tex := A.Sequences[A.Current].texture()

	return builder.DrawTexture(
		tex,
		A.Container.Position,
		A.Container.Rotation,
		renderder)
}

func (A *Animator) OnCollition(other *element.Element) error {
	return nil
}

func (seq *Sequence) texture() *sdl.Texture {
	return seq.Textures[seq.Frame]
}

func (s *Sequence) NextFrame() bool {
	if s.Frame == len(s.Textures)-1 {
		if s.Loop {
			s.Frame = 0
		} else {
			return true
		}
	} else {
		s.Frame++
	}
	return false
}

func NewSequence(filepath string, sampleRate float64, loop bool, renderer *sdl.Renderer) (*Sequence, error) {
	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		return nil, err
	}
	var Seq Sequence
	for _, file := range files {
		name := path.Join(filepath, file.Name())
		tex, err := builder.LoadTextureFromBMP(name, renderer)
		if err != nil {
			return nil, err
		}
		Seq.Textures = append(Seq.Textures, tex)
	}
	Seq.SampleRate = sampleRate
	Seq.Loop = loop
	return &Seq, nil
}

func (a *Animator) SetSequnce(name string) {
	a.Current = name
	a.Lastframechange = time.Now()
}
