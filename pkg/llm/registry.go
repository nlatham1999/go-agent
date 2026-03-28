package llm

import (
	"fmt"
	"math"

	"github.com/nlatham1999/go-agent/pkg/model"
	"github.com/nlatham1999/go-agent/pkg/state"
)

const (
	CommandForward         = "forward"
	CommandTurn            = "turn"
	CommandGetHeading      = "get_heading"
	CommandGetPosition     = "get_position"
	CommandGetCurrentPatch = "get_current_patch"
	CommandGetPatchInFront = "get_patch_in_front"
	CommandSetPatchColor   = "set_patch_color"
)

type Registry struct {
	commands map[string]CommandHandler
}

var commandFormats = map[string]string{
	CommandForward:         CommandForward + " <float>",
	CommandTurn:            CommandTurn + " <float>",
	CommandGetHeading:      CommandGetHeading,
	CommandGetPosition:     CommandGetPosition,
	CommandGetCurrentPatch: CommandGetCurrentPatch,
	CommandGetPatchInFront: CommandGetPatchInFront + " <float>",
	CommandSetPatchColor:   CommandSetPatchColor + " <r> <g> <b>",
}

func NewRegistry() *Registry {
	return &Registry{
		commands: map[string]CommandHandler{
			CommandForward:         forwardCommand,
			CommandTurn:            turnCommand,
			CommandGetHeading:      getHeadingCommand,
			CommandGetPosition:     getPositionCommand,
			CommandGetCurrentPatch: getCurrentPatchCommand,
			CommandGetPatchInFront: getPatchInFrontCommand,
			CommandSetPatchColor:   setPatchColorCommand,
		},
	}
}

func (r *Registry) CommandFormats(commands []string) []string {
	formats := make([]string, 0, len(commands))
	for _, name := range commands {
		if _, ok := r.commands[name]; !ok {
			continue
		}
		format, ok := commandFormats[name]
		if !ok {
			formats = append(formats, name)
			continue
		}
		formats = append(formats, format)
	}
	return formats
}

func (r *Registry) Register(name string, handler CommandHandler) {
	if r.commands == nil {
		r.commands = map[string]CommandHandler{}
	}
	r.commands[name] = handler
}

func (r *Registry) Execute(m *model.Model, agent Agent, command string, args []any) (any, error) {
	if m == nil {
		return nil, fmt.Errorf("model is nil")
	}

	if !agent.CanExecute(command) {
		return nil, fmt.Errorf("command %q is not allowed for agent %q", command, agent.Name)
	}

	handler, ok := r.commands[command]
	if !ok {
		return nil, fmt.Errorf("command %q is not registered", command)
	}

	if args == nil {
		args = []any{}
	}

	return handler(m, agent, args)
}

func getTurtle(m *model.Model, agent Agent) (*model.Turtle, error) {
	turtle := m.Turtle(agent.TurtleWho)
	if turtle == nil {
		return nil, fmt.Errorf("turtle %d not found for agent %q", agent.TurtleWho, agent.Name)
	}
	return turtle, nil
}

func getFloatArg(args []any, index int) (float64, error) {
	if index >= len(args) {
		return 0, fmt.Errorf("missing required argument at index %d", index)
	}
	raw := args[index]

	switch value := raw.(type) {
	case float64:
		return value, nil
	case float32:
		return float64(value), nil
	case int:
		return float64(value), nil
	case int32:
		return float64(value), nil
	case int64:
		return float64(value), nil
	default:
		return 0, fmt.Errorf("argument at index %d must be numeric", index)
	}
}

func forwardCommand(m *model.Model, agent Agent, args []any) (any, error) {
	turtle, err := getTurtle(m, agent)
	if err != nil {
		return nil, err
	}

	value, err := getFloatArg(args, 0)
	if err != nil {
		return nil, err
	}

	if math.Abs(value) > agent.ActionLimits.MaxSpeed {
		return nil, fmt.Errorf("forward value %v exceeds max speed %v", value, agent.ActionLimits.MaxSpeed)
	}

	turtle.Forward(value)
	return Position{
		X: turtle.XCor(),
		Y: turtle.YCor(),
		Z: turtle.ZCor(),
	}, nil
}

func turnCommand(m *model.Model, agent Agent, args []any) (any, error) {
	turtle, err := getTurtle(m, agent)
	if err != nil {
		return nil, err
	}

	value, err := getFloatArg(args, 0)
	if err != nil {
		return nil, err
	}

	if math.Abs(value) > agent.ActionLimits.MaxTurnDegrees {
		return nil, fmt.Errorf("turn value %v exceeds max turn %v", value, agent.ActionLimits.MaxTurnDegrees)
	}

	turtle.SetHeading(turtle.GetHeading() + value)
	return turtle.GetHeading(), nil
}

func getHeadingCommand(m *model.Model, agent Agent, args []any) (any, error) {
	turtle, err := getTurtle(m, agent)
	if err != nil {
		return nil, err
	}

	return turtle.GetHeading(), nil
}

func getPositionCommand(m *model.Model, agent Agent, args []any) (any, error) {
	turtle, err := getTurtle(m, agent)
	if err != nil {
		return nil, err
	}

	return Position{
		X: turtle.XCor(),
		Y: turtle.YCor(),
		Z: turtle.ZCor(),
	}, nil
}

func getCurrentPatchCommand(m *model.Model, agent Agent, args []any) (any, error) {
	turtle, err := getTurtle(m, agent)
	if err != nil {
		return nil, err
	}

	patch := turtle.PatchHere()
	if patch == nil {
		return nil, nil
	}

	return state.FromPatch(patch), nil
}

func getPatchInFrontCommand(m *model.Model, agent Agent, args []any) (any, error) {
	turtle, err := getTurtle(m, agent)
	if err != nil {
		return nil, err
	}

	value, err := getFloatArg(args, 0)
	if err != nil {
		return nil, err
	}

	if math.Abs(value) > agent.ActionLimits.VisionDistance {
		return nil, fmt.Errorf("patch-in-front value %v exceeds vision distance %v", value, agent.ActionLimits.VisionDistance)
	}

	patch := turtle.PatchAhead(value)
	if patch == nil {
		return nil, nil
	}

	return state.FromPatch(patch), nil
}

func setPatchColorCommand(m *model.Model, agent Agent, args []any) (any, error) {
	turtle, err := getTurtle(m, agent)
	if err != nil {
		return nil, err
	}

	red, err := getFloatArg(args, 0)
	if err != nil {
		return nil, err
	}
	green, err := getFloatArg(args, 1)
	if err != nil {
		return nil, err
	}
	blue, err := getFloatArg(args, 2)
	if err != nil {
		return nil, err
	}

	patch := turtle.PatchAhead(1)
	if patch == nil {
		return nil, nil
	}

	patch.Color.SetColorRGB(int(red), int(green), int(blue))
	return state.FromPatch(patch), nil
}
