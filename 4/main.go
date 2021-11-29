package main

import (
	"fmt"
	"sort"
	"strings"
)

func anagram(input []string) [][]string {
	sortString := make(map[string][]string)

	for _, word := range input {
		alphabet := strings.Split(word, "")
		sort.Strings(alphabet)
		key := strings.Join(alphabet, "")
		
		sortString[key] = append(sortString[key], word)
	}

	result := [][]string{}
	for _, arrWord := range sortString {
		result = append(result, arrWord)
	}

	return result
}

func main() {
	fmt.Println(anagram([]string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}))
}