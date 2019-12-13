package amygdala

import (
	"image"
	"math"

	"golang.org/x/image/math/fixed"
)

func maxInt(left, right int) int {
	return int(math.Max(float64(left), float64(right)))
}

func pointToFixed(point image.Point) fixed.Point26_6 {
	return fixed.P(point.X, point.Y)
}

func rectangleFromFixed(rectangle fixed.Rectangle26_6) image.Rectangle {
	return image.Rect(
		rectangle.Min.X.Floor(),
		rectangle.Min.Y.Floor(),
		rectangle.Max.X.Ceil(),
		rectangle.Max.Y.Ceil(),
	)
}
