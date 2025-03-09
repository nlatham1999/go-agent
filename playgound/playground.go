package playgound

import (
	"github.com/nlatham1999/go-agent/api"
	"github.com/nlatham1999/go-agent/model"
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
		WrappingX: false,
		WrappingY: false,
		MinPxCor:  -10,
		MaxPxCor:  10,
		MinPyCor:  -10,
		MaxPyCor:  10,
	}

	s.model = model.NewModel(settings)
}

func (s *Sim) SetUp() error {

	s.model.ClearAll()

	return nil
}

func (s *Sim) Go() {

	s.model.Patches.Ask(func(p *model.Patch) {
		p.PColor.SetColor(s.model.RandomColor())
	})
	// t1 := s.model.Turtle("", 2)
	// t1.Forward(10)

	s.model.Tick()

}

func (s *Sim) Stats() map[string]interface{} {
	return nil
}

func (s *Sim) Stop() bool {
	return false
}

func (s *Sim) Widgets() []api.Widget {
	return []api.Widget{}
}
