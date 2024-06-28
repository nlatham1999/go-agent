package universe

//agentset types

//set of links
type LinkSet []*Link

//set of Patches
type PatchSet []*Patch

//set of Turtles
type TurtleSet []*Turtle

//operation types

//general function for acting on a link
type LinkOperation func(t *Link)

//general function that takes in a patch and returns a bool
type LinkBoolOperation func(t *Link) bool

//general function for acting on a patch
type PatchOperation func(t *Patch)

//general function that takes in a patch and returns a bool
type PatchBoolOperation func(t *Patch) bool

//general function for acting on a turtle
type TurtleOperation func(t *Turtle)

//general function that takes in a turtle and returns a bool
type TurtleBoolOperation func(t *Turtle) bool
