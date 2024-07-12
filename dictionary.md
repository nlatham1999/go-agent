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
```(t *TurtleAgentSet) AtPoints(u *Universe, points []Coordinate) *TurtleAgentSet```
```(p *PatchAgentSet) AtPoints(u *Universe, points []Coordinate) *PatchAgentSet```


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
```(l *LinkAgentSet) Count()```
```(p *PatchAgentSet) Count() int```
```(t *TurtleAgentSet) Count()```

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
```Universe.DirectedLinkBreeds```

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
```LinkSet(links []*Link) *LinkAgentSet```

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
```Universe.MaxPxCor```

### max-pycor
```Universe.MaxPyCor```

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
```Universe.MinPxCor```

### min-pycor
```Universe.MinPyCor```

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
```(u *Universe) Patch(pxcor float64, pycor float64) *Patch```

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
```Universe.Patches```


### patches-own
```Universe.PatchesOwn```

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
```(u *Universe) ResetTicks()```

### reset-timer
```(u *Universe) ResetTimer()```

### resize-world
```(u *Universe) ResizeWorld(minPxcor int, maxPxcor int, minPycor int, maxPycor int)```

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
```(u *Universe) SetDefaultShapeLinks(shape string)```
```(u *Universe) SetDefaultShapeTurtles(shape string)```
```(u *Universe) SetDefaultShapeLinkBreed(breed string, shape string)```
```(u *Universe) SetDefaultShapeTurtleBreed(breed string, shape string)```

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
```ShadeOf(color1 float64, color2 float64) bool```

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
```(u *Universe) Tick()```

### tick-advance
```(u *Universe) TickAdvance(amount int)```

### ticks
```Universe.Ticks```

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
```(u *Universe) Turtle(breed string) *Turtle```

### turtle-set
```TurtleSet(turtles []*Turtle) *TurtleAgentSet```

### turtles
```Universe.Turtles```

### turtles-at
```(t *TurtleAgentSet) AtPoints(points []Coordinate) *TurtleAgentSet```

### turtles-here
```(p *Patch) TurtlesHere(breed string) *TurtleAgentSet```
```(t *Turtle) TurtlesHere(breed string) *TurtleAgentSet```

### turtles-on
```(u *Universe) TurtlesOnPatch(patch *Patch) *TurtleAgentSet```
```(u *Universe) TurtlesOnPatches(patches *PatchAgentSet) *TurtleAgentSet```
```(u *Universe) TurtlesWithTurtle(turtle *Turtle) *TurtleAgentSet```
```(u *Universe) TurtlesWithTurtles(turtles *TurtleAgentSet) *TurtleAgentSet```

### turtles-own
```Universe.TurtlesOwn```
```Universe.TurtleBreedsOwn```

### type
implement?

### undirected-link-breed
```Universe.UndirectedLinkBreeds```

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
```Universe.WorldWidth```

### world-height
```Universe.WorldHeight```

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
