package main

import (
	"fmt"
	"strings"
)

func findFirstStringInBracket(str string) string {
	if indexFirstBracketFound := strings.Index(str,"("); indexFirstBracketFound >= 0 {
		if indexClosingBracketFound := strings.Index(str,")"); indexClosingBracketFound > indexFirstBracketFound {
			return str[indexFirstBracketFound + 1:indexClosingBracketFound - 1]
		}
	}
	return ""
}

func main() {
	str := "K(evin Chris)tian"
	fmt.Println(findFirstStringInBracket(str))
}