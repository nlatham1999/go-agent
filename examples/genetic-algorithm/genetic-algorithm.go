package geneticalgorithm

import (
	"context"
	"strconv"
	"sync"

	"github.com/nlatham1999/go-agent/pkg/api"
	"github.com/nlatham1999/go-agent/pkg/model"
)

// concurrently runs multiple models every step
// demonstrates a basic genetic algorithm
// taken from Melanie Mitchell's trash picking robot example from Complexity: A Guided Tour

var _ api.ModelInterface = &GA{}

var possibleActions []string = []string{"MOVENORTH", "MOVESOUTH", "MOVEWEST", "MOVEEAST", "PICKUP"}

type GA struct {
	models []*model.Model

	rules map[string]int // map of states to rule numbers

	numRobots int

	highScore           int
	generation          int
	averageScore        int
	highestAmountPicked int

	generationLength    int
	mutationProbability int
}

func NewGeneticAlgorithm() *GA {
	return &GA{}
}

func (ga *GA) Model() *model.Model {

	if ga.models == nil {
		return nil
	}

	if len(ga.models) == 0 {
		return nil
	}

	return ga.models[0]
}

func (ga *GA) Init() {

	ga.loadRules()

	ga.numRobots = 1500

	ga.generationLength = 200

	ga.mutationProbability = 1

	_ = ga.SetUp()
}

func (ga *GA) SetUp() error {

	for i := 0; i < ga.numRobots; i++ {
		cans := model.NewTurtleBreed("cans", "circle", map[string]interface{}{
			"picked": false,
		})

		robots := model.NewTurtleBreed("robots", "circle", map[string]interface{}{
			"score":         0,
			"amount-picked": 0,
			"dna":           nil,
		})

		modelSettings := model.ModelSettings{
			TurtleBreeds: []*model.TurtleBreed{robots, cans},
			WrappingX:    false,
			WrappingY:    false,
			MinPxCor:     0,
			MaxPxCor:     9,
			MinPyCor:     0,
			MaxPyCor:     9,
			RandomSeed:   uint64(i),
			RandomSeed2:  uint64(i),
		}

		m := model.NewModel(modelSettings)

		m.Patches.Ask(func(p *model.Patch) {
			if p.XCor() == 0 && p.YCor() == 0 {
				return
			}
			if m.RandomInt(100) < 50 {
				cans.CreateAgents(1, func(t *model.Turtle) {
					t.Color = model.Red
					t.SetSize(.25)
					t.SetXY(float64(p.XCor()), float64(p.YCor()))
					t.SetProperty("picked", false)
				})
			}
		})

		robots.CreateAgents(1, func(t *model.Turtle) {
			t.Color = model.Blue
			t.SetSize(.5)
			t.SetXY(0, 0)
			t.SetProperty("score", 0)
			t.SetProperty("amount-picked", 0)
			t.SetProperty("dna", []string{})
			t.SetProperty("picked", map[*model.Turtle]interface{}{})
			ga.LoadDNAIntoTurtle(m, t)
		})

		ga.models = append(ga.models, m)
	}

	ga.highScore = 0
	ga.generation = 0
	ga.highestAmountPicked = 0
	ga.averageScore = 0
	ga.highScore = 0

	return nil
}

func (ga *GA) Go() {

	// context with cancellation
	ctx := context.Background()
	ctx, _ = context.WithCancel(ctx)

	waitgroup := sync.WaitGroup{}
	concurrency := 10

	modelChannel := make(chan *model.Model, len(ga.models))
	for _, m := range ga.models {
		modelChannel <- m
	}
	close(modelChannel)

	for i := 0; i < concurrency; i++ {
		waitgroup.Add(1)
		go func() {
			defer waitgroup.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case m, ok := <-modelChannel:
					if !ok {
						return
					}

					robots := m.TurtleBreed("robots")
					cans := m.TurtleBreed("cans")

					for m.Ticks < ga.generationLength {
						robots.Agents().Ask(func(t *model.Turtle) {
							ga.doAction(t)
							t.SetLabel(strconv.Itoa(t.GetProperty("score").(int)))
						})

						// recolor cans based on how many times they have been picked up
						cans.Agents().Ask(func(t *model.Turtle) {
							picked := t.GetProperty("picked").(bool)
							if picked {
								t.Color = model.Black
							} else {
								t.Color = model.Red
							}
						})

						m.Tick()
					}
				}
			}

		}()
	}

	// wait for all goroutines to finish
	waitgroup.Wait()

	ga.averageScore = 0
	ga.highScore = -100000
	ga.highestAmountPicked = 0
	for _, m := range ga.models {
		robot, _ := m.TurtleBreed("robots").Agents().First()
		if robot != nil {
			ga.averageScore += robot.GetProperty("score").(int)
			if robot.GetProperty("score").(int) > ga.highScore {
				ga.highScore = robot.GetProperty("score").(int)
			}
		} else {
		}

		cansPicked := m.TurtleBreed("cans").Agents().With(func(t *model.Turtle) bool {
			return t.GetProperty("picked").(bool)
		}).Count()

		if cansPicked > ga.highestAmountPicked {
			ga.highestAmountPicked = cansPicked
		}
	}

	ga.averageScore = ga.averageScore / ga.numRobots
	ga.generation++
	ga.newGeneration()

}

