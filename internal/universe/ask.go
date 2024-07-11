// universe functions that deal with patches
package universe

//ask functions

func (u *Universe) AskLinks(agentset LinkAgentSet, operations []LinkOperation) {
	for i := 0; i < len(agentset.links); i++ {
		for j := 0; j < len(operations); j++ {
			operations[j](agentset.links[i])
		}
	}
}

func (u *Universe) AskLink(agent *Link, operations []LinkOperation) {
	for j := 0; j < len(operations); j++ {
		operations[j](agent)
	}
}

func (u *Universe) AskPatches(agentset *PatchAgentSet, operations []PatchOperation) {
	for i := 0; i < len(agentset.patches); i++ {
		for j := 0; j < len(operations); j++ {
			operations[j](agentset.patches[i])
		}
	}
}

func (u *Universe) AskPatch(agent *Patch, operations []PatchOperation) {
	for j := 0; j < len(operations); j++ {
		operations[j](agent)
	}
}

func (u *Universe) AskTurtles(agentset *TurtleAgentSet, operations []TurtleOperation) {
	for i := 0; i < len(agentset.turtles); i++ {
		for j := 0; j < len(operations); j++ {
			operations[j](agentset.turtles[i])
		}
	}
}

func (u *Universe) AskTurtle(agent *Turtle, operations []TurtleOperation) {
	for j := 0; j < len(operations); j++ {
		operations[j](agent)
	}
}
