package converter

type astType string

var (
	expression astType = "expression"
	value      astType = "value"
)

type ast struct {
	token    token
	astType  astType
	children []ast
}

func (a *ast) print(tabs int) {
	for i := 0; i < tabs; i++ {
		print("\t")
	}
	println(a.token.tokenType, a.token.lexeme)
	for _, child := range a.children {
		child.print(tabs + 1)
	}
}