func (ga *GA) Stats() map[string]interface{} {
	stats := map[string]interface{}{
		"high score":            ga.highScore,
		"average score":         ga.averageScore,
		"generation":            ga.generation,
		"highest amount picked": ga.highestAmountPicked,
	}

	return stats
}

func (ga *GA) Stop() bool {
	return false
}

func (ga *GA) Widgets() []api.Widget {
	return []api.Widget{
		api.NewButtonWidget("New Generation", "new-gen", ga.newGeneration),
		api.NewIntSliderWidget("Generation Length", "generation-length", "1", "1000", strconv.Itoa(ga.generationLength), "1", &ga.generationLength),
		api.NewIntSliderWidget("Number of Robots", "num-robots", "1", "1000", strconv.Itoa(ga.numRobots), "1", &ga.numRobots),
		api.NewIntSliderWidget("Mutation Probability", "mutation-probability", "0", "100", strconv.Itoa(ga.mutationProbability), "1", &ga.mutationProbability),
	}
}

// load all possible combinations for North, South, East, West and current
func (ga *GA) loadRules() {

	ga.rules = map[string]int{}

	possibleStates := []string{"empty", "can", "wall"}

	count := 0
	for i := 0; i < len(possibleStates); i++ { //North
		for j := 0; j < len(possibleStates); j++ { //South
			for k := 0; k < len(possibleStates); k++ { //East
				for l := 0; l < len(possibleStates); l++ { //West
					for m := 0; m < len(possibleStates); m++ { //current
						state := possibleStates[i] + possibleStates[j] + possibleStates[k] + possibleStates[l] + possibleStates[m]
						ga.rules[state] = count
						count++
					}
				}
			}
		}
	}
}

func (ga *GA) LoadDNAIntoTurtle(m *model.Model, t *model.Turtle) {
	dna := []string{}

	for i := 0; i < 243; i++ {
		index := m.RandomInt(len(possibleActions))
		action := possibleActions[index]
		dna = append(dna, action)
	}

	t.SetProperty("dna", dna)
}

func (ga *GA) doAction(t *model.Turtle) {

	m := t.Model()

	cans := m.TurtleBreed("cans")

	// get the states of the surrounding patches
	northPatchState := ga.getPatchState(t, m.Patch(t.XCor(), t.YCor()+1))
	southPatchState := ga.getPatchState(t, m.Patch(t.XCor(), t.YCor()-1))
	eastPatchState := ga.getPatchState(t, m.Patch(t.XCor()+1, t.YCor()))
	westPatchState := ga.getPatchState(t, m.Patch(t.XCor()-1, t.YCor()))
	currentPatchState := ga.getPatchState(t, m.Patch(t.XCor(), t.YCor()))

	state := northPatchState + southPatchState + eastPatchState + westPatchState + currentPatchState
	actionIndex := ga.rules[state]
	action := t.GetProperty("dna").([]string)[actionIndex]

	if action == "MOVERANDOM" {
		// pick a random move action
		index := m.RandomInt(4)
		switch index {
		case 0:
			action = "MOVENORTH"
		case 1:
			action = "MOVESOUTH"
		case 2:
			action = "MOVEWEST"
		case 3:
			action = "MOVEEAST"
		}
	}

	switch action {
	case "MOVENORTH":
		if m.Patch(t.XCor(), t.YCor()+1) != nil {
			t.SetXY(t.XCor(), t.YCor()+1)
		} else {
			t.SetProperty("score", t.GetProperty("score").(int)-5) //penalize for hitting a wall
		}
	case "MOVESOUTH":
		if m.Patch(t.XCor(), t.YCor()-1) != nil {
			t.SetXY(t.XCor(), t.YCor()-1)
		} else {
			t.SetProperty("score", t.GetProperty("score").(int)-5) //penalize for hitting a wall
		}
	case "MOVEWEST":
		if m.Patch(t.XCor()-1, t.YCor()) != nil {
			t.SetXY(t.XCor()-1, t.YCor())
		} else {
			t.SetProperty("score", t.GetProperty("score").(int)-5) //penalize for hitting a wall
		}
	case "MOVEEAST":
		if m.Patch(t.XCor()+1, t.YCor()) != nil {
			t.SetXY(t.XCor()+1, t.YCor())
		} else {
			t.SetProperty("score", t.GetProperty("score").(int)-5) //penalize for hitting a wall
		}
	case "PICKUP":
		// check if the patch has a can on it
		cans := cans.AgentsOnPatch(m.Patch(t.XCor(), t.YCor()))
		if cans.Count() > 0 {
			can, _ := cans.First()
			picked := can.GetProperty("picked").(bool)
			// check if the can has already been picked up
			if picked {
				t.SetProperty("score", t.GetProperty("score").(int)-1) //penalize for picking up a can that has already been picked up
			} else {
				t.SetProperty("score", t.GetProperty("score").(int)+10) //increment score
				t.SetProperty("amount-picked", t.GetProperty("amount-picked").(int)+1)
				can.SetProperty("picked", true) // set the can to be picked up
			}
		} else {
			t.SetProperty("score", t.GetProperty("score").(int)-1) //penalize for choosing to pick up a can when there isn't one
		}
	case "STAY":
		// do nothing
	default:
		// do nothing
	}
}

