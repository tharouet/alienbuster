package spriterenderer

import (
	"github.com/sandbox/builder"
	"github.com/sandbox/element"
	"github.com/veandco/go-sdl2/sdl"
)

type SpriteRenderer struct {
	Container     *element.Element
	Tex           *sdl.Texture
	Width, Height float64
}

func New(Container *element.Element, renderer *sdl.Renderer, filePath string) *SpriteRenderer {
	tex, err := builder.LoadTextureFromBMP(filePath, renderer)
	_, _, width, height, err := tex.Query()
	if err != nil {
		panic(err)
	}
	return &SpriteRenderer{
		Container: Container,
		Tex:       tex,
		Height:    float64(height),
		Width:     float64(width),
	}
}

func (sr *SpriteRenderer) OnDraw(Renderer *sdl.Renderer) error {
	return builder.DrawTexture(
		sr.Tex,
		sr.Container.Position,
		sr.Container.Rotation,
		Renderer)
}

func (sr *SpriteRenderer) OnUpdate() error {
	return nil

}

func (sr *SpriteRenderer) OnCollition(other *element.Element) error {
	return nil
}

func TextureFromBmp(Renderer *sdl.Renderer, filepath string) (tex *sdl.Texture) {

	img, err := sdl.LoadBMP(filepath)
	if err != nil {
		panic(err)
	}
	defer img.Free()

	tex, err = Renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(err)
	}
	return
}
