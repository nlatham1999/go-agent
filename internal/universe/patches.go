//universe functions that deal with patches
package universe

func (u *Universe) AskPatches(operations []PatchOperation) {
	for y := 0; y < len(u.PatchesArray2D); y++ {
		for x := 0; x < len(u.PatchesArray2D[y]); x++ {
			for j := 0; j < len(operations); j++ {
				operations[j](u.PatchesArray2D[y][x])
			}
		}
	}
}
