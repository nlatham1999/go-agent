# Dictionary

### e 
```math.E```

### pi
```math.Pi```

### true
```true```

### false
```false```

### Color Constants
```<Color> (float64)```

### abs
```math.Abs()```

### acos
```math.Acos```


### all?
```Universe.AllLinks(agentset LinkSet, operation TurtleLinkOperation) bool```  
```Universe.AllPatches(agentset PatchSet, operation PatchBoolOperation) bool```  
```Universe.AllTurtles(agentset TurtleSet, operation TurtleBoolOperation) bool```  

### any?
```Universe.AnyLinks(agentset LinkSet, operation TurtleLinkOperation) bool```  
```Universe.AnyPatches(agentset PatchSet, operation PatchBoolOperation) bool```  
```Universe.AnyTurtles(agentset TurtleSet, operation TurtleBoolOperation) bool```  

### approximate-hsb
```ApproximateHSB(hue float64, saturation float64, brightness float64) int```

### approximate-rgb
``` ApproximateRGB(red float64, green float64, blue float64) int```

### ask
```(u *Universe) AskLinks(agentset LinkSet, operations []LinkOperation)```  
```(u *Universe) AskLink(agent *Link, operations []LinkOperation)```
```(u *Universe) AskPatches(agentset PatchSet, operations []PatchOperation)```
```(u *Universe) AskPatch(agent *Patch, operations []PatchOperation)```
```(u *Universe) AskTurtles(agentset TurtleSet, operations []TurtleOperation)```
```(u *Universe) AskTurtle(agent *Turtle, operations []TurtleOperation)```

### at-points
```(t *Base) AtPointsLinks(points []Coordinate) []*Link```  
```(t *Base) AtPointsPatches(points []Coordinate)```  
```(t *Base) AtPointsTurtles(points []Coordinate)```  


### autoplot?
implement?

### auto-plot-off
implement?

### auto-plot-on
implement?

### back, bk
```(t *Turtle) Back(distance float64)```

### base-colors
```BaseColors() []float64```

### beep
implement?

### behaviorspace-experiment-name
implement?

### behaviorspace-run-number
implement?

### both-ends
```(u *Universe) BothEnds(link *Link) []*Turtle```

### breed
```(t *Link) GetBreedName() string```
```(t *Link) GetBreedSet() []*Link```
```(t *Link) SetBreed(name string)```
```(t *Turtle) GetBreedName() string```  
```(t *Turtle) GetBreedSet() []*Turtle```
```(t *Turtle) SetBreed(name string)```

### but-first, butfirst, bf
```ButFirst(arr []interface{}) []interface{}```

### but-last, butlast, bl
```ButLast(arr []interface{}) []interface{}```

### can-move?
```(t *Turtle) CanMove(distance float64) bool```

### carefully
implement?

### clear-all
```(u *Universe) ClearAll()```

### clear-all-plots
implement?

### clear-drawing
implement?

### clear-globals
```(u *Universe) ClearGlobals()```

### clear-links
```(u *Universe) ClearLinks()```

### clear-output
```(u *Universe) ClearOutput()```

### clear-patches, cp
```(u *Universe) ClearPatches()```  

### clear-plot
implement?

### clear-ticks
```(u *Universe) ClearLinks()```

### clear-turtles, ct
```(u *Universe) ClearTurtles()```

### Color
turtles and links have a color struct

### cos
math library

### count
use the len() on the array

### create ordered turtles
```(u *Universe) CreateOrderedTurtles(breed string, amount float64, operations []TurtleOperation)```

### create-\<breed\>-to 
```(t *Turtle) CreateBreedTo(breed string, turtle *Turtle, operations []TurtleOperation)```

### create-\<breeds\>-to
```(t *Turtle) CreateBreedsTo(breed string, turtles []*Turtle, operations []TurtleOperation)```

### create-\<breed\>-from
```(t *Turtle) CreateBreedFrom(breed string, turtle *Turtle, operations []TurtleOperation)```

### create-\<breeds\>-from
```(t *Turtle) CreateBreedsFrom(breed string, turtles []*Turtle, operations []TurtleOperation)```

### create-\<breed\>-with
```(t *Turtle) CreateBreedWith(breed string, turtle *Turtle, operations []TurtleOperation)```

### create-\<breeds\>-with
```(t *Turtle) CreateBreedsWith(breed string, turtles []*Turtle, operations []TurtleOperation)```

### create-link-to
```(t *Turtle) CreateLinkTo(turtle *Turtle, operations []TurtleOperation)```

### create-links-to
```(t *Turtle) CreateLinksTo(turtles []*Turtle, operations []TurtleOperation)```

### create-link-from
```(t *Turtle) CreateLinkFrom(turtle *Turtle, operations []TurtleOperation)```

### create-links-from
```(t *Turtle) CreateLinksFrom(turtles []*Turtle, operations []TurtleOperation)```

### create-link-with
```(t *Turtle) CreateLinkWith(turtle *Turtle, operations []TurtleOperation)```

### create-links-with
```(t *Turtle) CreateLinksWith(turtles []*Turtle, operations []TurtleOperation)```

### create-turtles
(u *Universe) CreateTurtles(amount int, operations []TurtleOperation)

### create-temporary-plot-pen
implement?

### date-and-time
just use the date time library

### die
```(u *Universe) DieTurtle(turtle *Turtle)```
```(u *Universe) DieLink(link *Link)```

### diffuse 
```(u *Universe) Diffuse(patchVariable string, percent float64) error```

### diffuse4
```(u *Universe) Diffuse4(patchVariable string, percent float64) error```

### directed-link-breed
will probably just be a variable in the universe struct

### display
implement?

### distance
```(p *Patch) DistanceTurtle (t *Turtle) float64```
```(p *Patch) DistancePatch(patch *Patch) float64```
```(t *Turtle) DistanceTurtle(turtle *Turtle) float64```
```(t *Turtle) DistancePatch(patch *Patch) float64```

### distancexy
```(p *Patch) DistanceXY(x float64, y float64) float64```
```(p *Turtle) DistanceXY(x float64, y float64) float64```

### downhill
```(t *Turtle) Downhill(patchVariable string)```

### downhll4 
```(t *Turtle) Downhill4(patchVariable string)```

## dx
```(t *Turtle) DX() float64```

### dy
```(t *Turtle) DY() float64```