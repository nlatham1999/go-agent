package llm

import "strings"

type Agent struct {
	Name            string
	Instruction     string
	TurtleWho       int
	AllowedCommands []string
	ActionLimits    ActionLimits
	History         []Message
	LastResponseID  string
}

func (a Agent) CanExecute(command string) bool {
	for _, allowed := range a.AllowedCommands {
		if allowed == command {
			return true
		}
	}
	return false
}

func (a *Agent) AddMessage(role, content string) {
	if a == nil {
		return
	}
	a.History = append(a.History, Message{
		Role:    role,
		Content: content,
	})
}

func (a Agent) InstructionSuffix(registry *Registry) string {
	if registry == nil {
		return "Respond with only one command per step."
	}

	commandFormats := registry.CommandFormats(a.AllowedCommands)
	if len(commandFormats) == 0 {
		return "Respond with only one command per step."
	}

	return "Allowed commands:\n- " +
		strings.Join(commandFormats, "\n- ") +
		"\nRespond with only one command per step. Do not include any explanation or extra text."
}

func (a Agent) InitialPrompt(registry *Registry) string {
	suffix := a.InstructionSuffix(registry)
	if a.Instruction == "" {
		return suffix
	}
	if suffix == "" {
		return a.Instruction
	}
	return a.Instruction + "\n\n" + suffix
}
