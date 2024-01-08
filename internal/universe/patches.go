//universe functions that deal with patches
package universe

import patch "github.com/nlatham1999/go-agent/internal/patches"

func (u *Universe) AskPatches(operations []patch.PatchOperation) {
	for i := 0; i < len(u.Patches); i++ {
		for j := 0; j < len(operations); j++ {
			operations[j](u.Patches[i])
		}
	}
}
