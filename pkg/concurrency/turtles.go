package concurrency

import (
	"sync"

	"github.com/nlatham1999/go-agent/pkg/model"
)

func AskTurtles(turtleList []*model.Turtle, operation model.TurtleOperation, numGoRoutines int) {

	if len(turtleList) == 0 {
		return
	}

	turtleChannel := loadTurtleChannel(turtleList)

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	for i := 0; i < numGoRoutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for turtle := range turtleChannel {
				operation(turtle)
			}
		}()
	}
	// Wait for all goroutines to finish
	wg.Wait()
}

func loadTurtleChannel(turtleList []*model.Turtle) <-chan *model.Turtle {

	channel := make(chan *model.Turtle, len(turtleList))

	go func() {
		for _, turtle := range turtleList {
			channel <- turtle
		}
		close(channel)
	}()

	return channel
}
