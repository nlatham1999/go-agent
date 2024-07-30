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
```(l *LinkAgentSet) All(operation LinkBoolOperation) bool```  
```(p *PatchAgentSet) All(operation PatchBoolOperation) bool```  
```(t *TurtleAgentSet) All(operation TurtleBoolOperation) bool```  

### any?
```(l *LinkAgentSet) Any(operation LinkBoolOperation) bool```  
```(p *PatchAgentSet) Any(operation PatchBoolOperation) bool```  
```(t *TurtleAgentSet) Any(operation TurtleBoolOperation) bool```  

### approximate-hsb
no need

### approximate-rgb
no need

### ask
```(m *Model) AskLinks(agentset LinkSet, operations []LinkOperation)```  
```(m *Model) AskLink(agent *Link, operations []LinkOperation)```
```(m *Model) AskPatches(agentset PatchSet, operations []PatchOperation)```
```(m *Model) AskPatch(agent *Patch, operations []PatchOperation)```
```(m *Model) AskTurtles(agentset TurtleSet, operations []TurtleOperation)```
```(m *Model) AskTurtle(agent *Turtle, operations []TurtleOperation)```

### at-points
```(t *TurtleAgentSet) AtPoints(m *Model, points []Coordinate) *TurtleAgentSet```
```(p *PatchAgentSet) AtPoints(m *Model, points []Coordinate) *PatchAgentSet```


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
```(m *Model) BothEnds(link *Link) []*Turtle```

### breed
```(t *Link) BreedName() string```
```(t *Link) Breed() []*Link```
```(t *Link) SetBreed(name string)```
```(t *Turtle) BreedName() string```  
```(t *Turtle) Breed() []*Turtle```
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
```(m *Model) ClearAll()```

### clear-all-plots
implement?

### clear-drawing
implement?

### clear-globals
```(m *Model) ClearGlobals()```

### clear-links
```(m *Model) ClearLinks()```

### clear-output
```(m *Model) ClearOutput()```

### clear-patches, cp
```(m *Model) ClearPatches()```  

### clear-plot
implement?

### clear-ticks
```(m *Model) ClearLinks()```

### clear-turtles, ct
```(m *Model) ClearTurtles()```

### Color
turtles and links have a color struct

### cos
math library

### count
```(l *LinkAgentSet) Count()```
```(p *PatchAgentSet) Count() int```
```(t *TurtleAgentSet) Count()```

### create ordered turtles
```(m *Model) CreateOrderedTurtles(breed string, amount float64, operations []TurtleOperation)```

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
(m *Model) CreateTurtles(amount int, operations []TurtleOperation)

### create-temporary-plot-pen
implement?

### date-and-time
just use the date time library

### die
```(m *Model) DieTurtle(turtle *Turtle)```
```(m *Model) DieLink(link *Link)```

### diffuse 
```(m *Model) Diffuse(patchVariable string, percent float64) error```

### diffuse4
```(m *Model) Diffuse4(patchVariable string, percent float64) error```

### directed-link-breed
```Model.DirectedLinkBreeds```

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
not going to be implemented

### extract-rbg
not going to be implemented

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
not going to be implemented

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

### in-\<breed\>-neighbors?
```(t *Turtle) InLinkNeighbors(turtle *Turtle) []*Turtle```

### in-link-neighbors?
```(t *Turtle) InLinkNeighbors(turtle *Turtle) []*Turtle```

### in-\<breed\>-from
```(t *Turtle) InLinkBreedFrom(breed string, turtle *Turtle) *Link```

### in-link-from
```(t *Turtle) InLinkFrom(turtle *Turtle) *Link```

### __includes
implement?

### in-radius
```(p PatchAgentSet) InRadiusPatch(radius float64, patch *Patch) *PatchAgentSet```
```(p PatchAgentSet) InRadiusTurtle(radius float64, turtle *Turtle) *TurtleAgentSet```
```(t *TurtleAgentSet) InRadiusPatch(radius float64, patch *Patch) *TurtleAgentSet```
```(t *TurtleAgentSet) InRadiusTurtle(radius float64, turtle *Turtle) *TurtleAgentSet```

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
```(m *Model) LayoutCircle(turtles []*Turtle, radius float64)```

### layout-radial
```(m *Model) LayoutRadial(turtles []*Turtle, links []*Link, root *Turtle)```

### layout-spring
```(m *Model) LayoutSpring(turtles []*Turtle, links []*Link, springConstant float64, springLength float64, repulsionConstant float64)```

### layout-tutte
```(m *Model) LayoutTutte(turtles []*Turtle, links []*Link, radius float64)```

### left
```(t *Turtle) Left(number float64)```

