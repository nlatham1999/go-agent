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

	//open file and print out tokens
	// create output file
	// outputFile, err := os.Create("output.txt")
	// if err != nil {
	// 	return fmt.Errorf("could not create output file: %v", err)
	// }
	// defer outputFile.Close()

	// // write tokens to output file
	// for _, token := range tokens {
	// 	_, err := outputFile.WriteString(token.tokenType +
	// 		"\t" + token.lexeme + "\n")
	// 	if err != nil {
	// 		return fmt.Errorf("could not write token to output file: %v", err)
	// 	}
	// }

	for _, token := range tokens {
		fmt.Println(token.tokenType + "\t" + token.lexeme)
	}

	return nil
}
