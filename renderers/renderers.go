package renderers

import (
	"errors"
	"fmt"
	"github.com/Team-TCG/canvas"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"path/filepath"
	"strings"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

func Write(filename string, c *canvas.Canvas, opts ...interface{}) error {
	switch ext := strings.ToLower(filepath.Ext(filename)); ext {
	case ".png":
		return c.WriteFile(filename, PNG(opts...))
	case ".jpg", ".jpeg":
		return c.WriteFile(filename, JPEG(opts...))
	case ".gif":
		return c.WriteFile(filename, GIF(opts...))
	case ".tif", ".tiff":
		return c.WriteFile(filename, TIFF(opts...))
	case ".bmp":
		return c.WriteFile(filename, BMP(opts...))
	default:
		return errors.New(fmt.Sprintf("unknown file extension: %v", ext))
	}
}

func errorWriter(err error) canvas.Writer {
	return func(w io.Writer, c *canvas.Canvas) error {
		return err
	}
}

func PNG(opts ...interface{}) canvas.Writer {
	resolution := canvas.DPMM(1.0)
	colorSpace := canvas.DefaultColorSpace
	for _, opt := range opts {
		switch o := opt.(type) {
		case canvas.Resolution:
			resolution = o
		case canvas.ColorSpace:
			colorSpace = o
		default:
			return errorWriter(fmt.Errorf("unknown PNG option: %T(%v)", opt, opt))
		}
	}
	return func(w io.Writer, c *canvas.Canvas) error {
		img := Draw(c, resolution, colorSpace)
		return png.Encode(w, img)
	}
}

func JPEG(opts ...interface{}) canvas.Writer {
	resolution := canvas.DPMM(1.0)
	colorSpace := canvas.DefaultColorSpace
	var options *jpeg.Options
	for _, opt := range opts {
		switch o := opt.(type) {
		case canvas.Resolution:
			resolution = o
		case canvas.ColorSpace:
			colorSpace = o
		case *jpeg.Options:
			options = o
		default:
			return errorWriter(fmt.Errorf("unknown JPEG option: %T(%v)", opt, opt))
		}
	}
	return func(w io.Writer, c *canvas.Canvas) error {
		img := Draw(c, resolution, colorSpace)
		return jpeg.Encode(w, img, options)
	}
}

func GIF(opts ...interface{}) canvas.Writer {
	resolution := canvas.DPMM(1.0)
	colorSpace := canvas.DefaultColorSpace
	var options *gif.Options
	for _, opt := range opts {
		switch o := opt.(type) {
		case canvas.Resolution:
			resolution = o
		case canvas.ColorSpace:
			colorSpace = o
		case *gif.Options:
			options = o
		default:
			return errorWriter(fmt.Errorf("unknown option: %T(%v)", opt, opt))
		}
	}
	return func(w io.Writer, c *canvas.Canvas) error {
		img := Draw(c, resolution, colorSpace)
		return gif.Encode(w, img, options)
	}
}

func TIFF(opts ...interface{}) canvas.Writer {
	resolution := canvas.DPMM(1.0)
	colorSpace := canvas.DefaultColorSpace
	var options *tiff.Options
	for _, opt := range opts {
		switch o := opt.(type) {
		case canvas.Resolution:
			resolution = o
		case canvas.ColorSpace:
			colorSpace = o
		case *tiff.Options:
			options = o
		default:
			return errorWriter(fmt.Errorf("unknown option: %T(%v)", opt, opt))
		}
	}
	return func(w io.Writer, c *canvas.Canvas) error {
		img := Draw(c, resolution, colorSpace)
		return tiff.Encode(w, img, options)
	}
}

func BMP(opts ...interface{}) canvas.Writer {
	resolution := canvas.DPMM(1.0)
	colorSpace := canvas.DefaultColorSpace
	for _, opt := range opts {
		switch o := opt.(type) {
		case canvas.Resolution:
			resolution = o
		case canvas.ColorSpace:
			colorSpace = o
		default:
			return errorWriter(fmt.Errorf("unknown option: %T(%v)", opt, opt))
		}
	}
	return func(w io.Writer, c *canvas.Canvas) error {
		img := Draw(c, resolution, colorSpace)
		return bmp.Encode(w, img)
	}
}
