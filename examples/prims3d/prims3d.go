package prims3d

import (
	"fmt"
	"time"

	"github.com/nlatham1999/go-agent/pkg/api"
	"github.com/nlatham1999/go-agent/pkg/model"
)

type Prims struct {
	model *model.Model

	placedTurtleBreed   *model.TurtleBreed
	unplacedTurtleBreed *model.TurtleBreed
	placedLinkBreed     *model.LinkBreed
	unplacedLinkBreed   *model.LinkBreed

	liveColor model.Color

	nodeSize float64

	nodes int

	sortedLinks []*model.Link // cached sorted list of unplaced links
}

func NewPrims() *Prims {
	return &Prims{}
}

func (p *Prims) Init() {

	p.placedTurtleBreed = model.NewTurtleBreed("placed", "", nil)
	p.unplacedTurtleBreed = model.NewTurtleBreed("unplaced", "", nil)

	p.placedLinkBreed = model.NewLinkBreed("unplaced")
	p.unplacedLinkBreed = model.NewLinkBreed("placed")

	modelSettings := model.ModelSettings{
		TurtleBreeds:         []*model.TurtleBreed{p.placedTurtleBreed, p.unplacedTurtleBreed},
		UndirectedLinkBreeds: []*model.LinkBreed{p.placedLinkBreed, p.unplacedLinkBreed},
		MinPxCor:             0,
		MinPyCor:             0,
		MaxPxCor:             10,
		MaxPyCor:             10,
		MinPzCor:             0,
		MaxPzCor:             10,
	}

	// p.liveColor = model.White
	// neon green 57, 255, 20
	p.liveColor = model.Color{Red: 57, Green: 255, Blue: 20, Alpha: 255}
	p.nodeSize = 0.2

	p.model = model.NewModel(modelSettings)

	p.nodes = 1500
}

func (p *Prims) SetUp() error {

	p.model.ClearAll()

	unplaced := p.model.TurtleBreed("unplaced")
	unplaced.CreateAgents(p.nodes,
		p.placeInitialNodes,
	)

	//for each turtle create a link with every other turtle
	unplaced.Agents().Ask(
		p.createInitialLinks,
	)

	p.unplacedLinkBreed.Links().SortAsc(func(l *model.Link) float64 {
		return l.Length()
	})

	// Cache sorted links as a list for faster iteration
	p.sortedLinks = p.unplacedLinkBreed.Links().List()

	fmt.Println("Initial links: ", len(p.sortedLinks))

	t0 := p.model.Turtle(0)
	t0.Color.SetColor(p.liveColor)

	placed := p.model.TurtleBreed("placed")

	t0.SetBreed(placed)

	p.model.ResetTicks()
	return nil
}

func (p *Prims) placeInitialNodes(t *model.Turtle) {
	t.Color.SetColor(model.Gray)
	t.SetSize(p.nodeSize)
	t.SetXYZ(p.model.RandomXCor(), p.model.RandomYCor(), p.model.RandomZCor())
}

// returns the approximate average distance between nodes based on current node count and area
func (p *Prims) densityFunc() float64 {
	volume := p.volumeFunc()
	avgVolumePerNode := volume / float64(p.nodes)
	return avgVolumePerNode + volume*0.0003
}

func (p *Prims) volumeFunc() float64 {
	x := p.model.MaxPxCor() - p.model.MinPxCor() + 1
	y := p.model.MaxPyCor() - p.model.MinPyCor() + 1
	z := p.model.MaxPzCor() - p.model.MinPzCor() + 1

	volume := float64(x * y * z)
	return volume
}

func (p *Prims) createInitialLinks(t *model.Turtle) {

	distance := p.densityFunc()
	fmt.Println("Density Value", distance)

	// first check to see if there's any other turtles in the radius that are unplaced
	turtlesToLinkTo := p.model.TurtlesInRadiusXYZ(t.XCor(), t.YCor(), t.ZCor(), distance+.1).With(func(t2 *model.Turtle) bool {
		return t.BreedName() == "unplaced" && t != t2
	})

	// if none, link to all unplaced turtles
	if turtlesToLinkTo.Count() == 0 {
		fmt.Println("No nearby turtles, linking to all unplaced")
		turtlesToLinkTo = p.model.TurtleBreed("unplaced").Agents()
	}

	turtlesToLinkTo.Ask(
		func(t2 *model.Turtle) {
			if t != t2 { // && t.DistanceTurtle(t2) < 10 {
				t.CreateLinkWithTurtle(p.unplacedLinkBreed, t2,
					func(l *model.Link) {
						l.Color.SetColor(model.Gray)
						l.Hide()
					},
				)
			}
		},
	)
}

func (p *Prims) Go() {
	start := time.Now()

	// find the closest link to the cluster by iterating through cached sorted list
	var closestLink *model.Link
	var closestTurtle *model.Turtle
	newIndex := 0

	// Iterate through sorted links, removing invalid ones as we go
	for i, l := range p.sortedLinks {
		breedName1 := l.End1().BreedName()
		breedName2 := l.End2().BreedName()

		// Both placed - remove this link
		if breedName1 == "placed" && breedName2 == "placed" {
			l.Die()
			continue // Skip this link, don't add to newIndex
		}

		// Both unplaced - keep in list but skip for now
		if breedName1 == "unplaced" && breedName2 == "unplaced" {
			p.sortedLinks[newIndex] = l
			newIndex++
			continue
		}

		// Found the first valid link (one placed, one unplaced)
		closestLink = l
		if breedName1 == "unplaced" {
			closestTurtle = l.End1()
		} else {
			closestTurtle = l.End2()
		}

		// Copy remaining links and break
		copy(p.sortedLinks[newIndex:], p.sortedLinks[i+1:])
		newIndex += len(p.sortedLinks) - i - 1
		break
	}

	// Trim the slice to remove processed items
	p.sortedLinks = p.sortedLinks[:newIndex]

	if closestLink == nil {
		return
	}

	placed := p.model.TurtleBreed("placed")

	//add the link and turtle to the cluster
	closestLink.SetBreed(p.placedLinkBreed)
	closestLink.Color.SetColor(p.liveColor)
	closestLink.Show()
	closestTurtle.SetBreed(placed)
	closestTurtle.Color.SetColor(p.liveColor)

	// if all nodes are placed, kill all remaining unplaced links
	if placed.Agents().Count() == p.nodes {
		for _, l := range p.sortedLinks {
			l.Die()
		}
		p.sortedLinks = nil
	}

	p.model.Tick()

	// if there are no unplaced links left, run setup to reset everything
	if len(p.sortedLinks) == 0 {
		p.SetUp()
	}
	fmt.Println("Time taken: ", time.Since(start).Seconds())

}

func (p *Prims) Model() *model.Model {
	return p.model
}

func (p *Prims) Stats() map[string]interface{} {

	placed := p.model.TurtleBreed("placed")
	unplaced := p.model.TurtleBreed("unplaced")

	return map[string]interface{}{
		"Placed nodes":    placed.Agents().Count(),
		"Unplaced nodes":  unplaced.Agents().Count(),
		"potential links": len(p.sortedLinks),
	}
}

func (p *Prims) Stop() bool {
	// return p.placedLinkBreed.Links().Count() >= p.nodes-2
	return false
}

func (p *Prims) Widgets() []api.Widget {
	return []api.Widget{
		api.NewIntSliderWidget("Nodes", "nodes", "0", "4000", "1500", "100", &p.nodes),
		api.NewFloatSliderWidget("node size", "node-size", ".1", "1.0", ".2", ".1", &p.nodeSize),
	}
}
