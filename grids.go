package blgrids

import (
	"math"

	"github.com/bit101/blgo"
)

// RectGrid draws a rectangular grid
func RectGrid(surface *blgo.Surface, x, y, w, h, xres, yres float64) {
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

// SquareGrid draws a square grid
func SquareGrid(surface *blgo.Surface, x, y, w, h, res float64) {
	RectGrid(surface, x, y, w, h, res, res)
}

// HexGrid draws a hex grid
func HexGrid(surface *blgo.Surface, x, y, w, h, r1, r2 float64) {
	HexLayout(x, y, w, h, r1, func(x1, y1 float64) {
		surface.StrokePolygon(0, 0, r2, 6, math.Pi/2)
	})
}
