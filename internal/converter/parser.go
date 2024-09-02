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
		tokenType: tROOT,
	}

	index := 0

	if tokens[0].tokenType != tROOT {
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
				return ast{}, fmt.Errorf(errorMessage(a.token, tokens[*index-1], valids.tokens, tokens[*index]))
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
						return ast{}, fmt.Errorf(errorMessage(a.token, tokens[*index-1], valids.tokens, tokens[*index]))
					}
					found = false
					*index--
				}
			}

		} else if valids.validType == "optional" {
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
				*index--
			}
		} else if valids.validType == "skip" {
			*index++
			if _, ok := valids.tokens[tokens[*index].tokenType]; !ok {
				return ast{}, fmt.Errorf(errorMessage(a.token, tokens[*index-1], valids.tokens, tokens[*index]))
			}
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

func errorMessage(parentToken token, previousToken token, validTokens map[string]interface{}, currentToken token) string {

	tokensArr := []string{}
	for k := range validTokens {
		tokensArr = append(tokensArr, k)
	}

	return fmt.Sprintf(
		"valid next token not found.\n parent token: %s\n previous token: %s\n current token: %s\n line: %d, column: %d\n Valid next tokens: %v",
		parentToken.tokenType, previousToken.tokenType, currentToken.tokenType, currentToken.line+1, currentToken.column, tokensArr,
	)
}

var rootExp = parseExpression{
	subTokens: []validTokens{
		{
			validType: "many",
			tokens:    map[string]interface{}{tBREED: nil, tGLOBALS: nil, tTO: nil, tTO_REPORT: nil},
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
		functionNameValidType,
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
		agentSetValidType,
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
		closeParenthesisValidType,
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

var maxExpr = parseExpression{
	subTokens: []validTokens{
		maxValidType,
	},
}

//this is really for of statements but they start with an open bracket hence the naming
var openBracketOfExpr = parseExpression{
	subTokens: []validTokens{
		rValuleValidType,
		closeBracketOfValidType,
		ofValidType,
		ofRValueValidType,
	},
}

var ifExpr = parseExpression{
	subTokens: []validTokens{
		rValuleValidType,
		openBracketValidType,
		toBodyValidType,
		closeBracketValidType,
	},
}

var ifElseExpr = parseExpression{
	subTokens: []validTokens{
		rValuleValidType,
		openBracketValidType,
		toBodyValidType,
		closeBracketValidType,
		openBracketValidType,
		toBodyValidType,
		closeBracketValidType,
	},
}

var allExpr = parseExpression{
	subTokens: []validTokens{
		agentSetValidType,
		openBracketValidType,
		rValuleValidType,
		closeBracketValidType,
	},
}

var functionCallExpr = parseExpression{
	subTokens: []validTokens{
		functionParametersValidType,
	},
}

var faceXYCallExpr = parseExpression{
	subTokens: []validTokens{
		varOrNumberValidType,
		varOrNumberValidType,
	},
}

var forwardCallExpr = parseExpression{
	subTokens: []validTokens{
		varOrNumberValidType,
	},
}

var faceExpr = parseExpression{
	subTokens: []validTokens{
		turtleAgentValidType,
	},
}

var functionNameExpr = parseExpression{
	subTokens: []validTokens{
		openBracketFunctionParameterValidType,
	},
}

var openBracketFunctionParameterExpr = parseExpression{
	subTokens: []validTokens{
		varValidType,
		closeBracketFunctionParameterValidType,
	},
}

var rightExpr = parseExpression{
	subTokens: []validTokens{
		rValuleValidType,
	},
}

var randomFloatExpr = parseExpression{
	subTokens: []validTokens{
		rValuleValidType,
	},
}

var leftExpr = parseExpression{
	subTokens: []validTokens{
		rValuleValidType,
	},
}

var patchAtExpr = parseExpression{
	subTokens: []validTokens{
		varOrNumberValidType,
		varOrNumberValidType,
	},
}

var reportExpr = parseExpression{
	subTokens: []validTokens{
		rValuleValidType,
	},
}

var randomExpr = parseExpression{
	subTokens: []validTokens{
		varOrNumberValidType,
	},
}
