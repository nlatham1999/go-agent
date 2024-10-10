package model

// holds RGBA values
type Color struct {
	Red   int
	Blue  int
	Green int
	Alpha int
}

// creates a new color with the same values as the given color
func (c *Color) SetColor(c2 Color) {
	c.Red = c2.Red
	c.Blue = c2.Blue
	c.Green = c2.Green
	c.Alpha = c2.Alpha
}

// creates a new color with the given red, blue, and green values
func (c *Color) SetColorRGB(red int, green int, blue int) {
	c.Red = red
	c.Blue = blue
	c.Green = green
	c.Alpha = 1
}

// creates a new color with the given red, blue, green, and alpha values
func (c *Color) SetColorRGBA(red int, green int, blue int, alpha int) {
	c.Red = red
	c.Blue = blue
	c.Green = green
	c.Alpha = alpha
}
