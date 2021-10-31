package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortWord(word string) string {
	str := strings.Split(word, "")
    sort.Strings(str)
    return strings.Join(str, "")
}

func main() {
	listOfString := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}

	anagrams := make(map[string][]string)
	for _, v := range listOfString {
		sorted := sortWord(v)
		anagrams[sorted] = append(anagrams[sorted], v)
	}
	
	for _, str := range anagrams {
		fmt.Println(str)
	}
}