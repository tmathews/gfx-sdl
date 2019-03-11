package gfx

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Text struct {
	Content string
	Color   sdl.Color
	Texture *sdl.Texture
	Surface *sdl.Surface
	Font    *ttf.Font
	Size    int
	X       int32
	Y       int32
}

func (t *Text) Draw(r *sdl.Renderer) {
	if t.Surface == nil {
		return
	}
	if t.Texture == nil {
		var err error
		t.Texture, err = r.CreateTextureFromSurface(t.Surface)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	dst := sdl.Rect{t.X, t.Y, t.Surface.W, t.Surface.H}
	r.Copy(t.Texture, &t.Surface.ClipRect, &dst)
}

func (t *Text) Free() {
	if t.Texture != nil {
		t.Texture.Destroy()
	}
	if t.Surface != nil {
		t.Surface.Free()
	}
	t.Font.Close()
}

func (t *Text) Render() error {
	// Free our old texture so the renderer will recreate
	if t.Texture != nil {
		t.Texture.Destroy()
		t.Texture = nil
	}
	if t.Surface != nil {
		t.Surface.Free()
	}
	var err error
	// TODO add ability to set render type: solid or blended AKA hard or smooth
	t.Surface, err = t.Font.RenderUTF8Blended(t.Content, t.Color)
	if err != nil {
		return err
	}
	return nil
}

func InitText() error {
	return ttf.Init()
}

func NewText(str string, size int, fontPath string) (*Text, error) {
	font, err := ttf.OpenFont(fontPath, size)
	if err != nil {
		return nil, err
	}
	t, err := newTextWithFont(str, font)
	if err != nil {
		return nil, err
	}
	t.Size = size
	return t, nil
}

func NewTextFromBufString(str string, size int, a string) (*Text, error) {
	font, err := FontFromBufString(a, size)
	if err != nil {
		return nil, err
	}
	return newTextWithFont(str, font)
}

func newTextWithFont(str string, font *ttf.Font) (*Text, error) {
	t := &Text{
		Content: str,
		Font: font,
	}
	return t, t.Render()
}