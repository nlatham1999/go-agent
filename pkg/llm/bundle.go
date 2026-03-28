package llm

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nlatham1999/go-agent/pkg/model"
)

type Client interface {
	Complete(agent *Agent, prompt string) (string, error)
}

// AgentBundle holds the shared runtime dependencies for a group of agents.
// Agents are stored both by name for lookup and in a stable slice order for stepping.
type AgentBundle struct {
	Registry   *Registry
	Client     Client
	Agents     map[string]*Agent
	AgentOrder []string
}

// NewAgentBundle builds a bundle with deterministic step order matching the input slice.
func NewAgentBundle(registry *Registry, client Client, agents []Agent) (*AgentBundle, error) {
	if registry == nil {
		registry = NewRegistry()
	}

	agentMap := make(map[string]*Agent, len(agents))
	agentOrder := make([]string, 0, len(agents))
	for _, agent := range agents {
		if agent.Name == "" {
			return nil, fmt.Errorf("agent name cannot be empty")
		}
		if _, exists := agentMap[agent.Name]; exists {
			return nil, fmt.Errorf("duplicate agent name %q", agent.Name)
		}

		agentCopy := agent
		agentMap[agent.Name] = &agentCopy
		agentOrder = append(agentOrder, agent.Name)
	}

	return &AgentBundle{
		Registry:   registry,
		Client:     client,
		Agents:     agentMap,
		AgentOrder: agentOrder,
	}, nil
}

// Agent returns a named agent from the bundle.
func (b *AgentBundle) Agent(name string) (*Agent, error) {
	if b == nil {
		return nil, fmt.Errorf("agent bundle is nil")
	}

	agent, ok := b.Agents[name]
	if !ok {
		return nil, fmt.Errorf("agent %q not found", name)
	}
	return agent, nil
}

// Step advances every agent by one LLM round.
// On the first round, the agent instruction is used as the prompt.
// On later rounds, the previous execution result is used as the next prompt.
func (b *AgentBundle) Step(m *model.Model) error {
	if b == nil {
		return fmt.Errorf("agent bundle is nil")
	}
	if m == nil {
		return fmt.Errorf("model is nil")
	}
	if b.Registry == nil {
		return fmt.Errorf("registry is nil")
	}
	if b.Client == nil {
		return fmt.Errorf("client is nil")
	}

	for _, name := range b.AgentOrder {
		agent := b.Agents[name]
		prompt := lastUserPrompt(agent)
		if prompt == "" {
			prompt = agent.InitialPrompt(b.Registry)
			agent.AddMessage(RoleSystem, prompt)
		}

		response, err := b.Client.Complete(agent, prompt)
		if err != nil {
			return fmt.Errorf("agent %q: %w", agent.Name, err)
		}

		agent.AddMessage(RoleAssistant, response)

		result, err := b.executeMessage(m, agent, response)
		if err != nil {
			return fmt.Errorf("agent %q: %w", agent.Name, err)
		}

		agent.AddMessage(RoleUser, formatExecutionResult(response, result))
	}

	return nil
}

// executeMessage parses and runs a plain text command returned by the LLM.
func (b *AgentBundle) executeMessage(m *model.Model, agent *Agent, message string) (any, error) {
	command, args, err := parseCommand(message)
	if err != nil {
		return nil, err
	}
	return b.Registry.Execute(m, *agent, command, args)
}

func lastUserPrompt(agent *Agent) string {
	if agent == nil {
		return ""
	}

	for i := len(agent.History) - 1; i >= 0; i-- {
		if agent.History[i].Role == RoleUser {
			return agent.History[i].Content
		}
	}
	return ""
}

// parseCommand converts a plain text response like "forward 1" into a command name and positional args.
func parseCommand(message string) (string, []any, error) {
	fields := strings.Fields(message)
	if len(fields) == 0 {
		return "", nil, fmt.Errorf("empty command")
	}

	command := fields[0]
	args := make([]any, 0, len(fields)-1)
	for _, field := range fields[1:] {
		value, err := strconv.ParseFloat(field, 64)
		if err != nil {
			args = append(args, field)
			continue
		}
		args = append(args, value)
	}

	return command, args, nil
}

// formatExecutionResult creates the feedback message that is fed into the next LLM turn.
func formatExecutionResult(command string, result any) string {
	if result == nil {
		return fmt.Sprintf("Executed command: %s\nResult: null", command)
	}
	return fmt.Sprintf("Executed command: %s\nResult: %v", command, result)
}
