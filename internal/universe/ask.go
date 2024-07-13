// universe functions that deal with patches
package universe

//ask functions

func AskLinks(agentset LinkAgentSet, operations []LinkOperation) {
	for link := range agentset.links {
		for j := 0; j < len(operations); j++ {
			operations[j](link)
		}
	}
}

func AskLink(agent *Link, operations []LinkOperation) {
	for j := 0; j < len(operations); j++ {
		operations[j](agent)
	}
}

func AskPatches(agentset *PatchAgentSet, operations []PatchOperation) {
	for patch := range agentset.patches {
		for j := 0; j < len(operations); j++ {
			operations[j](patch)
		}
	}
}

func AskPatch(agent *Patch, operations []PatchOperation) {
	for j := 0; j < len(operations); j++ {
		operations[j](agent)
	}
}

func AskTurtles(agentset *TurtleAgentSet, operations []TurtleOperation) {
	for turtle := range agentset.turtles {
		for j := 0; j < len(operations); j++ {
			operations[j](turtle)
		}
	}
}

func AskTurtle(agent *Turtle, operations []TurtleOperation) {
	for j := 0; j < len(operations); j++ {
		operations[j](agent)
	}
}
