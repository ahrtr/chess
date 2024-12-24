package rules

import "image"

func isPointInsideRect(pt image.Point, rect image.Rectangle) bool {
	return rect.Min.X <= pt.X && pt.X <= rect.Max.X &&
		rect.Min.Y <= pt.Y && pt.Y <= rect.Max.Y
}

func abs[T int | float64](n T) T {
	if n < 0 {
		return -n
	}
	return n
}