### length
use the length of the slice

### let
not necessary

### link
```(m *Model) Link(breed string, turtle1 int, turtle2 int) *Link```
```(m *Model) LinkDirected(breed string, turtle1 int, turtle2 int) *Link```

### link-heading
```(t *Link) Heading() float64```

### link-length
```(t *Link) Length() float64```

### link-set
```LinkSet(links []*Link) *LinkAgentSet```

### link-shapes
```m *Model) LinkShapes() []string```

### links-own
Model.LinksOwn
Model.LinkBreedsOwn

### list
```(l *LinkAgentSet) List() []*Link```

### ln
math library

### log
math library

### loop
for loop

### lput
slice function

### map
pretty sure theres a map library

### max
implement?

### max-n-of
```(l *LinkAgentSet) MaxNOf(n int, operation LinkFloatOperation) *LinkAgentSet```
```(l *PatchAgentSet) MaxNOf(n int, operation PatchFloatOperation) *PatchAgentSet```
```(l *TurtleAgentSet) MaxNOf(n int, operation TurtleFloatOperation) *TurtleAgentSet```

### max-one-of
```(l *LinkAgentSet) MaxOneOf(operation LinkFloatOperation) *Link```
```(p *PatchAgentSet) MaxOneOf(operation PatchFloatOperation) *Patch```
```(t *TurtleAgentSet) MaxOneOf(operation TurtleFloatOperation) *Turtle```

### max-pxcor
```Model.MaxPxCor```

### max-pycor
```Model.MaxPyCor```

### mean
probably a built in function that does this

### median
probably a function that does this

### member?
built in

### min
built in

### min-n-of
```(l *LinkAgentSet) MinNOf(n int, operation LinkFloatOperation) *LinkAgentSet```
```(p *PatchAgentSet) MinNOf(n int, operation PatchFloatOperation) *PatchAgentSet```
```(t *TurtleAgentSet) MinNOf(n int, operation TurtleFloatOperation) *TurtleAgentSet```

### min-one-of
```(l *LinkAgentSet) MinOneOf(operation LinkFloatOperation) *Link```
```(p *PatchAgentSet) MinOneOf(operation PatchFloatOperation) *Patch```
```(t *TurtleAgentSet) MinOneOf(operation TurtleFloatOperation) *Turtle```

### min-pxcor
```Model.MinPxCor```

### min-pycor
```Model.MinPyCor```

### mod
%

### modes
probably built in

### mouse-down?
implement?

### mouse-inside?
implement?

### mouse-xcor
implement?

### mouse-xcor
implement?

### move-to
```(t *Turtle) MoveToPatch(patch *Patch)```
```(t *Turtle) MoveToTurtle(turtle *Turtle)```

### my-links
```(t *Turtle) MyLinks(breed string) []*Link```

### my-in-links
```(t *Turtle) MyInLinks(breed string) []*Link```

### my-out-links
```(t *Turtle) MyOutLinks(breed string) []*Link```

### myself
the turtle, patch, or link is part of the operater parameter so it can be accessed in that way

### n-of
implement?

### n-values
implement?

### neighbors
```(t *Turtle) Neighbors() []*Patch```
```(p *Patch) Neighbors() []*Patch```

### neighbors4
```(t *Turtle) Neighbors4() []*Patch```
```(p *Patch) Neighbors4() []*Patch```

### link-neighbors
```(t *Turtle) LinkNeighbors(breed string) []*Turtle```

### link-neighbor?
```(t *Turtle) LinkNeighbor(turtle *Turtle)```

### new-seed
random library

### no-display
implement?

### nobody
this is nil

### no-links
empty array

### no-patches
empty array

### not
!

### no-turtles
empty array

### of
this is the . operator

### one-of
```(l *LinkAgentSet) OneOf(operation LinkBoolOperation) *Link```
```(p *PatchAgentSet) OneOf() *Patch```
```(t *TurtleAgentSet) OneOf() *Turtle```

### or
||

### other
```(p *Patch) Other(patches *PatchAgentSet) *PatchAgentSet```
```(t *Turtle) Other(turtles TurtleAgentSet) *TurtleAgentSet```


### other-end
```(l *Link) OtherEnd(t *Turtle) *Turtle```
```(t *Turtle) OtherEnd(link *Link) *Turtle```

### out-link-neighbor?
```(t *Turtle) OutLinkNeighbor(breed string, turtle *Turtle) bool```


### out-link-neighbors
```(t *Turtle) OutLinkNeighbors(breed string, turtle *Turtle) *TurtleAgentSet```

### out-link-to
```(t *Turtle) OutLinkTo(breed string, turtle *Turtle) *Link```

