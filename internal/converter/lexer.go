package converter

import (
	"fmt"
	"strconv"
)

func Lexer(lines []string) ([]token, error) {
	tokens := []token{}

	// loop through each line in the file
	for lineNum, line := range lines {

		column := 0

		text := ""
		// fmt.Println(line)
		for column < len(line) {

			// fmt.Println(lineNum, column, line[column])

			// eat whitespace
			for column < len(line) && isWhitespace(line[column]) {
				if text != "" {
					token, err := generateToken(text, lineNum, column-len(text))
					if err != nil {
						return nil, err
					}
					tokens = append(tokens, token)
					text = ""
				}

				column++
				// fmt.Println("eating whitespace")
			}

			// check if we're at the end of the line
			if column >= len(line) {
				// fmt.Println("end of line")
				if text != "" {
					token, err := generateToken(text, lineNum, column-len(text))
					if err != nil {
						return nil, err
					}
					tokens = append(tokens, token)
				}
				continue
			}

			// check for a comment
			if line[column] == ';' {
				if text != "" {
					token, err := generateToken(text, lineNum, column-len(text))
					if err != nil {
						return nil, err
					}
					tokens = append(tokens, token)
				}
				// fmt.Println("comment")
				break
			}

			//check for a string
			if line[column] == '"' {
				if text != "" {
					//generate error because why is there text and then quotes
					// fmt.Println(text)
					return nil, fmt.Errorf("unexpected character %s at line %d, column %d", string(line[column]), lineNum, column)
				}

				text += string(line[column])
				//find the end of the string
				column++
				for column < len(line) && line[column] != '"' {
					text += string(line[column])
					column++
				}

				if column < len(line) {
					text += string(line[column])
					column++
				}

				//make sure the string is closed
				if text[len(text)-1] != '"' {
					return nil, fmt.Errorf("unterminated string at line %d, column %d", lineNum, column)
				}

				//generate the token
				// fmt.Println("gerating token " + text)
				token, err := generateToken(text, lineNum, column-len(text))
				if err != nil {
					return nil, err
				}
				tokens = append(tokens, token)
				text = ""

				// fmt.Println("string")

				continue
			}

			// check for a no space character
			if _, ok := noSpaceTokens[string(line[column])]; ok {
				if text != "" {
					token, err := generateToken(text, lineNum, column-len(text))
					if err != nil {
						return nil, err
					}
					tokens = append(tokens, token)
					text = ""
				}

				// if this character plus the next character is a two character token
				if next := nextCharacter(line, column); next != "" {
					if _, ok := doubleNoSpaceTokens[string(line[column])+next]; ok {
						token, err := generateToken(string(line[column])+next, lineNum, column)
						if err != nil {
							return nil, err
						}
						tokens = append(tokens, token)
						column += 2
						continue
					}
				}

				token, err := generateToken(string(line[column]), lineNum, column)
				if err != nil {
					return nil, err
				}
				tokens = append(tokens, token)
				column++
				// fmt.Println("no space token")
				continue
			}

			// check for a space character
			if isWhitespace(line[column]) {
				if text != "" {
					token, err := generateToken(text, lineNum, column-len(text))
					if err != nil {
						return nil, err
					}
					tokens = append(tokens, token)
					text = ""
				}
				column++
				// fmt.Println("space")
				continue
			}

			// add the character to the text
			text += string(line[column])
			column++
			if column >= len(line) {
				token, err := generateToken(text, lineNum, column-len(text))
				if err != nil {
					return nil, err
				}
				tokens = append(tokens, token)
				text = ""
			}
			// fmt.Println("adding to text")
		}
		lineNum++
	}

	return tokens, nil
}

func generateToken(text string, line int, column int) (token, error) {

	// make sure the token is valid
	if _, ok := notImplementedTokens[text]; ok {
		return token{}, fmt.Errorf("keyword %s not implemented at line %d, column %d", text, line, column)
	}

	// check to see if the token is a keyword
	if tokenType, ok := keywordToTokenType[text]; ok {
		return token{tokenType: tokenType, lexeme: text, line: line, column: column}, nil
	}

	// check to see if the token is a number
	if _, err := strconv.ParseFloat(text, 64); err == nil {
		return token{tokenType: "NUMBER", lexeme: text, line: line, column: column}, nil
	}

	// check to see if the token is a string
	if text[0] == '"' && text[len(text)-1] == '"' {
		return token{tokenType: "STRING", lexeme: text, line: line, column: column}, nil
	}

	//token is an identifier
	return token{tokenType: "VAR", lexeme: text, line: line, column: column}, nil
}

func isWhitespace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n'
}

func nextCharacter(str string, index int) string {
	if index+1 < len(str) {
		return string(str[index+1])
	}
	return ""
}