package converter

var openBracketValidType = validTokens{
	validType: "skip",
	tokens:    map[string]interface{}{"OPEN-BRACKET": nil},
}

var closeBracketValidType = validTokens{
	validType: "skip",
	tokens:    map[string]interface{}{"CLOSE-BRACKET": nil},
}

var varValidType = validTokens{
	validType: "any",
	tokens:    map[string]interface{}{"VAR": nil},
}

var endValidType = validTokens{
	validType: "skip",
	tokens:    map[string]interface{}{"END": nil},
}

var toBodyValidType = validTokens{
	validType: "many",
	tokens: map[string]interface{}{
		"ASK":               nil,
		"CLEAR-ALL":         nil,
		"CREATE-BREEDS":     nil,
		"RESET-TICKS":       nil,
		"SET":               nil,
		"SET-DEFAULT-SHAPE": nil,
		"IF":                nil,
		"STOP":              nil,
		"FACEXY":            nil,
		"FORWARD":           nil,
		"FACE":              nil,
		"TICK":              nil,
		"RIGHT":             nil,
		"LEFT":              nil,
		"IFELSE":            nil,
		"REPORT":            nil,
	},
}

var stringValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		"STRING": nil,
	},
}

var breedsTurtlesLinksValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		"TURTLES": nil,
		"LINKS":   nil,
		"VAR":     nil,
	},
}

var rValuleValidType = validTokens{
	validType: "many",
	tokens: func() map[string]interface{} {
		result := map[string]interface{}{
			"NUMBER":           nil,
			"PLUS":             nil,
			"MIN-PXCOR":        nil,
			"MAX-PXCOR":        nil,
			"MINUS":            nil,
			"STRING":           nil,
			"OPEN-PARENTHESIS": nil,
			"VAR":              nil,
			"MAX":              nil,
			"WHO":              nil,
			"ALL":              nil,
			"EQUAL":            nil,
			"XCOR":             nil,
			"GREATER-EQUAL":    nil,
			"MULTIPLY":         nil,
			"DIVIDE":           nil,
			"POWER":            nil,
			"LESS":             nil,
			"GREATER":          nil,
			"LESS-EQUAL":       nil,
			"NOT-EQUAL":        nil,
			"AND":              nil,
			"RANDOM-FLOAT":     nil,
			"HEADING":          nil,
			"PATCH-AT":         nil,
			"NOBODY":           nil,
			"FUNCTION-CALL":    nil,
			"OPEN-BRACKET-OF":  nil,
			"TURTLE":           nil,
			"RANDOM":           nil,
		}
		for k, v := range colorTokens {
			result[k] = v
		}
		return result
	}(),
}

var varOrNumberValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		"VAR":    nil,
		"NUMBER": nil,
	},
}

var askBodyValidType = validTokens{
	validType: "many",
	tokens: map[string]interface{}{
		"SPROUT":        nil,
		"SET-XY":        nil,
		"SET":           nil,
		"FUNCTION-CALL": nil,
		"IF":            nil,
		"FACE":          nil,
		"FORWARD":       nil,
		"FACE-XY":       nil,
	},
}

var numberValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		"NUMBER": nil,
	},
}

var agentInstructionsValidType = validTokens{
	validType: "many",
	tokens: map[string]interface{}{
		"SET":   nil,
		"STAMP": nil,
		"DIE":   nil,
	},
}

var lValueValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		"VAR":     nil,
		"COLOR":   nil,
		"SHAPE":   nil,
		"SIZE":    nil,
		"HEADING": nil,
	},
}

var closeParenthesisValidType = validTokens{
	validType: "skip",
	tokens:    map[string]interface{}{"CLOSE-PARENTHESIS": nil},
}

var varNumberOrParenthesisValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		"VAR":              nil,
		"NUMBER":           nil,
		"OPEN-PARENTHESIS": nil,
	},
}

var maxValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		"OPEN-BRACKET-OF": nil,
	},
}

var closeBracketOfValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		"CLOSE-BRACKET-OF": nil,
	},
}

var ofValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		"OF": nil,
	},
}

var ofRValueValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		"TURTLES":          nil,
		"OPEN-PARENTHESIS": nil,
	},
}

var agentSetValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		"VAR":     nil,
		"PATCH":   nil,
		"TURTLES": nil,
		"TURTLE":  nil,
	},
}

var functionNameValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		"FUNCTION-NAME": nil,
	},
}

var functionParametersValidType = validTokens{
	validType: "optional",
	tokens: map[string]interface{}{
		"VAR": nil,
	},
}

var turtleAgentValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		"TURTLE": nil,
	},
}

var openBracketFunctionParameterValidType = validTokens{
	validType: "optional",
	tokens:    map[string]interface{}{"OPEN-BRACKET-FUNCTION-PARAMETER": nil},
}

var closeBracketFunctionParameterValidType = validTokens{
	validType: "any",
	tokens:    map[string]interface{}{"CLOSE-BRACKET-FUNCTION-PARAMETER": nil},
}
