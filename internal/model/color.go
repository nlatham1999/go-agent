package model

type Color struct {
	Red   int
	Blue  int
	Green int
	Alpha int //used with rbg it is the transparency
}

func (c *Color) SetColor(c2 Color) {
	c.Red = c2.Red
	c.Blue = c2.Blue
	c.Green = c2.Green
	c.Alpha = c2.Alpha
}

func (c *Color) SetColorRGB(red int, blue int, green int) {
	c.Red = red
	c.Blue = blue
	c.Green = green
}

func (c *Color) SetColorRGBA(red int, blue int, green int, alpha int) {
	c.Red = red
	c.Blue = blue
	c.Green = green
	c.Alpha = alpha
}
