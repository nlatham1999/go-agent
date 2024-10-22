package playgound

import (
	"fmt"
	"sync"
	"time"

	"github.com/nlatham1999/go-agent/api"
	"github.com/nlatham1999/go-agent/model"
	"go.uber.org/atomic"
)

// Init()        // runs at the very beginning
// SetUp() error // sets up the model
// Go()          // runs the model

// Model() *model.Model           // returns the model
// Stats() map[string]interface{} //returns the stats of the model
// Stop() bool                    // on whether to stop the model
// Widgets() []Widget             // returns the widgets of the model

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
		WrappingX: true,
		WrappingY: true,
		MinPxCor:  -100,
		MaxPxCor:  100,
		MinPyCor:  -100,
		MaxPyCor:  100,
	}

	s.model = model.NewModel(settings)
}

func (s *Sim) SetUp() error {

	fmt.Println("hitting setup")

	s.model.ClearAll()

	s.model.CreateTurtles(180, "", []model.TurtleOperation{
		func(t *model.Turtle) {
			t.SetXY(-90, float64(t.Who()-90))
			t.SetHeading(0)
		},
	})

	return nil
}

func (s *Sim) Go() {

	// t1 := s.model.Turtle("", 2)
	// t1.Forward(10)

	timeNow := time.Now()

	// s.Routines()
	count := 0
	s.model.Turtles("").Ask(func(t *model.Turtle) {
		t.Forward(float64(count))
		count++
	})

	fmt.Println("Time taken: ", time.Since(timeNow))
}

func (s *Sim) Routines() {

	turtleChannel := make(chan *model.Turtle)

	go func() {
		s.model.Turtles("").Ask(func(t *model.Turtle) {
			turtleChannel <- t
		})
		close(turtleChannel)
	}()

	var wg sync.WaitGroup
	numRoutines := 45

	//count as a uber atomic counter
	count := atomic.NewInt64(0)

	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range turtleChannel {
				t.Forward(float64(count.Load()))
				count.Inc()
			}
		}()
	}

	wg.Wait()
}

func (s *Sim) Stats() map[string]interface{} {
	return nil
}

func (s *Sim) Stop() bool {
	return false
}

func (s *Sim) Widgets() []api.Widget {
	return nil
}
