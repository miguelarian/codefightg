package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

const Alphabet = "abcdefghijklmnopqrstuvwxyz"

type ResultItem struct {
	letter string
	count  int
	source int
}

type ByLexicographic []ResultItem

func (items ByLexicographic) Len() int      { return len(items) }
func (items ByLexicographic) Swap(i, j int) { items[i], items[j] = items[j], items[i] }
func (items ByLexicographic) Less(i, j int) bool {

	if items[i].count > items[j].count {
		return true
	} else if items[i].count < items[j].count {
		return false
	} else {
		return items[i].toString() < items[j].toString()
	}
}

func (item *ResultItem) toString() string {
	var result string
	if item.source == 0 {
		result = "=:"
	} else {
		result = strconv.Itoa(item.source) + ":"
	}
	result += strings.Repeat(string(item.letter), item.count)
	return result
}

func getLowercaseLettersCount(input string) map[string]int {
	var result map[string]int
	result = make(map[string]int)

	for _, letter := range input {
		if unicode.IsLower(letter) {
			key := string(letter)
			count, containsKey := result[key]
			result[key] = 1
			if containsKey {
				result[key] += count
			}
		}
	}

	return result
}

func buildResult(allItems []ResultItem) string {

	sort.Sort(ByLexicographic(allItems))

	var result []string
	for _, item := range allItems {
		result = append(result, item.toString())
	}

	return strings.Join(result, "/")
}

func Mix(arg1 string, arg2 string) string {
	var ocurrences map[string]ResultItem
	ocurrences = make(map[string]ResultItem)

	// Get lowercase letters from inputs
	lowercaseLetters1 := getLowercaseLettersCount(arg1)
	lowercaseLetters2 := getLowercaseLettersCount(arg2)

	var allItems []ResultItem
	for _, letter := range Alphabet {
		// Foreach letter, check if present in S1 or S2
		letterKey := string(letter)
		countInS1, isPresentInS1 := lowercaseLetters1[letterKey]
		countInS2, isPresentInS2 := lowercaseLetters2[letterKey]

		// Skip iteration if letter not present in both or less than 2
		if !isPresentInS1 && !isPresentInS2 {
			continue
		}

		// Add to S1
		if isPresentInS1 && !isPresentInS2 && countInS1 > 1 {
			ocurrences[letterKey] = ResultItem{source: 1, letter: letterKey, count: countInS1}
		}

		// Add to S2
		if !isPresentInS1 && isPresentInS2 && countInS2 > 1 {
			ocurrences[letterKey] = ResultItem{source: 2, letter: letterKey, count: countInS2}
		}

		// Draw - check both
		if isPresentInS1 && isPresentInS2 {
			if countInS1 == countInS2 && countInS1 > 1 {
				ocurrences[letterKey] = ResultItem{source: 0, letter: letterKey, count: countInS1}
			} else if countInS1 > countInS2 && countInS1 > 1 {
				ocurrences[letterKey] = ResultItem{source: 1, letter: letterKey, count: countInS1}
			} else if countInS2 > countInS1 && countInS2 > 1 {
				ocurrences[letterKey] = ResultItem{source: 2, letter: letterKey, count: countInS2}
			}
		}

		item, letterAdded := ocurrences[letterKey]
		if letterAdded {
			allItems = append(allItems, item)
		}
	}

	return buildResult(allItems)
}

func main() {
	fmt.Println(Mix("Are they here", "yes, they are here"))
	fmt.Println("2:eeeee/2:yy/=:hh/=:rr")
}
