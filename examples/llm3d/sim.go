package llm3d

import (
	"fmt"
	"log"
	"os"

	"github.com/nlatham1999/go-agent/pkg/api"
	"github.com/nlatham1999/go-agent/pkg/llm"
	"github.com/nlatham1999/go-agent/pkg/model"
)

var _ api.ModelInterface = (*Sim)(nil)

const (
	llm3dPromptID      = "pmpt_69b61b4a64f88193820c245a95ece9970855413b5f098b65"
	llm3dPromptVersion = "1"
)

type Sim struct {
	model  *model.Model
	bundle *llm.AgentBundle

	lastError string
}

func NewSim() *Sim {
	return &Sim{}
}

func (s *Sim) Init() {
	settings := model.ModelSettings{
		MinPxCor: 0,
		MaxPxCor: 9,
		MinPyCor: 0,
		MaxPyCor: 9,
		MinPzCor: 0,
		MaxPzCor: 9,
	}
	s.model = model.NewModel(settings)
}

func (s *Sim) SetUp() error {
	s.model.ClearAll()
	s.lastError = ""
	log.Println("[llm3d] setup starting")

	_, err := s.model.CreateTurtles(1, func(t *model.Turtle) {
		t.SetXYZ(5, 5, 5)
		t.Color.SetColor(model.White)
		t.SetHeading(0)
		t.Shape = "triangle"
	})
	if err != nil {
		return err
	}

	client := llm.NewChatGPTClient(os.Getenv("OPENAI_API_KEY"), os.Getenv("OPENAI_MODEL"))
	client.PromptID = llm3dPromptID
	client.PromptVersion = llm3dPromptVersion
	bundle, err := llm.NewAgentBundle(llm.NewRegistry(), client, []llm.Agent{
		{
			Name:        "pilot",
			Instruction: "You are controlling a single turtle in a 10x10x10 world. Explore carefully, have fun, and build things by coloring patches with set_patch_color. Use sensing commands often, make small moves, and create visible structures or trails in the world.",
			TurtleWho:   0,
			AllowedCommands: []string{
				llm.CommandForward,
				llm.CommandTurn,
				llm.CommandGetHeading,
				llm.CommandGetPosition,
				llm.CommandGetCurrentPatch,
				llm.CommandGetPatchInFront,
				llm.CommandSetPatchColor,
			},
			ActionLimits: llm.ActionLimits{
				MaxSpeed:       1,
				MaxTurnDegrees: 45,
				VisionDistance: 2,
			},
		},
	})
	if err != nil {
		return err
	}

	s.bundle = bundle
	s.model.ResetTicks()
	log.Println("[llm3d] setup complete")
	return nil
}

func (s *Sim) Go() {
	if s.bundle == nil {
		s.lastError = "agent bundle is nil"
		log.Printf("[llm3d] tick=%d error=%s\n", s.model.Ticks, s.lastError)
		return
	}

	if err := s.bundle.Step(s.model); err != nil {
		s.lastError = err.Error()
		log.Printf("[llm3d] tick=%d error=%s\n", s.model.Ticks, s.lastError)
		return
	}

	turtle := s.model.Turtle(0)
	if turtle != nil {
		log.Printf(
			"[llm3d] tick=%d pos=(%.2f, %.2f, %.2f) heading=%.2f history=%d lastCommand=%q\n",
			s.model.Ticks,
			turtle.XCor(),
			turtle.YCor(),
			turtle.ZCor(),
			turtle.GetHeading(),
			s.historyLen(),
			s.lastAssistantCommand(),
		)
	}

	s.lastError = ""
	s.model.Tick()
}

func (s *Sim) Model() *model.Model {
	return s.model
}

func (s *Sim) Stats() map[string]interface{} {
	if s.model == nil {
		return map[string]interface{}{
			"lastError": "model is nil",
		}
	}

	turtle := s.model.Turtle(0)
	if turtle == nil {
		return map[string]interface{}{
			"lastError": s.lastError,
		}
	}

	historyLen := 0
	if s.bundle != nil {
		agent, err := s.bundle.Agent("pilot")
		if err == nil {
			historyLen = len(agent.History)
		}
	}

	return map[string]interface{}{
		"x":          fmt.Sprintf("%.2f", turtle.XCor()),
		"y":          fmt.Sprintf("%.2f", turtle.YCor()),
		"z":          fmt.Sprintf("%.2f", turtle.ZCor()),
		"heading":    fmt.Sprintf("%.2f", turtle.GetHeading()),
		"historyLen": historyLen,
		"lastError":  s.lastError,
	}
}

func (s *Sim) Stop() bool {
	return false
}

func (s *Sim) Widgets() []api.Widget {
	return nil
}

func (s *Sim) historyLen() int {
	if s.bundle == nil {
		return 0
	}
	agent, err := s.bundle.Agent("pilot")
	if err != nil {
		return 0
	}
	return len(agent.History)
}

func (s *Sim) lastAssistantCommand() string {
	if s.bundle == nil {
		return ""
	}
	agent, err := s.bundle.Agent("pilot")
	if err != nil {
		return ""
	}
	for i := len(agent.History) - 1; i >= 0; i-- {
		if agent.History[i].Role == llm.RoleAssistant {
			return agent.History[i].Content
		}
	}
	return ""
}
