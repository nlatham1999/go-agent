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

### empty?
use len()

### end
not needed

### end1
```Link.End1```

### end2
```Link.End2```

### error
use the error library

### every
use go routines

### exp
math library

### export-*
not needed

### extensions
not needed

### extract-hsb
```ExtractHSBFromScale(scale float64) (int, int, int)```
```ExtractHSBFromRBG(red int, green int, blue int) (int, int, int)```

### extract-rbg
```ExtractRGBFromScale(scale float64) (int, int, int)```

### face
```(t *Turtle) FaceTurtle(turtle *Turtle)```
```(t *Turtle) FacePatch(patch *Patch)```


### facexy
```(t *Turtle) FaceXY(x float64, y float64)```

### file-at-end?
implement?

### file-close
implement?

### file-close-all
implement?

### file-delete
implement?

### file-exists?
implement?

### file-flush
implement?

### file-open
implement?

### file-print
implement?

### file-read
implement?

### file-read-characters
implement?

### file-read-line
implement?

### file-show
implement?

### file-type
implement?

### file-write
implement?

### filter
```Filter(arr []interface{}, pred func(interface{}) bool)```

### first
use the 0th position

### floor
math library

### follow
implement?

### follow-me
implement?

### foreach
for loops

### forward
```t *Turtle) Forward(distance float64)```

### fput
we have array concatonation

### globals
```GlobalFloats map[string]float64```
```GlobalBools  map[string]bool```
ummm why can't we just use them as regulary globals

### hatch
```(t *Turtle) Hatch(amount int, operations []TurtleOperation)```
```(t *Turtle) HatchBreed(breed string, amount int, operations []TurtleOperation)```

### heading
```Turtle.Heading```

### hidden?
```Link.Hidden```
```Turtle.Hidden```

### hide-link  
```(t *Link) Hide()```

### hide-turtle
```(t *Turtle) Hide()```

### histogram
implement?

### home
```(t *Turtle) Home()```

### hsb
```HSB(hue int, saturation int, brightness int) (int, int, int)```

### hubnet-*
implement?

### if
if statement

### ifelse
if else statement

### ifelse-value
switch case statement

### import-*
implement?

### in-cone
```(t *Turtle) InConePatches(distance float64, angle float64) []*Patch```
```(t *Turtle) InConeTurtles(distance float64, angle float64) []*Turtle```

### in-\<breed\>-neighbor?
```(t *Turtle) InLinkBreedNeighbor(breed string, turtle *Turtle)```

### in-link-neighbor?
```(t *Turtle) InLinkNeighbor(turtle *Turtle) bool```

### in-\<breed\>-neighbor?
```(t *Turtle) InLinkBreedNeighbors(breed string, turtle *Turtle) []*Turtle```

### in-link-neighbor?
```(t *Turtle) InLinkNeighbors(turtle *Turtle) []*Turtle```

### in-\<breed\>-from
```(t *Turtle) InLinkBreedFrom(breed string, turtle *Turtle) *Link```

### in-link-from
```(t *Turtle) InLinkFrom(turtle *Turtle) *Link```

### __includes
implement?

### in-radius
```(p *Patch) InRadiusPatches(radius float64) []*Patch```
```(p *Patch) InRadiusTurtles(radius float64) []*Turtle```
```(t *Turtle) InRadiusPatches(distance float64) []*Patch```
```(t *Turtle) InRadiusTurtles(distance float64) []*Turtle```

### insert-item
use slice functions

### inspect
implement?

### int
just convert to int

### is-*?
we use strict typing so not really necessary except for directed and breeds
```Link.Directed```
```Link.Breed```
```Turtle.Breed```

### item
we have indeces on slices

### jump
```(t *Turtle) Jump(distance float64)```

### label
```Link.Label```
```Turtle.Label```

### label-color
```Link.LabelColor```
```Turtle.LabelColor```

### last
last element in slice

### layout-circle
```(u *Universe) LayoutCircle(turtles []*Turtle, radius float64)```

### layout-radial
```(u *Universe) LayoutRadial(turtles []*Turtle, links []*Link, root *Turtle)```

### layout-spring
```(u *Universe) LayoutSpring(turtles []*Turtle, links []*Link, springConstant float64, springLength float64, repulsionConstant float64)```

### layout-tutte
```(u *Universe) LayoutTutte(turtles []*Turtle, links []*Link, radius float64)```

### left
```(t *Turtle) Left(number float64)```

### length
use the length of the slice

### let
not necessary

### link
```(u *Universe) Link(breed string, turtle1 int, turtle2 int) *Link```
```(u *Universe) LinkDirected(breed string, turtle1 int, turtle2 int) *Link```

### link-heading
```(t *Link) Heading() float64```

### link-length
```(t *Link) Length() float64```

### link-set
implement?

### link-shapes
```u *Universe) LinkShapes() []string```

### links-own
Universe.LinksOwn
Universe.LinkBreedsOwn

### list
implement?

### ln
math library

### log
math library

### loop
for loop

### lput
slice function