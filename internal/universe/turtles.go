//universe functions that deal with turtles

package universe

import turtle "github.com/nlatham1999/go-agent/internal/turtles"

func (u *Universe) AskTurtles(operations []turtle.TurtleOperation) {
	for i := 0; i < len(u.Turtles); i++ {
		for j := 0; j < len(operations); j++ {
			operations[j](u.Turtles[i])
		}
	}
}