### patch
```(m *Model) Patch(pxcor float64, pycor float64) *Patch```

### patch-ahead
```(t *Turtle) PatchAhead(distance float64) *Patch```

### patch-at
```(p *Patch) PatchAt(dx float64, dy float64) *Patch```
```(t *Turtle) PatchAt(dx float64, dy float64) *Patch```

### patch-at-heading-and-distance
```(p *Patch) PatchAtHeadingAndDistance(heading float64, distance float64) *Patch```
```(t *Turtle) PatchAtHeadingAndDistance(heading float64, distance float64) *Patch```

### patch-here
```(t *Turtle) PatchHere() *Patch```

### patch-left-and-ahead
```(t *Turtle) PatchLeftAndAhead(angle float64, distance float64) *Patch```

### patch-right-and-ahead
```(t *Turtle) PatchRightAndAhead(angle float64, distance float64) *Patch```


### patch-set
```PatchSet(patches []*Patch) *PatchAgentSet```

### patch-size
implement?

### patches
```Model.Patches```


### patches-own
```Model.PatchesOwn```

### pcolor
```Patch.PColor```
```Turtle.PatchHere().PColor```

### pen-down
implement?

### pen-erase
implement?

### pen-up
implement?

### pen-mode
implement?

### pen-size
implement?

### plabel
```Patch.Label```
```Turtle.PatchHere().Label```

### plabel-color
```Patch.PLabelColor```
```Turtle.PatchHere().PLabelColor```

### plot
implement?

### plot-name
implement?

### plot-pen-exists?
implement?

### plot-pen-down
implement?

### plot-pen-up
implement?

### plot-pen-reset
implement?

### plotxy
implement?

### plot-x-min
implement?

### plot-x-max
implement?

### plot-y-min
implement?

### plot-y-max
implement?

### position
this is basically a find

### precision
implement?

### print
implement?

### pxcor
```(p *Patch) PXCor() int```
```Turtle.PatchHere().PXCor```

### pycor
```(p *Patch) PYCor() int```
```Turtle.PatchHere().PYCor```

### random
implement?

### random-float
implement?

### random-*
not really needed

### range
for loop

### read-from-string
implement?

### reduce
implement?

### remainder
implement?

### implement?
remove

### remove-duplicates
implement?

### remove-item
implement?

### repeat
implement?

### replace-item
implement?

### report
return

### reset-perspective
implement?

### reset-ticks
```(m *Model) ResetTicks()```

### reset-timer
```(m *Model) ResetTimer()```

### resize-world
```(m *Model) ResizeWorld(minPxcor int, maxPxcor int, minPycor int, maxPycor int)```

### reverse
implement?

### rgb
implement?

### ride
implement?

### ride-me
implement?

### right
```(t *Turtle) Right(number float64)```

### round
implement?

### run
implement?

### self
not needed

### sentence
not needed

### set
not needed

### set-current-directory
implement?

### set-current-plot
implement?

### set-current-plot-pen
implement?

### set-default-shape
```(m *Model) SetDefaultShapeLinks(shape string)```
```(m *Model) SetDefaultShapeTurtles(shape string)```
```(m *Model) SetDefaultShapeLinkBreed(breed string, shape string)```
```(m *Model) SetDefaultShapeTurtleBreed(breed string, shape string)```

### set-histogram-num-bars
implement?

### __set-line-thickness
implement?

### set-patch-size
implement?

### set-plot-background-color
implement?

### set-plot-pen-color
implement?

### set-plot-pen-interval
implement?

### set-plot-pen-mode
implement?

### setup-plots
implement?

### set-plot-x-range
implement?

### set-plot-y-range
implement?

### setxy
```(t *Turtle) SetXY(x float64, y float64)```

### shade-of?
not going to be implemented

### shape
```Link.Shape```
```Turtle.Shape```

### shapes
implement?

### show
implement?

### show-turtle
```(t *Turtle) Show()```

### show-link
```(l *Link) Show()```

### shuffle
implement?

### sin
math library

### size
```Turtle.Size```

### sort
implement?

### sort-by
implement?

### sort-on
implement?

### sprout
```(p *Patch) Sprout(breed string, number int, operations []TurtleOperation)```

### sqrt
math library

### stamp
implement?

### stamp-erase
implement?

### standard-deviation
implement?

### startup
called in main

### stop
break or return

### stop-inspecting
implement?

### stop-inspecting-dead-agents
implement?

### subject
implement?

### sublist
implement?

### substring
implement?

### subtract-headings
```SubtractHeadings(h1 float64, h2 float64) float64```

### sum
implement?

### tan
implement?

