package blgrids

import (
	"fmt"
	"math"

	"github.com/bit101/blgo"
)

// Rect draws a rectangular grid
func Rect(surface *blgo.Surface, x, y, w, h, xres, yres float64) {
	fmt.Printf("yres = %+v\n", yres)
	for y1 := y; y1 < y+h; y1 += yres {
		surface.MoveTo(x, y1)
		surface.LineTo(x+w, y1)
	}
	for x1 := x; x1 < x+w; x1 += xres {
		surface.MoveTo(x1, y)
		surface.LineTo(x1, y+h)
	}
	surface.Stroke()
}

// Square draws a square grid
func Square(surface *blgo.Surface, x, y, w, h, res float64) {
	Rect(surface, x, y, w, h, res, res)
}

// Hex draws a hex grid
func Hex(surface *blgo.Surface, x, y, w, h, r1, r2 float64) {
	sin60 := math.Sin(math.Pi / 3)
	xinc := r1 * 2 * sin60
	yinc := r1 * 1.5
	offset := 0.0

	for y1 := y; y1 < y+h*2+r1; y1 += yinc {
		for x1 := x; x1 < x+w+r1; x1 += xinc {
			surface.Save()
			surface.Translate(x1+offset, y1)
			surface.Rotate(math.Pi / 2)
			surface.StrokePolygon(0, 0, r2, 6, 0)
			surface.Restore()
		}
		if offset == 0 {
			offset = r1 * sin60
		} else {
			offset = 0
		}
	}
}

// HexWithRenderer draws a hex grid with a render function
func HexWithRenderer(surface *blgo.Surface, x, y, w, h, r float64, render HexagonRenderer) {
	sin60 := math.Sin(math.Pi / 3)
	xinc := r * 2 * sin60
	yinc := r * 1.5
	offset := 0.0

	for y1 := y; y1 < y+h*2+r; y1 += yinc {
		for x1 := x; x1 < x+w+r; x1 += xinc {
			surface.Save()
			surface.Translate(x1+offset, y1)
			surface.Rotate(math.Pi / 2)
			render(surface, r)
			surface.Restore()
		}
		if offset == 0 {
			offset = r * sin60
		} else {
			offset = 0
		}
	}
}

// HexagonRenderer is the interface for a function which renders a hexagon
type HexagonRenderer func(surface *blgo.Surface, r float64)
