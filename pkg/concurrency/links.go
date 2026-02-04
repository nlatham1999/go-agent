package concurrency

import (
	"sync"

	"github.com/nlatham1999/go-agent/pkg/model"
)

func AskLinks(links []*model.Link, operation model.LinkOperation, numGoRoutines int) {
	if len(links) == 0 {
		return
	}

	linkChannel := loadLinkChannel(links)

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	for i := 0; i < numGoRoutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for link := range linkChannel {
				operation(link)
			}
		}()
	}
	// Wait for all goroutines to finish
	wg.Wait()
}

func loadLinkChannel(links []*model.Link) <-chan *model.Link {
	linkChannel := make(chan *model.Link, len(links))

	go func() {
		for _, link := range links {
			linkChannel <- link
		}
		close(linkChannel)
	}()

	return linkChannel
}
