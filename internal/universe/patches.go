//universe functions that deal with patches
package universe

func (u *Universe) AskPatches(operations []PatchOperation) {
	for i := 0; i < len(u.Patches); i++ {
		for j := 0; j < len(operations); j++ {
			operations[j](u.Patches[i])
		}
	}
}
