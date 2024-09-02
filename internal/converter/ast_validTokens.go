package converter

var openBracketValidType = validTokens{
	validType: "skip",
	tokens:    map[string]interface{}{tOPEN_BRACKET: nil},
}

var closeBracketValidType = validTokens{
	validType: "skip",
	tokens:    map[string]interface{}{tCLOSE_BRACKET: nil},
}

var varValidType = validTokens{
	validType: "any",
	tokens:    map[string]interface{}{tVAR: nil},
}

var endValidType = validTokens{
	validType: "skip",
	tokens:    map[string]interface{}{"END": nil},
}

var toBodyValidType = validTokens{
	validType: "many",
	tokens: map[string]interface{}{
		tASK:               nil,
		tCLEAR_ALL:         nil,
		tCREATE_BREEDS:     nil,
		tRESET_TICKS:       nil,
		tSET:               nil,
		tSET_DEFAULT_SHAPE: nil,
		tIF:                nil,
		tSTOP:              nil,
		tFACEXY:            nil,
		tFORWARD:           nil,
		tFACE:              nil,
		tTICK:              nil,
		tRIGHT:             nil,
		tLEFT:              nil,
		tIFELSE:            nil,
		tREPORT:            nil,
	},
}

var stringValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		tSTRING: nil,
	},
}

var breedsTurtlesLinksValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		tTURTLES: nil,
		tLINKS:   nil,
		tVAR:     nil,
	},
}

var rValuleValidType = validTokens{
	validType: "many",
	tokens: func() map[string]interface{} {
		result := map[string]interface{}{
			tNUMBER:           nil,
			tPLUS:             nil,
			tMIN_PXCOR:        nil,
			tMAX_PXCOR:        nil,
			tMINUS:            nil,
			tSTRING:           nil,
			tOPEN_PARENTHESIS: nil,
			tVAR:              nil,
			tMAX:              nil,
			tWHO:              nil,
			tALL:              nil,
			tEQUAL:            nil,
			tXCOR:             nil,
			tGREATER_EQUAL:    nil,
			tMULTIPLY:         nil,
			tDIVIDE:           nil,
			tPOWER:            nil,
			tLESS:             nil,
			tGREATER:          nil,
			tLESS_EQUAL:       nil,
			tNOT_EQUAL:        nil,
			tAND:              nil,
			tRANDOM_FLOAT:     nil,
			tHEADING:          nil,
			tPATCH_AT:         nil,
			tNOBODY:           nil,
			tFUNCTION_CALL:    nil,
			tOPEN_BRACKET_OF:  nil,
			tTURTLE:           nil,
			tRANDOM:           nil,
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
		tVAR:    nil,
		tNUMBER: nil,
	},
}

var askBodyValidType = validTokens{
	validType: "many",
	tokens: map[string]interface{}{
		tSPROUT:        nil,
		tSET_XY:        nil,
		tSET:           nil,
		tFUNCTION_CALL: nil,
		tIF:            nil,
		tFACE:          nil,
		tFORWARD:       nil,
		tFACEXY:        nil,
	},
}

var numberValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		tNUMBER: nil,
	},
}

var agentInstructionsValidType = validTokens{
	validType: "many",
	tokens: map[string]interface{}{
		tSET:   nil,
		tSTAMP: nil,
		tDIE:   nil,
	},
}

var lValueValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		tVAR:     nil,
		tCOLOR:   nil,
		tSHAPE:   nil,
		tSIZE:    nil,
		tHEADING: nil,
	},
}

var closeParenthesisValidType = validTokens{
	validType: "skip",
	tokens:    map[string]interface{}{tCLOSE_PARENTHESIS: nil},
}

var varNumberOrParenthesisValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		tVAR:              nil,
		tNUMBER:           nil,
		tOPEN_PARENTHESIS: nil,
	},
}

var maxValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		tOPEN_BRACKET_OF: nil,
	},
}

var closeBracketOfValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		tCLOSE_BRACKET_OF: nil,
	},
}

var ofValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		tOF: nil,
	},
}

var ofRValueValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		tTURTLES:          nil,
		tOPEN_PARENTHESIS: nil,
	},
}

var agentSetValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		tVAR:     nil,
		tPATCH:   nil,
		tTURTLES: nil,
		tTURTLE:  nil,
	},
}

var functionNameValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		tFUNCTION_NAME: nil,
	},
}

var functionParametersValidType = validTokens{
	validType: "optional",
	tokens: map[string]interface{}{
		tVAR: nil,
	},
}

var turtleAgentValidType = validTokens{
	validType: "any",
	tokens: map[string]interface{}{
		tTURTLE: nil,
	},
}

var openBracketFunctionParameterValidType = validTokens{
	validType: "optional",
	tokens:    map[string]interface{}{tOPEN_BRACKET_FUNCTION_PARAMETER: nil},
}

var closeBracketFunctionParameterValidType = validTokens{
	validType: "any",
	tokens:    map[string]interface{}{tCLOSE_BRACKET_FUNCTION_PARAMETER: nil},
}
