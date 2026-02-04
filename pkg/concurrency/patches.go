package concurrency

import (
	"sync"

	"github.com/nlatham1999/go-agent/pkg/model"
)

func AskPatches(patchList []*model.Patch, operation model.PatchOperation, numGoRoutines int) {
	if len(patchList) == 0 {
		return
	}

	patchChannel := loadPatchChannel(patchList)

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	for i := 0; i < numGoRoutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for patch := range patchChannel {
				operation(patch)
			}
		}()
	}
	// Wait for all goroutines to finish
	wg.Wait()
}

func loadPatchChannel(patchList []*model.Patch) <-chan *model.Patch {
	patchChannel := make(chan *model.Patch, len(patchList))

	go func() {
		for _, patch := range patchList {
			patchChannel <- patch
		}
		close(patchChannel)
	}()

	return patchChannel
}
