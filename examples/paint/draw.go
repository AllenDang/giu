package main

import (
	"image"
	"image/color"
	"math"
)

// DrawLine draws a line from (x1, y1) to (x2, y2) with the specified color and line width on the given *image.RGBA.
func DrawLine(x1, y1, x2, y2 int, c color.Color, linewidth float32, img *image.RGBA) {
	// Use Bresenham's line algorithm to get all points along the line
	dx := math.Abs(float64(x2 - x1))
	dy := math.Abs(float64(y2 - y1))

	sx := -1
	if x1 < x2 {
		sx = 1
	}

	sy := -1
	if y1 < y2 {
		sy = 1
	}

	err := dx - dy

	for {
		// Draw a circle at each point to simulate line width
		drawCircle(img, x1, y1, linewidth/2, c)

		if x1 == x2 && y1 == y2 {
			break
		}

		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x1 += sx
		}

		if e2 < dx {
			err += dx
			y1 += sy
		}
	}
}

// drawCircle draws a filled circle with the specified radius and color at (cx, cy) on the *image.RGBA.
func drawCircle(img *image.RGBA, cx, cy int, radius float32, c color.Color) {
	r := int(radius)
	for x := -r; x <= r; x++ {
		for y := -r; y <= r; y++ {
			if x*x+y*y <= r*r {
				img.Set(cx+x, cy+y, c)
			}
		}
	}
}

// Floodfill fills an area of an image with a given color starting at point (x, y).
// The fill continues for all adjacent pixels of the same starting color.
func Floodfill(input *image.RGBA, c color.Color, x, y int) {
	// Get the color of the starting pixel
	startColor := input.At(x, y)

	// If the starting pixel is already the target color, return
	if colorsEqual(startColor, c) {
		return
	}

	// A queue to process pixels
	queue := []image.Point{{X: x, Y: y}}
	bounds := input.Bounds()

	// Process the queue iteratively
	for len(queue) > 0 {
		// Dequeue a pixel
		point := queue[0]
		queue = queue[1:]

		px, py := point.X, point.Y

		// Ignore if out of bounds
		if px < bounds.Min.X || px >= bounds.Max.X || py < bounds.Min.Y || py >= bounds.Max.Y {
			continue
		}

		// Check the color of the current pixel
		if !colorsEqual(input.At(px, py), startColor) {
			continue
		}

		// Set the new color
		input.Set(px, py, c)

		// Add the neighboring pixels to the queue
		queue = append(queue, image.Point{X: px + 1, Y: py}, image.Point{X: px - 1, Y: py}, image.Point{X: px, Y: py + 1}, image.Point{X: px, Y: py - 1})
	}
}

// colorsEqual compares two colors and returns true if they are equal.
func colorsEqual(c1, c2 color.Color) bool {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()

	return r1 == r2 && g1 == g2 && b1 == b2 && a1 == a2
}
