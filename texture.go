package main

import (
	"errors"
	"image"
	"image/color"
	"github.com/banthar/gl"
)

var (
	CustomColorModels = make(map[color.Model]*GLColorModel)
)

type EngineColorModel interface {
	color.Model
	Data() interface{}
}

type GLColorModel struct {
	InternalFormat int
	Type gl.GLenum
	Format gl.GLenum
	Target gl.GLenum
	PixelBytesSize int
	Model EngineColorModel
}

func ColorModelToGLTypes(model color.Model) (internalFormat int, typ gl.GLenum, format gl.GLenum, target gl.GLenum, err error) {

  switch model.(type) {
  case color.Palette:
    return gl.RGBA8, gl.RGBA, gl.UNSIGNED_BYTE, gl.TEXTURE_2D, nil
  }

  switch model {
  case color.RGBAModel, color.NRGBAModel:
    return gl.RGBA8, gl.RGBA, gl.UNSIGNED_BYTE, gl.TEXTURE_2D, nil
  case color.RGBA64Model, color.NRGBAModel:
    return gl.RGBA16, gl.RGBA, gl.UNSIGNED_SHORT, gl.TEXTURE_2D, nil
  case color.AlphaModel:
    return gl.ALPHA, gl.ALPHA, gl.UNSIGNED_BYTE, gl.TEXTURE_2D, nil
  case color.Alpha16Model:
    return gl.ALPHA16, gl.ALPHA, gl.UNSIGNED_SHORT, gl.TEXTURE_2D, nil
  case color.GrayModel:
    return gl.LUMINANCE, gl.LUMINANCE, gl.UNSIGNED_BYTE, gl.TEXTURE_2D, nil
	case color.Gray16Model:
    return gl.LUMINANCE16, gl.LUMINANCE, gl.UNSIGNED_SHORT, gl.TEXTURE_2D, nil
  case color.YCbCrModel:
    return gl.RGB8, gl.RGB, gl.UNSIGNED_BYTE, gl.TEXTURE_2D, nil
  default:
    m, e := CustomColorModels[model]
    if e {
      return m.InternalFormat, m.Type, m.Format, m.Target, nil
    }
    break
  }
  return 0, 0, 0, 0, errors.New("unsupported format")
}

func ImageData(image image.Image) (data interface{}, err error) {
	w := image.Bounds().Dx()
	h := image.Bounds().Dy()
	model := image.ColorModel()

	switch model.(type) {
	case color.Palette:
		data := make([]byte, 4*h*w)
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				offset := (x + (y*w)) * 4
				r, g, b, a := image.At(x, y).RGBA()
				data[offset] = byte(r / 257)
				data[offset+1] = byte(g / 257)
				data[offset+2] = byte(b / 257)
				data[offset+3] = byte(a / 257)
			}
		}
		return data, nil
	}

	switch model {
	case color.YCbCrModel:
		data := make([]byte, 3*h*w)
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				offset := (x + (y*w)) * 3
				r, g, b, _ := image.At(x, y).RGBA()
				data[offset] = byte(r / 257)
				data[offset+1] = byte(g / 257)
				data[offset+2] = byte(b / 257)
			}
		}
		return data, nil
	case color.RGBAModel, color.NRGBAModel:
		data := make([]byte, 4*h*w)
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				offset := (x + (y*w)) * 4
				r, g, b, a := image.At(x, y).RGBA()
				data[offset] = byte(r / 257)
				data[offset+1] = byte(g / 257)
				data[offset+2] = byte(b / 257)
				data[offset+3] = byte(a / 257)
			}
		}
		return data, nil
	case color.RGBA64Model, color.NRGBA64Model:
		data := make([]byte, 4*h*w)
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				offset := (x + (y*w)) * 4
				r, g, b, a := image.At(x, y).RGBA()
				data[offset] = byte(r / 257)
				data[offset+1] = byte(g / 257)
				data[offset+2] = byte(b / 257)
				data[offset+3] = byte(a / 257)
			}
		}
		return data, nil
	default:
		m, e := CustomColorModels[model]
		if e {
			return m.Model.Data(), nil
		}
	}
	return nil, errors.New("unsupported format")
}

