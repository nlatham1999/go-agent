package llm

import (
	"strings"
	"testing"

	"github.com/nlatham1999/go-agent/pkg/model"
	"github.com/nlatham1999/go-agent/pkg/state"
)

type testClient struct {
	responses []string
	prompts   []string
}

func (c *testClient) Complete(agent *Agent, prompt string) (string, error) {
	c.prompts = append(c.prompts, prompt)
	if len(c.responses) == 0 {
		return "", nil
	}
	response := c.responses[0]
	c.responses = c.responses[1:]
	return response, nil
}

func testModel() *model.Model {
	m := model.NewModel(model.ModelSettings{
		MinPxCor: -5,
		MaxPxCor: 5,
		MinPyCor: -5,
		MaxPyCor: 5,
	})
	_, _ = m.CreateTurtles(1, nil)
	return m
}

func testAgent() Agent {
	return Agent{
		Name:            "agent-1",
		TurtleWho:       0,
		AllowedCommands: []string{CommandForward, CommandTurn, CommandGetHeading, CommandGetPosition, CommandGetCurrentPatch, CommandGetPatchInFront, CommandSetPatchColor},
		ActionLimits: ActionLimits{
			MaxSpeed:       2,
			MaxTurnDegrees: 45,
			VisionDistance: 1,
		},
	}
}

func TestExecuteForward(t *testing.T) {
	registry := NewRegistry()
	m := testModel()
	agent := testAgent()

	result, err := registry.Execute(m, agent, CommandForward, []any{1.5})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	position, ok := result.(Position)
	if !ok {
		t.Fatalf("expected Position result, got %T", result)
	}

	if position.X == 0 && position.Y == 0 {
		t.Fatalf("expected turtle to move, got position %+v", position)
	}
}

func TestExecuteRejectsDisallowedCommand(t *testing.T) {
	registry := NewRegistry()
	m := testModel()
	agent := testAgent()
	agent.AllowedCommands = []string{CommandGetHeading}

	_, err := registry.Execute(m, agent, CommandForward, []any{1.0})
	if err == nil {
		t.Fatalf("expected disallowed command error")
	}
}

func TestExecuteRejectsSpeedAboveLimit(t *testing.T) {
	registry := NewRegistry()
	m := testModel()
	agent := testAgent()

	_, err := registry.Execute(m, agent, CommandForward, []any{3.0})
	if err == nil {
		t.Fatalf("expected max speed error")
	}
}

func TestExecuteRejectsTurnAboveLimit(t *testing.T) {
	registry := NewRegistry()
	m := testModel()
	agent := testAgent()

	_, err := registry.Execute(m, agent, CommandTurn, []any{90.0})
	if err == nil {
		t.Fatalf("expected max turn error")
	}
}

func TestExecuteGetPatchInFront(t *testing.T) {
	registry := NewRegistry()
	m := testModel()
	agent := testAgent()

	result, err := registry.Execute(m, agent, CommandGetPatchInFront, []any{1.0})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if _, ok := result.(*state.Patch); !ok {
		t.Fatalf("expected *state.Patch result, got %T", result)
	}
}

