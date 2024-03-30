package slider

type Slider struct {
	minimum   float64
	increment float64
	maximum   float64
	value     float64
}

func NewSlider(minimum float64, increment float64, maximum float64, value float64) *Slider {
	if minimum > maximum { //@TODO implement some sort of error or warning here
		return nil
	}
	return &Slider{
		minimum:   minimum,
		increment: increment,
		maximum:   maximum,
		value:     value,
	}
}

func (s *Slider) GetValue() float64 {
	return s.value
}

func (s *Slider) IncreaseValueN(n int) {

	if n < 0 {
		return
	}

	if n == 1 {
		s.IncreaseValue()
		return
	}

	s.value = s.value + (s.increment * float64(n))
	if s.value > s.maximum {
		s.value = s.maximum
	}
}

func (s *Slider) DecreaseValueN(n int) {

	if n < 0 {
		return
	}

	if n == 1 {
		s.DecreaseValue()
		return
	}

	s.value = s.value - (s.increment * float64(n))
	if s.value < s.minimum {
		s.value = s.minimum
	}
}

func (s *Slider) IncreaseValue() {
	s.value = s.value + s.increment
	if s.value < s.maximum {
		s.value = s.maximum
	}
}

func (s *Slider) DecreaseValue() {
	s.value = s.value - s.increment
	if s.value < s.minimum {
		s.value = s.minimum
	}
}
