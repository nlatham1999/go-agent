package mouseinteractions

import (
	"fmt"

	"github.com/nlatham1999/go-agent/pkg/api"
	"github.com/nlatham1999/go-agent/pkg/model"
)

var _ api.ModelInterface = &Sim{}

type Sim struct {
	model *model.Model

	MouseClicked  bool
	MouseXClicked float64
	MouseYClicked float64

	MouseMoved bool
	MouseX     float64
	MouseY     float64
}

func NewSim() *Sim {
	return &Sim{}
}

func (s *Sim) Model() *model.Model {
	return s.model
}

func (s *Sim) Init() {
	settings := model.ModelSettings{
		WrappingX:  false,
		WrappingY:  false,
		MinPxCor:   -10, // min x patch coordinate
		MaxPxCor:   3,   // max x patch coordinate
		MinPyCor:   -10, // min y patch coordinate
		MaxPyCor:   3,   // max y patch coordinate
		RandomSeed: 10,  // random seed
		TurtleProperties: map[string]interface{}{
			"newHeading": nil,
		},
	}

	s.model = model.NewModel(settings) // create a new model with the settings
}

func (s *Sim) SetUp() error {

	s.model.ClearAll()

	s.model.CreateTurtles(1, func(t *model.Turtle) {
		t.Color = model.Red
		t.SetSize(.25)
		t.SetXY(0, 0)
	})

	return nil
}

func (s *Sim) Go() {
	if s.MouseClicked {
		fmt.Println("Mouse clicked at", s.MouseXClicked, s.MouseYClicked)
		s.MouseClicked = false
		p := s.model.Patch(s.MouseXClicked, s.MouseYClicked)
		if p == nil {
			return
		}
		p.Color.SetColor(model.White)
	}

	if s.MouseMoved {
		t := s.model.Turtle(0)
		t.FaceXY(s.MouseX, s.MouseY)
		t.Forward(0.1)
		s.MouseMoved = false
	}
}

func (s *Sim) Stats() map[string]interface{} {
	return nil
}

func (s *Sim) Stop() bool {
	return false
}

func (s *Sim) TurnPatchWhite() {
	p := s.model.Patch(-10, -10)
	p.Color.SetColor(model.White)
}

func (s *Sim) TurnPatchGreen() {
	p := s.model.Patch(-10, -10)
	p.Color.SetColor(model.Green)
}

func (s *Sim) Widgets() []api.Widget {
	return []api.Widget{
		api.NewButtonWidget("Turn Patch White", "turn-patch-white", s.TurnPatchWhite),
		api.NewButtonWidget("Turn Patch Green", "turn-patch-green", s.TurnPatchGreen),
		api.NewMouseXClickedHook(&s.MouseXClicked),
		api.NewMouseYClickedHook(&s.MouseYClicked),
		api.NewMouseClickedHook(&s.MouseClicked),
		api.NewMouseXMovedHook(&s.MouseX),
		api.NewMouseYMovedHook(&s.MouseY),
		api.NewMouseMovedHook(&s.MouseMoved),
	}
}
