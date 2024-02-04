//universe functions that deal with turtles

package universe

func (u *Universe) AskTurtles(operations []TurtleOperation) {
	for i := 0; i < len(u.Turtles); i++ {
		for j := 0; j < len(operations); j++ {
			operations[j](u.Turtles[i])
		}
	}
}
