package blgrids

import (
	"math"
)

// LayoutCallback is a function that gets called from a layout
type LayoutCallback func(x, y float64)

// RectLayout draws a rectangular grid
func RectLayout(x, y, w, h, xres, yres float64, callback LayoutCallback) {
	for y1 := y; y1 < y+h; y1 += yres {
		for x1 := x; x1 < x+w; x1 += xres {
			callback(x1+xres/2, y1+yres/2)
		}
	}
}

// SquareLayout draws a square grid
func SquareLayout(x, y, w, h, res float64, callback LayoutCallback) {
	RectLayout(x, y, w, h, res, res, callback)
}

// HexLayout calls a callback at every location in a hex grid
func HexLayout(x, y, w, h, r float64, callback LayoutCallback) {
	sin60r := math.Sin(math.Pi/3) * r
	xinc := 2 * sin60r
	yinc := r * 1.5
	offset := 0.0

	for y1 := y; y1 < y+h*2+r; y1 += yinc {
		for x1 := x; x1 < x+w+r; x1 += xinc {
			callback(x1+offset, y1)
		}
		if offset == 0 {
			offset = sin60r
		} else {
			offset = 0
		}
	}
}
