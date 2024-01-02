package patches

type Patches struct {
	x int
	y int

	topLeft     *Patches
	top         *Patches
	topRight    *Patches
	left        *Patches
	right       *Patches
	bottomLeft  *Patches
	bottom      *Patches
	bottomRight *Patches
}
