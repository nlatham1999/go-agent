package patch

type Patch struct {
	x int
	y int

	customVars map[string]interface{}

	topLeft     *Patch
	top         *Patch
	topRight    *Patch
	left        *Patch
	right       *Patch
	bottomLeft  *Patch
	bottom      *Patch
	bottomRight *Patch
}
