package model

import "fmt"

var (
	ErrNoLinksInAgentSet   = fmt.Errorf("no links in agent set")
	ErrNoTurtlesInAgentSet = fmt.Errorf("no turtles in agent set")
	ErrNoPatchesInAgentSet = fmt.Errorf("no patches in agent set")
)
