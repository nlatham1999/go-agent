package simplesim

import (
	"fmt"

	"github.com/nlatham1999/go-agent/internal/model"
)

type SimpleSim struct {
	m      *model.Model
	repeat bool
}

func NewSimpleSim() *SimpleSim {
	return &SimpleSim{
		repeat: false,
	}
}

func (s *SimpleSim) Model() *model.Model {
	return s.m
}

func (s *SimpleSim) Repeat() {
	s.repeat = !s.repeat
}

func (s *SimpleSim) Init() {

	fmt.Println("Initializing model")

	settings := model.ModelSettings{
		WrappingX: true,
		WrappingY: true,
	}

	s.m = model.NewModel(settings)

	if s.m != nil {
		fmt.Println("Model initialized")
	}
}

func (s *SimpleSim) SetUp() {

	s.m.ClearAll()

	fmt.Println(s.m.Turtles("").Count())

	s.m.CreateTurtles(1, "", nil)

	t1 := s.m.Turtle("", 0)
	// t2 := s.m.Turtle("", 1)
	// t3 := s.m.Turtle("", 2)
	// t4 := s.m.Turtle("", 3)

	t1.SetXY(0, 0)
	t1.SetHeading(0)

	// t2.SetXY(2, 2)
	// t3.SetXY(2, 0)
	// t4.SetXY(3, 1)
}

func (s *SimpleSim) Go() {
	t1 := s.m.Turtle("", 0)
	t1.Forward(.2)
}
