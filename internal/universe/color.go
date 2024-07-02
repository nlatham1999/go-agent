package universe

type Color struct {
	usingScaleColor bool
	scaleColor      float64
	rbg             RGB
	alpha           int //used with rbg it is the transparency
}

type RGB struct {
	red   int
	blue  int
	green int
}

type HSB struct {
	hue        int
	saturation int
	brightness int
}

func (c *Color) SetColorScale(val float64) {
	c.scaleColor = val
	c.usingScaleColor = true
}

func (c *Color) SetColorRGB(red int, blue int, green int) {
	c.rbg = RGB{
		red:   red,
		blue:  blue,
		green: green,
	}
	c.alpha = 0
	c.usingScaleColor = false
}

func (c *Color) SetColorRGBA(red int, blue int, green int, alpha int) {
	c.rbg = RGB{
		red:   red,
		green: green,
		blue:  blue,
	}
	c.alpha = alpha
	c.usingScaleColor = false
}

func (c *Color) GetColorScale() float64 {
	if c.usingScaleColor {
		return c.scaleColor
	} else {
		return ApproximateRGB(c.rbg.red, c.rbg.green, c.rbg.blue)
	}
}
