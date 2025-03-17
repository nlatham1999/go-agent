package schelling

import (
	"fmt"
	"time"

	"github.com/nlatham1999/go-agent/pkg/api"
	"github.com/nlatham1999/go-agent/pkg/model"
)

var _ api.ModelInterface = (*Schelling)(nil)

type Schelling struct {
	model *model.Model

	redBluePercentage    float64 // percentage of red and blue agents - .3 means 30% red, 70% blue
	emptyPercentage      float64 // percentage of empty patches
	segregationThreshold float64

	populatedPatches   *model.PatchAgentSet
	unpopulatedPatches *model.PatchAgentSet

	totalUnhappy int
	totalHappy   int
}

func NewSchelling() *Schelling {
	return &Schelling{}
}

func (s *Schelling) Init() {

	modelSettings := model.ModelSettings{
		MinPxCor: 0,
		MinPyCor: 0,
		MaxPxCor: 49,
		MaxPyCor: 49,
		PatchProperties: map[string]interface{}{
			"group": "none",
		},
	}

	s.model = model.NewModel(modelSettings)

	s.redBluePercentage = .5
	s.emptyPercentage = .3
	s.segregationThreshold = .40
}

func (s *Schelling) SetUp() error {

	s.model.ClearAll()
	s.populatedPatches = model.NewPatchAgentSet(nil)
	s.unpopulatedPatches = model.NewPatchAgentSet(nil)
	s.totalHappy = -1
	s.totalUnhappy = -1

	redBluePercentageAdjusted := 1 - s.emptyPercentage
	redPercentage := redBluePercentageAdjusted * s.redBluePercentage
	bluePercentage := redBluePercentageAdjusted * (1 - s.redBluePercentage)

	// create patches and assign to red, blue, or empty
	s.model.Patches.Ask(
		func(p *model.Patch) {
			randomInt := s.model.RandomFloat(1)
			if randomInt < redPercentage {
				p.SetProperty("group", "red")
			} else if randomInt < redPercentage+bluePercentage {
				p.SetProperty("group", "blue")
			}
		},
	)

	s.model.Patches.Ask(
		func(p *model.Patch) {
			if p.GetProperty("group") == "none" {
				p.PColor = model.Black
				s.unpopulatedPatches.Add(p)
			} else if p.GetProperty("group") == "red" {
				p.PColor = model.Red
				s.populatedPatches.Add(p)
			} else if p.GetProperty("group") == "blue" {
				p.PColor = model.Blue
				s.populatedPatches.Add(p)
			}
		},
	)

	return nil
}

func (s *Schelling) Go() {

	timeNow := time.Now()

	happy := 0
	unhappy := 0

	// go through each patch and check if it is happy
	s.populatedPatches.Ask(
		func(p *model.Patch) {
			group := p.GetProperty("group")

			if group == "none" {
				return
			}

			percentage := s.getPercentage(p)

			// if less than threshold, move to a random empty
			if percentage < s.segregationThreshold {
				// look through the unpoulated patches for a new home where the number of
				newHome := s.unpopulatedPatches.RandomWhere(func(p2 *model.Patch) bool {
					return p2.GetProperty("group") == "none" && s.getPercentage(p2) >= s.segregationThreshold
				})

				if newHome != nil {
					p.SetProperty("group", "none")
					newHome.SetProperty("group", group)
					s.populatedPatches.Add(newHome)
					s.unpopulatedPatches.Remove(newHome)
					s.populatedPatches.Remove(p)
					s.unpopulatedPatches.Add(p)
				}
				unhappy++

			} else {
				happy++
			}

		},
	)

	// recolor the patches
	s.model.Patches.Ask(
		func(p *model.Patch) {
			if p.GetProperty("group") == "none" {
				p.PColor = model.Black
			} else if p.GetProperty("group") == "red" {
				p.PColor = model.Red
			} else if p.GetProperty("group") == "blue" {
				p.PColor = model.Blue
			}
		},
	)

	fmt.Println("Time taken: ", time.Since(timeNow))

	s.totalHappy = happy
	s.totalUnhappy = unhappy

	s.model.Tick()

}

// get the percentage of neighbors that are the same group compared to all neighbors, excluding empty patches
func (s *Schelling) getPercentage(p *model.Patch) float64 {
	neighbors := p.Neighbors()

	sameNeighbors := 0
	otherNeighbors := 0
	neighbors.Ask(
		func(p2 *model.Patch) {
			if p2.GetProperty("group") == p.GetProperty("group") {
				sameNeighbors++
			} else if p2.GetProperty("group") != "none" {
				otherNeighbors++
			}
		},
	)

	return float64(sameNeighbors) / float64(sameNeighbors+otherNeighbors)
}

func (s *Schelling) Model() *model.Model {
	return s.model
}

func (s *Schelling) Stats() map[string]interface{} {

	var occupied int
	var unoccupied int

	if s.populatedPatches != nil {
		occupied = s.populatedPatches.Count()
	} else {
		occupied = 0
	}

	if s.unpopulatedPatches != nil {
		unoccupied = s.unpopulatedPatches.Count()
	} else {
		unoccupied = 0
	}

	return map[string]interface{}{
		"occupied":   occupied,
		"unoccupied": unoccupied,
		"happy":      s.totalHappy,
		"unhappy":    s.totalUnhappy,
	}
}

func (s *Schelling) Stop() bool {
	return s.totalUnhappy == 0
}

func (s *Schelling) Widgets() []api.Widget {
	return []api.Widget{
		api.NewFloatSliderWidget("Segregation Threshold", "segregationThreshold", "0", "1", ".40", ".01", &s.segregationThreshold),
		api.NewFloatSliderWidget("Red/Blue Ratio", "redBluePercentage", "0", "1", ".5", ".01", &s.redBluePercentage),
		api.NewFloatSliderWidget("Empty Percentage", "emptyPercentage", "0", "1", ".3", ".01", &s.emptyPercentage),
	}
}