func TestExecuteGetCurrentPatch(t *testing.T) {
	registry := NewRegistry()
	m := testModel()
	agent := testAgent()

	result, err := registry.Execute(m, agent, CommandGetCurrentPatch, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if _, ok := result.(*state.Patch); !ok {
		t.Fatalf("expected *state.Patch result, got %T", result)
	}
}

func TestExecuteRejectsPatchInFrontAboveVisionLimit(t *testing.T) {
	registry := NewRegistry()
	m := testModel()
	agent := testAgent()

	_, err := registry.Execute(m, agent, CommandGetPatchInFront, []any{2.0})
	if err == nil {
		t.Fatalf("expected vision distance error")
	}
}

func TestCommandFormats(t *testing.T) {
	registry := NewRegistry()

	formats := registry.CommandFormats([]string{
		CommandTurn,
		"unknown_command",
		CommandGetPosition,
		CommandForward,
	})
	expected := []string{
		CommandTurn + " <float>",
		CommandGetPosition,
		CommandForward + " <float>",
	}

	if len(formats) != len(expected) {
		t.Fatalf("expected %d formats, got %d", len(expected), len(formats))
	}

	for i := range expected {
		if formats[i] != expected[i] {
			t.Fatalf("expected format %q at index %d, got %q", expected[i], i, formats[i])
		}
	}
}

func TestInstructionSuffix(t *testing.T) {
	registry := NewRegistry()
	agent := Agent{
		AllowedCommands: []string{CommandForward, CommandGetHeading},
	}

	suffix := agent.InstructionSuffix(registry)

	expectedParts := []string{
		"Allowed commands:",
		CommandForward + " <float>",
		CommandGetHeading,
		"Respond with only one command per step.",
		"Do not include any explanation or extra text.",
	}

	for _, expected := range expectedParts {
		if !strings.Contains(suffix, expected) {
			t.Fatalf("expected suffix to contain %q, got %q", expected, suffix)
		}
	}
}

func TestExecuteSetPatchColor(t *testing.T) {
	registry := NewRegistry()
	m := testModel()
	agent := testAgent()

	result, err := registry.Execute(m, agent, CommandSetPatchColor, []any{255.0, 0.0, 0.0})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	patch, ok := result.(*state.Patch)
	if !ok {
		t.Fatalf("expected *state.Patch result, got %T", result)
	}

	if patch.Color.Red != 255 || patch.Color.Green != 0 || patch.Color.Blue != 0 {
		t.Fatalf("expected patch color to be rgb(255,0,0), got %+v", patch.Color)
	}
}

func TestExecuteRejectsMissingPatchColorRGBArg(t *testing.T) {
	registry := NewRegistry()
	m := testModel()
	agent := testAgent()

	_, err := registry.Execute(m, agent, CommandSetPatchColor, []any{255.0, 0.0})
	if err == nil {
		t.Fatalf("expected missing rgb argument error")
	}
}

func TestInitialPrompt(t *testing.T) {
	registry := NewRegistry()
	agent := Agent{
		Instruction:     "You are a turtle.",
		AllowedCommands: []string{CommandForward},
	}

	prompt := agent.InitialPrompt(registry)
	if !strings.Contains(prompt, "You are a turtle.") {
		t.Fatalf("expected initial prompt to include instruction, got %q", prompt)
	}
	if !strings.Contains(prompt, CommandForward+" <float>") {
		t.Fatalf("expected initial prompt to include command format, got %q", prompt)
	}
}

func TestAgentAddMessage(t *testing.T) {
	agent := &Agent{}

	agent.AddMessage(RoleUser, "move forward")
	agent.AddMessage(RoleAssistant, "forward 1")

	if len(agent.History) != 2 {
		t.Fatalf("expected 2 history entries, got %d", len(agent.History))
	}

	if agent.History[0].Role != RoleUser || agent.History[0].Content != "move forward" {
		t.Fatalf("unexpected first history entry: %+v", agent.History[0])
	}

	if agent.History[1].Role != RoleAssistant || agent.History[1].Content != "forward 1" {
		t.Fatalf("unexpected second history entry: %+v", agent.History[1])
	}
}

func TestNewAgentBundle(t *testing.T) {
	registry := NewRegistry()
	agents := []Agent{
		{Name: "a1"},
		{Name: "a2"},
	}

	bundle, err := NewAgentBundle(registry, nil, agents)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if bundle.Registry != registry {
		t.Fatalf("expected bundle to use provided registry")
	}

	if len(bundle.Agents) != 2 {
		t.Fatalf("expected 2 agents, got %d", len(bundle.Agents))
	}
	if len(bundle.AgentOrder) != 2 || bundle.AgentOrder[0] != "a1" || bundle.AgentOrder[1] != "a2" {
		t.Fatalf("expected stable agent order [a1 a2], got %+v", bundle.AgentOrder)
	}

	agent, err := bundle.Agent("a1")
	if err != nil {
		t.Fatalf("expected to find agent a1, got %v", err)
	}

	if agent.Name != "a1" {
		t.Fatalf("expected agent name a1, got %q", agent.Name)
	}
}

func TestNewAgentBundleRejectsDuplicateNames(t *testing.T) {
	_, err := NewAgentBundle(nil, nil, []Agent{
		{Name: "dup"},
		{Name: "dup"},
	})
	if err == nil {
		t.Fatalf("expected duplicate agent name error")
	}
}

func TestAgentBundleStep(t *testing.T) {
	client := &testClient{
		responses: []string{"forward 1", "get_position"},
	}
	bundle, err := NewAgentBundle(NewRegistry(), client, []Agent{
		{
			Name:            "agent-1",
			Instruction:     "Move around.",
			TurtleWho:       0,
			AllowedCommands: []string{CommandForward, CommandGetPosition},
			ActionLimits: ActionLimits{
				MaxSpeed: 1,
			},
		},
	})
	if err != nil {
		t.Fatalf("expected no error creating bundle, got %v", err)
	}

	m := testModel()

	if err := bundle.Step(m); err != nil {
		t.Fatalf("expected first step to succeed, got %v", err)
	}

	agent, err := bundle.Agent("agent-1")
	if err != nil {
		t.Fatalf("expected to find agent, got %v", err)
	}

	if len(agent.History) != 3 {
		t.Fatalf("expected history length 3 after first step, got %d", len(agent.History))
	}
	if agent.History[0].Role != RoleSystem {
		t.Fatalf("expected first history entry to be system prompt, got %+v", agent.History[0])
	}
	if agent.History[1].Role != RoleAssistant || agent.History[1].Content != "forward 1" {
		t.Fatalf("unexpected assistant history entry after first step: %+v", agent.History[1])
	}
	if agent.History[2].Role != RoleUser || !strings.Contains(agent.History[2].Content, "Executed command: forward 1") {
		t.Fatalf("unexpected execution result after first step: %+v", agent.History[2])
	}

	if err := bundle.Step(m); err != nil {
		t.Fatalf("expected second step to succeed, got %v", err)
	}

	if len(agent.History) != 5 {
		t.Fatalf("expected history length 5 after second step, got %d", len(agent.History))
	}
	if agent.History[3].Role != RoleAssistant || agent.History[3].Content != "get_position" {
		t.Fatalf("unexpected assistant history entry after second step: %+v", agent.History[3])
	}
	if agent.History[4].Role != RoleUser || !strings.Contains(agent.History[4].Content, "Executed command: get_position") {
		t.Fatalf("unexpected execution result history entry after second step: %+v", agent.History[4])
	}
	if m.Turtle(0).XCor() == 0 && m.Turtle(0).YCor() == 0 {
		t.Fatalf("expected turtle to move after executing first command")
	}
}
