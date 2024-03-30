package beehive

import (
	"math/rand"
	"time"

	"github.com/nlatham1999/go-agent/internal/slider"
	"github.com/nlatham1999/go-agent/internal/universe"
)

var (
	environment *universe.Universe

	sliders map[string]*slider.Slider

	//globals
	colorList   []float64 // colors for hives, which keeps consistency among the hive colors, plots  pens colors, and committed bees' colors
	qualityList []int     // quality of hives
)

const (
	//patches own

	//widgets
	hiveNumber         = "hiveNumber"
	initialPercentage  = "initialPercentage"
	initialExploreTime = "initialExporePercentage"
	quorum             = "quorum"

	//breeds
	site  = "site"
	scout = "scout"

	//sitesOwn
	quality      = "quality"
	discovered   = "discovered"
	scoutsOnSite = "scoutsOnSite"

	//scoutsOwn
	myHome       = "myHome"       // a bee's original position
	nextTask     = "nextTask"     // the code block a bee is running
	taskString   = "taskString"   // the behavior a bee is displaying
	beeTimer     = "beeTimer"     // a timer keeping track of the length of the current state  or the waiting time before entering next state
	target       = "target"       // the hive that a bee is currently focusing on exploring
	interest     = "interest"     // a bee's interest in the target hive
	trips        = "trips"        // times a bee has visited the target
	initialScout = "initialScout" // true if it is an initial scout, who explores the unknown horizons
	noDiscovery  = "nodiscovery"  // true if it is an initial scout and fails to discover any hive site on its initial exploration
	onSite       = "onSite"       // true if it's inspecting a hive site
	piping       = "piping"       // a bee starts to "pipe" when the decision of the best
	distToHive   = "distToHive"   // the distance between the swarm and the hive that a bee is exploring
	circleSwitch = "circleSwitch" // when making a waggle dance, a bee alternates left and right to make the figure "8". circle-switch alternates between 1 and -1 to tell a bee which direction to turn.
	tempXDance   = "tempXDance"   // initial position of a dance
	tempYDance   = "tempYDance"

	//globals
	showDancePath   = "showDancePath"  // dance path is the circular patter with a zigzag line in the middle. when large amount of bees dance, the patterns overlaps each other, which makes them hard to distinguish. turn show-dance-path? off can clear existing patterns
	scoutsVisible   = "scoutsVisible"  // you can hide scouts and only look at the dance patterns to avoid distraction from bees' dancing movements
	watchDanceTask  = "watchDanceTask" // a list of tasks
	discoverTask    = "discoverTask"
	inspectHiveTask = "inspectHiveTask"
	goHomeTask      = "goHomeTask"
	danceTask       = "danceTask"
	reVisitTask     = "reVisitTask"
	pipeTask        = "pipeTask"
	takeOffTask     = "takeOffTask"
)

func Init() {

	rand.Seed(time.Now().UnixNano())

	sitesOwn := map[string]interface{}{
		quality:      0.0,
		discovered:   false,
		scoutsOnSite: 0.0,
	}

	scoutsOwn := map[string]interface{}{
		myHome:       0.0,
		nextTask:     0,
		taskString:   "",
		beeTimer:     0,
		target:       0,
		interest:     0,
		trips:        0,
		initialScout: false,
		noDiscovery:  false,
		onSite:       false,
		piping:       false,
		distToHive:   0.0,
		circleSwitch: 0,
		tempXDance:   0,
		tempYDance:   0,
	}

	breedsOwn := map[string]map[string]interface{}{
		site:  sitesOwn,
		scout: scoutsOwn,
	}

	environment = universe.NewUniverse(nil, nil, breedsOwn)

	sliders = map[string]*slider.Slider{
		hiveNumber:         slider.NewSlider(4, 1, 10, 10),
		initialPercentage:  slider.NewSlider(5, 1, 25, 12),
		initialExploreTime: slider.NewSlider(100, 10, 300, 200),
		quorum:             slider.NewSlider(0, 1, 50, 49),
	}

}

func setup() {
	// clear-all
	// setup-hives
	// setup-tasks
	// setup-bees
	// set show-dance-path? true
	// set scouts-visible? true
	// reset-ticks

	environment.ClearAll()
}

func setupHives() {
	// set color-list [ 97.9 94.5 57.5 63.8 17.6 14.9 27.5 25.1 117.9 114.4 ]
	// set quality-list [ 100 75 50 1 54 48 40 32 24 16 ]
	// ask n-of hive-number patches with [
	//   distancexy 0 0 > 16 and abs pxcor < (max-pxcor - 2) and
	//   abs pycor < (max-pycor - 2)
	// ] [
	//   ; randomly placing hives around the center in the
	//   ; view with a minimum distance of 16 from the center
	//   sprout-sites 1 [
	// 	set shape "box"
	// 	set size 2
	// 	set color gray
	// 	set discovered? false
	//   ]
	// ]
	// let i 0 ; assign quality and plot pens to each hive
	// repeat count sites [
	//   ask site i [
	// 	set quality item i quality-list
	// 	set label quality
	//   ]
	//   set-current-plot "on-site"
	//   create-temporary-plot-pen word "site" i
	//   set-plot-pen-color item i color-list
	//   set-current-plot "committed"
	//   create-temporary-plot-pen word "target" i
	//   set-plot-pen-color item i color-list
	//   set i i + 1
	// ]

	colorList = []float64{97.9, 94.5, 57.5, 63.8, 17.6, 14.9, 27.5, 25.1, 117.9, 114.4}
	qualityList = []int{100, 75, 50, 1, 54, 48, 40, 32, 24, 16}
}
