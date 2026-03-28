package llm

import "github.com/nlatham1999/go-agent/pkg/model"

type ActionLimits struct {
	MaxSpeed       float64 // maximum movement magnitude per action; abs(input) must be <= MaxSpeed
	MaxTurnDegrees float64 // abs(input) must be <= MaxTurnDegrees
	VisionDistance float64 // distance used for forward sensing commands
}

type CommandHandler func(*model.Model, Agent, []any) (any, error)

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