### thickness
```Link.Thickness```

### tick
```(m *Model) Tick()```

### tick-advance
```(m *Model) TickAdvance(amount int)```

### ticks
```Model.Ticks```

### tie
```(l *Link) Tie()```

### tie-mode
```Link.TieMode```

### timer
implement?

### to
func

### to-report
func

### towards
```(p *Patch) TowardsTurtle(t *Turtle) float64```
```(p *Patch) TowardsPatch(patch *Patch) float64```
```(t *Turtle) TowardsTurtle(turtle *Turtle)```
```(t *Turtle) TowardsPatch(patch *Patch) float64```

### towardsxy
```(p *Patch) TowardsXY(x float64, y float64) float64```
```(t *Turtle) TowardsXY(x float64, y float64) float64```

### turtle
```(m *Model) Turtle(breed string) *Turtle```

### turtle-set
```TurtleSet(turtles []*Turtle) *TurtleAgentSet```

### turtles
```Model.Turtles```

### turtles-at
```(t *TurtleAgentSet) AtPoints(points []Coordinate) *TurtleAgentSet```

### turtles-here
```(p *Patch) TurtlesHere(breed string) *TurtleAgentSet```
```(t *Turtle) TurtlesHere(breed string) *TurtleAgentSet```

### turtles-on
```(m *Model) TurtlesOnPatch(patch *Patch) *TurtleAgentSet```
```(m *Model) TurtlesOnPatches(patches *PatchAgentSet) *TurtleAgentSet```
```(m *Model) TurtlesWithTurtle(turtle *Turtle) *TurtleAgentSet```
```(m *Model) TurtlesWithTurtles(turtles *TurtleAgentSet) *TurtleAgentSet```

### turtles-own
```Model.TurtlesOwn```
```Model.TurtleBreedsOwn```

### type
implement?

### undirected-link-breed
```Model.UndirectedLinkBreeds```

### untie
```(l *Link) Untie()```

### up-to-n-of
```(l *LinkAgentSet) UpToNOf() *LinkAgentSet```
```(l *PatchAgentSet) UpToNOf(n int) *PatchAgentSet```
```(l *TurtleAgentSet) UpToNOf(n int) *TurtleAgentSet```

### update-plots
implement?

### uphill
```(t *Turtle) Uphill(patchVariable string)```

### uphill4
```(t *Turtle) Uphill4(patchVariable string)```

### user-*
implement?

### variance
implement?

### wait
implement?

### watch
implement?

### watch-me
implement?

### while
implement?

### who
```Turtle.Who```

### who-are-not
```(l *LinkAgentSet) WhoAreNot(links *LinkAgentSet) *LinkAgentSet```
```(l *LinkAgentSet) WhoAreNotLink(link *Link) *LinkAgentSet```
```(p *PatchAgentSet) WhoAreNot(patches *PatchAgentSet) *PatchAgentSet```
```(p *PatchAgentSet) WhoAreNotPatch(patch *Patch) *PatchAgentSet```
```(t *TurtleAgentSet) WhoAreNot(turtles *TurtleAgentSet) *TurtleAgentSet```
```(t *TurtleAgentSet) WhoAreNotTurtle(turtle *Turtle) *TurtleAgentSet```

### with
```(l *LinkAgentSet) With(operation LinkBoolOperation) *LinkAgentSet```
```(p *PatchAgentSet) With(operation PatchBoolOperation) *PatchAgentSet```
```(t *TurtleAgentSet) With(operation TurtleBoolOperation) *TurtleAgentSet```

### link-with
```(t *Turtle) LinkWith(turtle *Turtle) *Link```

### with-max
```(l *LinkAgentSet) WithMax(operation LinkFloatOperation) *LinkAgentSet```
```(p *PatchAgentSet) WithMax(operation PatchFloatOperation) *PatchAgentSet```
```(t *TurtleAgentSet) WithMax(operation TurtleFloatOperation) *TurtleAgentSet```

### with-min
```(l *LinkAgentSet) WithMin(operation LinkFloatOperation) *LinkAgentSet```
```(p *PatchAgentSet) WithMin(operation PatchFloatOperation) *PatchAgentSet```
```(t *TurtleAgentSet) WithMin(operation TurtleFloatOperation) *TurtleAgentSet```


### with-local-randomness
implement?

### without-interruption
implement?

### word
implement?

### world-width
```Model.WorldWidth```

### world-height
```Model.WorldHeight```

### wrap-color
```WrapColor(color float64) float64```

### write
implement?

### xcor
```(t *Turtle) XCor() float64 ```

### xor
not needed

### ycor
```(t *Turtle) YCor() float64```
