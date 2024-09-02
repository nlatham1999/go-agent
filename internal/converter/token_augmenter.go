// step between the lexer and the parser to create more complex tokens

package converter

func augmentTokens(tokens []token) []token {

	lookingForMaxCloseBracket := false
	//loop backwards to find OF tokens and make the close and open brackets before them into OF tokens
	for i := len(tokens) - 1; i >= 0; i-- {
		if tokens[i].tokenType == "OF" {
			lookingForMaxCloseBracket = true
		}
		if lookingForMaxCloseBracket && tokens[i].tokenType == "CLOSE-BRACKET" {
			tokens[i].tokenType = "CLOSE-BRACKET-OF"
		}
		if lookingForMaxCloseBracket && tokens[i].tokenType == "OPEN-BRACKET" {
			tokens[i].tokenType = "OPEN-BRACKET-OF"
			lookingForMaxCloseBracket = false
		}
	}

	//look for function names
	funcNames := map[string]interface{}{}
	for i := 0; i < len(tokens); i++ {
		if tokens[i].tokenType == "TO" || tokens[i].tokenType == "TO-REPORT" {
			//if the next token is a string, then it is a function name
			if tokens[i+1].tokenType == "VAR" {
				tokens[i+1].tokenType = "FUNCTION-NAME"
				funcNames[tokens[i+1].lexeme] = nil
			}
		}
	}
	for i := 0; i < len(tokens); i++ {
		if _, ok := funcNames[tokens[i].lexeme]; ok && tokens[i].tokenType == "VAR" {
			tokens[i].tokenType = "FUNCTION-CALL"
		}
	}

	//look for function parameter brackets
	lookingForFunctionParameters := false
	for i := 0; i < len(tokens); i++ {
		if tokens[i].tokenType == "FUNCTION-NAME" {
			//if the next token is an open bracket, then it is a function parameter
			if tokens[i+1].tokenType == "OPEN-BRACKET" {
				tokens[i+1].tokenType = "OPEN-BRACKET-FUNCTION-PARAMETER"
				lookingForFunctionParameters = true
			}
		}
		if lookingForFunctionParameters && tokens[i].tokenType == "CLOSE-BRACKET" {
			tokens[i].tokenType = "CLOSE-BRACKET-FUNCTION-PARAMETER"
			lookingForFunctionParameters = false
		}
	}

	return tokens
}
