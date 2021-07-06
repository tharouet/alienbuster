package builder

import (
	"github.com/alienbuster/element"
	"github.com/veandco/go-sdl2/sdl"
)

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

func DrawTexture(
	Texture *sdl.Texture,
	Postion element.Vector,
	Rotation float64,
	Renderer *sdl.Renderer) error {

	_, _, width, height, err := Texture.Query()
	if err != nil {
		return err
	}
	// convert coordinate to top left of sprite
	Postion.X -= float64(width / 2.0)
	Postion.Y -= float64(height / 2.0)

	return Renderer.CopyEx(
		Texture,
		&sdl.Rect{X: 0, Y: 0, W: int32(width), H: int32(height)},
		&sdl.Rect{
			X: int32(Postion.X),
			Y: int32(Postion.Y),
			W: int32(width / 1),
			H: int32(height / 1)},
		Rotation,
		&sdl.Point{X: int32(width) / 1, Y: int32(height) / 1},
		sdl.FLIP_NONE)

}

func LoadTextureFromBMP(filename string, renderer *sdl.Renderer) (*sdl.Texture, error) {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		return nil, err
	}
	defer img.Free()
	Tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		return nil, err
	}
	return Tex, err

}
