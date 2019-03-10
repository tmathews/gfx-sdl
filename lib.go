package gfx

import (
	"bytes"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func SurfaceFromFile(filename string) (*sdl.Surface, error) {
	return img.Load(filename)
}

func SurfaceFromBuf(buf []byte) (*sdl.Surface, error) {
	rw, err := sdl.RWFromMem(buf)
	if err != nil {
		return nil, err
	}
	return img.LoadRW(rw, true)
}

func SurfaceFromBufString(a string) (*sdl.Surface, error) {
	buf := bytes.NewBufferString(a)
	return SurfaceFromBuf(buf.Bytes())
}

func FontFromBufString(a string, size int) (*ttf.Font, error) {
	rw, err := sdl.RWFromMem(bytes.NewBufferString(a).Bytes())
	if err != nil {
		return nil, err
	}
	return ttf.OpenFontRW(rw, 1, size)
}