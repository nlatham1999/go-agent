package converter

import (
	"bufio"
	"fmt"
	"os"
)

func Convert(filename string) error {

	// open the file
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	tokens, err := Lexer(lines)
	if err != nil {
		return fmt.Errorf("could not lex file: %v", err)
	}

	tokens = augmentTokens(tokens)

	// for _, token := range tokens {
	// 	fmt.Println(token.tokenType)
	// }

	ast, err := parse(tokens)
	if err != nil {
		return fmt.Errorf("unable to parse file: %v", err)
	}
	ast.print(0)

	return nil
}
