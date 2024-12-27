package utils

import "image"

func IsPointInsideRect(pt image.Point, rect image.Rectangle) bool {
	return rect.Min.X <= pt.X && pt.X <= rect.Max.X &&
		rect.Min.Y <= pt.Y && pt.Y <= rect.Max.Y
}

func Abs[T int | float64](n T) T {
	if n < 0 {
		return -n
	}
	return n
}
