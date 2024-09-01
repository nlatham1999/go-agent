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
		}
		for k, v := range colorTokens {
			result[k] = v
		}
		return result
	}(),
}

var askAgentValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		"PATCH":   nil,
		"TURTLES": nil,
		"TURTLE":  nil,
	},
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
		"SPROUT": nil,
		"SET-XY": nil,
		"SET":    nil,
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
