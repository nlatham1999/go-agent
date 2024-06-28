//universe functions that deal with patches
package universe

//ask functions

func (u *Universe) AskLinks(agentset LinkSet, operations []LinkOperation) {
	for i := 0; i < len(agentset); i++ {
		for j := 0; j < len(operations); j++ {
			operations[j](agentset[i])
		}
	}
}

func (u *Universe) AskLink(agent *Link, operations []LinkOperation) {
	for j := 0; j < len(operations); j++ {
		operations[j](agent)
	}
}

func (u *Universe) AskPatches(agentset PatchSet, operations []PatchOperation) {
	for i := 0; i < len(agentset); i++ {
		for j := 0; j < len(operations); j++ {
			operations[j](agentset[i])
		}
	}
}

func (u *Universe) AskPatch(agent *Patch, operations []PatchOperation) {
	for j := 0; j < len(operations); j++ {
		operations[j](agent)
	}
}

func (u *Universe) AskTurtles(agentset TurtleSet, operations []TurtleOperation) {
	for i := 0; i < len(agentset); i++ {
		for j := 0; j < len(operations); j++ {
			operations[j](agentset[i])
		}
	}
}

func (u *Universe) AskTurtle(agent *Turtle, operations []TurtleOperation) {
	for j := 0; j < len(operations); j++ {
		operations[j](agent)
	}
}
