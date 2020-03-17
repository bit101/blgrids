package blgrids

import (
	"math"

	"github.com/bit101/blgo"
)

// Rect draws a rectangular grid
func Rect(surface *blgo.Surface, x, y, w, h, xres, yres float64) {
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
	HexLayout(x, y, w, h, r1, func(x1, y1 float64) {
		surface.StrokePolygon(0, 0, r2, 6, math.Pi/2)
	})
}
