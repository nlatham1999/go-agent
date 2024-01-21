package patch

type Patch struct {
	patchesOwn map[string]interface{}
}

func NewPatch(patchesOwn map[string]interface{}) *Patch {

	patch := &Patch{}

	patch.patchesOwn = map[string]interface{}{}
	for key, value := range patchesOwn {
		patch.patchesOwn[key] = value
	}

	return patch
}
