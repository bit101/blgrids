package blgrids

import (
	"math"
)

// LayoutCallback is a function that gets called from a layout
type LayoutCallback func(x, y float64)

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
