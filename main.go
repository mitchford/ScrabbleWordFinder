package main

import (
	"ScrabbleWordFinder/letterMapping"
	"fmt"
	"strings"
)

var letterMap = letterMapping.MakeMap()
var wordScoreMap = make(map[string]int)

func main() {
	letters := make([]string, 7)
	createPossibleWords(letters)
	for word, _ := range wordScoreMap {
		if isHighestScoringWord(word) {
			fmt.Print("The best scoring word you can make is: {}", word)
		}
	}
}

func createPossibleWords(letters []string) {
	word := letters[2]
	wordScoreMap[word] = getWordScore(word)
}

func isHighestScoringWord(word string) bool {
	highestScore := 0
	highestScoringWord := ""
	for k, v := range wordScoreMap{
		if v > highestScore {
			highestScore = v
			highestScoringWord = k
		}
	}
	if highestScoringWord == word {
		return true
	}
	return false
}

func getWordScore(word string) int {
	total := 0
	letters := strings.Split(word, "")
	for _, l := range letters {
		total += getLetterScore(l)
	}
	return total
}

func getLetterScore(letter string) int {
	return letterMap[letter]
}
