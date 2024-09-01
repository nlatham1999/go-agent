// step between the lexer and the parser to create more complex tokens

package converter

func condenseTokens(tokens []token) []token {

	newTokens := []token{}

	lookingForMaxCloseBracket := false
	for i := 0; i < len(tokens); i++ {

		// if the token is max and the next token is an open bracket, make the brackets open and close reporters
		if tokens[i].tokenType == "MAX" && tokens[i+1].tokenType == "OPEN-BRACKET" {
			tokens[i+1].tokenType = "MAX-OPEN-BRACKET"
			lookingForMaxCloseBracket = true
		}

		if lookingForMaxCloseBracket && tokens[i].tokenType == "CLOSE-BRACKET" {
			tokens[i].tokenType = "MAX-CLOSE-BRACKET"
			lookingForMaxCloseBracket = false
		}

		newTokens = append(newTokens, tokens[i])
	}

	return tokens
}
