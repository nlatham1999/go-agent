package converter

import "fmt"

type parseExpression struct {
	subTokens []validTokens
}

type validTokens struct {
	validType string
	tokens    map[string]interface{}
}

func parse(tokens []token) (ast, error) {

	// create the root of the AST

	//root token
	rootToken := token{
		tokenType: "ROOT",
	}

	index := 0

	if tokens[0].tokenType != "ROOT" {
		tokens = append([]token{rootToken}, tokens...)
	}

	root, err := parseExp(tokens, &index)
	if err != nil {
		return ast{}, err
	}

	return root, nil
}

func parseExp(tokens []token, index *int) (ast, error) {

	a := ast{
		token:    tokens[*index],
		astType:  expression,
		children: []ast{},
	}

	exp := parseMap[a.token.tokenType]
	if exp == nil {
		return ast{}, fmt.Errorf("Found non parsable token: %s, line: %d, column: %d", a.token.tokenType, a.token.line, a.token.column)
	}

	for _, valids := range exp.subTokens {
		if valids.validType == "any" {
			*index++
			if *index >= len(tokens) {
				break
			}
			nextToken := tokens[*index]
			if _, ok := valids.tokens[nextToken.tokenType]; ok {
				c, err := parseExp(tokens, index)
				if err != nil {
					return ast{}, err
				}
				a.children = append(a.children, c)
			} else {
				return ast{}, fmt.Errorf("valid next token not found. Current token: %s, line: %d, column: %d. Valid next tokens: %v. Given next token: %v", tokens[*index-1].tokenType, tokens[*index-1].line, tokens[*index-1].column, valids.tokens, tokens[*index])
			}
		} else if valids.validType == "many" {

			childrenFound := 0
			found := true
			for found {
				*index++
				if *index >= len(tokens) {
					break
				}
				nextToken := tokens[*index]
				if _, ok := valids.tokens[nextToken.tokenType]; ok {
					c, err := parseExp(tokens, index)
					if err != nil {
						return ast{}, err
					}
					a.children = append(a.children, c)
					childrenFound++

				} else {
					if childrenFound == 0 {
						return ast{}, fmt.Errorf("valid next token not found. Current token: %s, line: %d, column: %d. Valid next tokens in many: %v. Given next token: %v", tokens[*index-1].tokenType, tokens[*index-1].line, tokens[*index-1].column, valids.tokens, tokens[*index])
					}
					found = false
					*index--
				}
			}

		} else if valids.validType == "skip" {
			*index++
		} else {
			panic("not valid valid type")
		}
	}

	if len(exp.subTokens) == 0 {
		a.astType = value
	} else {
		a.astType = expression
	}

	return a, nil
}

var rootExp = parseExpression{
	subTokens: []validTokens{
		{
			validType: "many",
			tokens:    map[string]interface{}{"BREED": nil, "GLOBALS": nil, "TO": nil},
		},
	},
}

var breedExp = parseExpression{
	subTokens: []validTokens{
		openBracketValidType,
		varValidType,
		varValidType,
		closeBracketValidType,
	},
}

var singleValueExpr = parseExpression{
	subTokens: []validTokens{},
}

var globalsExpr = parseExpression{
	subTokens: []validTokens{
		openBracketValidType,
		varValidType,
		varValidType,
		varValidType,
		varValidType,
		closeBracketValidType,
	},
}

var toExpr = parseExpression{
	subTokens: []validTokens{
		varValidType,
		toBodyValidType,
		endValidType,
	},
}

var setDefaultShapeExpr = parseExpression{
	subTokens: []validTokens{
		breedsTurtlesLinksValidType,
		stringValidType,
	},
}

var setExpr = parseExpression{
	subTokens: []validTokens{
		lValueValidType,
		rValuleValidType,
	},
}

var askExpr = parseExpression{
	subTokens: []validTokens{
		askAgentValidType,
		openBracketValidType,
		askBodyValidType,
		closeBracketValidType,
	},
}

var patchExpr = parseExpression{
	subTokens: []validTokens{
		varOrNumberValidType,
		varOrNumberValidType,
	},
}

var sproutExpr = parseExpression{
	subTokens: []validTokens{
		numberValidType,
		openBracketValidType,
		agentInstructionsValidType,
		closeBracketValidType,
	},
}

var createBreedsExpr = parseExpression{
	subTokens: []validTokens{
		varValidType,
		rValuleValidType,
		openBracketValidType,
		agentInstructionsValidType,
		closeBracketValidType,
	},
}

var leftParenthesisExpr = parseExpression{
	subTokens: []validTokens{
		rValuleValidType,
		closeBracketValidType,
	},
}

var setXYExpr = parseExpression{
	subTokens: []validTokens{
		varOrNumberValidType,
		varOrNumberValidType,
	},
}

var turtleExpr = parseExpression{
	subTokens: []validTokens{
		varNumberOrParenthesisValidType,
	},
}
