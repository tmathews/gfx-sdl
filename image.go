package gfx

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type ImageMode int8
type ImageAlign int8

const (
	ImNone    = ImageMode(iota) // Will output the image as is at x, y
	ImFill                      // Will stretch to the fit bounds
	ImContain                   // The image will fit inside to the maximum bounds
	ImCover                     // All space will be fitted into
	ImCrop                      // Will crop to bounds
)

const (
	ImTopLeft = ImageAlign(iota)
	ImTopCenter
	ImTopRight
	ImCenterLeft
	ImCenter
	ImCenterRight
	ImBottomLeft
	ImBottomCenter
	ImBottomRight
)

type Image struct {
	Texture *sdl.Texture
	Surface *sdl.Surface
	Width   int32
	Height  int32
	X       int32
	Y       int32
	Mode    ImageMode
	Align   ImageAlign
}

func (i *Image) Draw(r *sdl.Renderer) {
	var dst sdl.Rect
	src := sdl.Rect{0, 0, i.Surface.W, i.Surface.H}

	switch i.Mode {
	case ImNone:
		dst = sdl.Rect{W: i.Surface.W, H: i.Surface.H}
	case ImFill:
		dst = sdl.Rect{W: i.Width, H: i.Height}
	case ImContain:
		dst = getContainRect(i.Surface.W, i.Surface.H, i.Width, i.Height, i.Align)
	case ImCover:
		dst.W = i.Width
		dst.H = i.Height
		src = getCoverRect(i.Surface.W, i.Surface.H, i.Width, i.Height, i.Align)
	case ImCrop:
		dst.W = i.Width
		dst.H = i.Height
		src = getCropRect(i.Surface.W, i.Surface.H, i.Width, i.Height, i.Align)
	}
	dst.X += i.X
	dst.Y += i.Y
	r.Copy(i.Texture, &src, &dst)
}

func (i *Image) Free() {
	i.Surface.Free()
	i.Texture.Destroy()
}

func (i *Image) Load(r *sdl.Renderer, filename string) error {
	var err error
	surface, err := img.Load(filename)
	if err != nil {
		return err
	}
	texture, err := r.CreateTextureFromSurface(surface)
	if err != nil {
		return err
	}
	i.Surface = surface
	i.Texture = texture
	return nil
}

func NewImage(r *sdl.Renderer, filename string) (*Image, error) {
	var i Image
	if err := i.Load(r, filename); err != nil {
		return nil, err
	}
	return &i, nil
}

func getContainRect(sw, sh, w, h int32, align ImageAlign) sdl.Rect {
	nw := sw
	nh := sh
	var x int32
	var y int32
	if sw > w {
		nw = w
		nh = int32((float32(nw) * float32(sh)) / float32(sw))
	}
	if nh > h {
		nh = h
		nw = int32((float32(nh) * float32(sw)) / float32(sh))
	}

	switch align {
	case ImCenter:
		x = (w / 2) - (nw / 2)
		y = (h / 2) - (nh / 2)
	}

	return sdl.Rect{
		W: nw,
		H: nh,
		X: x,
		Y: y,
	}
}

func getCoverRect(sw, sh, w, h int32, align ImageAlign) sdl.Rect {
	if w >= sw || h >= sh {
		return getContainRect(w, h, sw, sh, align)
	}

	nw := sw
	nh := sh
	var x int32
	var y int32

	nh = int32((float32(sw) / float32(w)) * float32(h))
	nw = int32((float32(sh) / float32(h)) * float32(w))

	switch align {
	case ImCenter:
		x = (sw - nw) / 2
		y = (sh - nh) / 2
	}

	return sdl.Rect{
		W: nw,
		H: nh,
		X: x,
		Y: y,
	}
}

func getCropRect(sw, sh, w, h int32, align ImageAlign) sdl.Rect {
	// TODO finish me
	return sdl.Rect{
		W: w,
		H: h,
	}
}