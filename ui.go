package gfx

import (
	"github.com/veandco/go-sdl2/sdl"
	"image/color"
)

type Drawable interface {
	Draw(renderer *sdl.Renderer)
}

type Resource interface {
	Free()
}

type Container struct {
	Width  int32
	Height int32
	Color  color.RGBA
}

func (c *Container) Draw(r *sdl.Renderer) {
	r.SetDrawColor(c.Color.R, c.Color.G, c.Color.B, c.Color.A)
	r.FillRect(&sdl.Rect{0, 0, c.Width, c.Height})
}
