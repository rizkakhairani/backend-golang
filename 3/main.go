package main

import (
	"fmt"
	"strings"
)

func findFirstStringInBracket(str string) string {
	if len(str) == 0 {
		return ""
	}

	indexFirstBracket := strings.Index(str, "(")
	indexClosingBracket := strings.Index(str, ")")
	countClosingBracket := strings.Count(str, ")")

	var wordsAfterFirstBracket string
	if indexFirstBracket >= 0 && indexClosingBracket > indexFirstBracket {
		wordsAfterFirstBracket = str[indexFirstBracket + 1:indexClosingBracket]
	} else if indexFirstBracket >= 0 && countClosingBracket > 1 {
		for indexClosingBracket < indexFirstBracket {
			if countClosingBracket == 0 {
				break
			}

			str = strings.Replace(str, ")", "", 1)
			indexFirstBracket = strings.Index(str, "(")
			indexClosingBracket = strings.Index(str, ")")
		}

		wordsAfterFirstBracket = str[indexFirstBracket + 1:indexClosingBracket]
	}
		
	return wordsAfterFirstBracket
}

func main() {
	fmt.Println(findFirstStringInBracket("test (first string) (second string)"))	// "first string"
	fmt.Println(findFirstStringInBracket("test) first string (second string)"))		// "second string"
	fmt.Println(findFirstStringInBracket("(test) (first string) (second string)"))	// "test"
	fmt.Println(findFirstStringInBracket("test (first string second string"))		// ""
}