func (ga *GA) getPatchState(t *model.Turtle, p *model.Patch) string {

	m := t.Model()

	cans := m.TurtleBreed("cans")

	state := ""
	if p == nil {
		state = "wall"
	} else {
		cans := cans.AgentsOnPatch(p)
		if cans.Count() > 0 {
			can, _ := cans.First()
			picked := can.GetProperty("picked").(bool)
			// check if the can has already been picked up
			if picked {
				state = "empty" // can has been picked up
			} else {
				state = "can"
			}
		} else {
			state = "empty"
		}
	}
	return state
}

func (ga *GA) newGeneration() {

	// aggregate all the robots from all the models
	allRobots := model.NewTurtleAgentSet([]*model.Turtle{})
	for _, m := range ga.models {
		robot, _ := m.TurtleBreed("robots").Agents().First()
		if robot != nil {
			allRobots.Add(robot)
		} else {
		}
	}

	//sort the robots by score
	allRobots.SortDesc(func(t *model.Turtle) float64 {
		return float64(t.GetProperty("score").(int))
	})

	topAmount := allRobots.Count() / 18

	// get the top part of the robots
	topHalfSet := allRobots.FirstNOf(topAmount)
	// create a weighted list where robots with higher scores are more likely to be selected
	topHalfUnadjusted := topHalfSet.List()
	topHalf := []*model.Turtle{}
	for i := 0; i < len(topHalfUnadjusted); i++ {
		for j := 0; j < len(topHalfUnadjusted)-i; j++ {
			topHalf = append(topHalf, topHalfUnadjusted[i])
		}
	}

	for _, m := range ga.models {

		parent1 := topHalf[m.RandomInt(len(topHalf)-1)]
		parent2 := topHalf[m.RandomInt(len(topHalf)-1)]

		dna1 := parent1.GetProperty("dna").([]string)
		dna2 := parent2.GetProperty("dna").([]string)

		robots := m.TurtleBreed("robots")
		robots.CreateAgents(1, func(t *model.Turtle) {
			t.SetXY(0, 0)
			newDNA := ga.blendDNA(dna1, dna2)
			newDNA2 := ga.applyMutations(m, newDNA)
			t.SetProperty("dna", newDNA2)
			t.SetProperty("score", 0)
			t.Color = model.Blue
			t.SetSize(.5)
		})

		cans := m.TurtleBreed("cans")
		cans.Agents().Ask(func(t *model.Turtle) {
			t.SetProperty("picked", false)
		})

		m.ResetTicks()
	}

	allRobots.Ask(func(t *model.Turtle) {
		t.Die()
	})

}

func (ga *GA) blendDNA(dna1 []string, dna2 []string) []string {
	newDNA := []string{}

	// len is 243
	newDNA = dna1[0:121]
	newDNA = append(newDNA, dna2[121:]...)
	return newDNA
}

func (ga *GA) applyMutations(m *model.Model, dna []string) []string {
	// apply mutations to the dna string
	newDNa := []string{}
	for i := 0; i < len(dna); i++ {
		index := m.RandomInt(100)
		newDNa = append(newDNa, dna[i])
		if index < ga.mutationProbability {
			newDNa[i] = possibleActions[m.RandomInt(len(possibleActions)-1)]
		}
	}
	return newDNa
}
