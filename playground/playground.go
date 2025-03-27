// playground for quick testing
package playground

import (
	"github.com/nlatham1999/go-agent/pkg/api"
	"github.com/nlatham1999/go-agent/pkg/model"
)

// Init()        // runs at the very beginning
// SetUp() error // sets up the model
// Go()          // runs the model

// Model() *model.Model           // returns the model
// Stats() map[string]interface{} //returns the stats of the model
// Stop() bool                    // on whether to stop the model
// Widgets() []Widget             // returns the widgets of the model

var _ api.ModelInterface = &Sim{}

type Sim struct {
	model *model.Model
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
		MaxPxCor:   10,  // max x patch coordinate
		MinPyCor:   -10, // min y patch coordinate
		MaxPyCor:   10,  // max y patch coordinate
		RandomSeed: 10,  // random seed
		TurtleProperties: map[string]interface{}{
			"newHeading": nil,
		},
	}

	s.model = model.NewModel(settings) // create a new model with the settings
}

func (s *Sim) SetUp() error {

	s.model.ClearAll()

	return nil
}

func (s *Sim) Go() {
}

func (s *Sim) Stats() map[string]interface{} {
	return nil
}

func (s *Sim) Stop() bool {
	return false
}

func (s *Sim) TurnPatchWhite() {
	p := s.model.Patch(0, 0)
	p.PColor.SetColor(model.White)
}

func (s *Sim) TurnPatchGreen() {
	p := s.model.Patch(0, 0)
	p.PColor.SetColor(model.Green)
}

func (s *Sim) Widgets() []api.Widget {
	return []api.Widget{
		api.NewButtonWidget("Turn Patch White", "turn-patch-white", s.TurnPatchWhite),
		api.NewButtonWidget("Turn Patch Green", "turn-patch-green", s.TurnPatchGreen),
	}
}
