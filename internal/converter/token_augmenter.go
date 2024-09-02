// step between the lexer and the parser to create more complex tokens

package converter

func augmentTokens(tokens []token) []token {

	lookingForMaxCloseBracket := false
	//loop backwards to find OF tokens and make the close and open brackets before them into OF tokens
	for i := len(tokens) - 1; i >= 0; i-- {
		if tokens[i].tokenType == tOF {
			lookingForMaxCloseBracket = true
		}
		if lookingForMaxCloseBracket && tokens[i].tokenType == "CLOSE-BRACKET" {
			tokens[i].tokenType = tCLOSE_BRACKET_OF
		}
		if lookingForMaxCloseBracket && tokens[i].tokenType == "OPEN-BRACKET" {
			tokens[i].tokenType = tOPEN_BRACKET_OF
			lookingForMaxCloseBracket = false
		}
	}

	//look for function names
	funcNames := map[string]interface{}{}
	for i := 0; i < len(tokens); i++ {
		if tokens[i].tokenType == "TO" || tokens[i].tokenType == "TO-REPORT" {
			//if the next token is a string, then it is a function name
			if tokens[i+1].tokenType == tVAR {
				tokens[i+1].tokenType = tFUNCTION_NAME
				funcNames[tokens[i+1].lexeme] = nil
			}
		}
	}
	for i := 0; i < len(tokens); i++ {
		if _, ok := funcNames[tokens[i].lexeme]; ok && tokens[i].tokenType == "VAR" {
			tokens[i].tokenType = tFUNCTION_CALL
		}
	}

	//look for function parameter brackets
	lookingForFunctionParameters := false
	for i := 0; i < len(tokens); i++ {
		if tokens[i].tokenType == tFUNCTION_NAME {
			//if the next token is an open bracket, then it is a function parameter
			if tokens[i+1].tokenType == tOPEN_BRACKET {
				tokens[i+1].tokenType = tOPEN_BRACKET_FUNCTION_PARAMETER
				lookingForFunctionParameters = true
			}
		}
		if lookingForFunctionParameters && tokens[i].tokenType == tCLOSE_BRACKET {
			tokens[i].tokenType = tCLOSE_BRACKET_FUNCTION_PARAMETER
			lookingForFunctionParameters = false
		}
	}

	return tokens
}
