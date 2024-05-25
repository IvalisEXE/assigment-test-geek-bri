package main

import (
	"fmt"
	"strings"
)

const (
	ALPHABET         = "abcdefghijklmnopqrstuvwxyz"
	SUCCESS_RESULT_1 = "Yes"
	FAILED_RESULT_1  = "No"
)

type WeightedStrings struct {
	weights map[rune]int
}

// Fungsi untuk menghitung total bobot dari sebuah string
func (ws *WeightedStrings) doValidationQueriesOnChar(input string, queries []int) []string {
	trimInput := strings.TrimSpace(input)
	fieldsInput := strings.Fields(trimInput)
	input = strings.Join(fieldsInput, " ")

	var result []string
	var hasQuery bool

	for queriesKey := range queries {
		repeatedSubstring := ws.findRepeatedSubstrings(input)
		for repeatedSubstringKey := range repeatedSubstring {
			hasQuery = false
			char := repeatedSubstring[repeatedSubstringKey]
			charWithTotalWeight := ws.findTotalWeight(char)
			if foundResult := SUCCESS_RESULT_1; queries[queriesKey] == charWithTotalWeight {
				result = append(result, foundResult)
				hasQuery = true
				break
			}
		}

		if noResult := FAILED_RESULT_1; !hasQuery {
			result = append(result, noResult)
		}

	}
	return result
}

// Fungsi untuk mendapatkan substring dari perulangan pertama hingga n
func (ws *WeightedStrings) findRepeatedSubstrings(input string) []string {
	substrings := make([]string, 0)

	// Indeks awal karakter dalam string
	startIndex := 0

	// Iterasi melalui string
	for i := 0; i < len(input)-1; i++ {
		// Jika karakter saat ini berbeda dengan karakter berikutnya, berarti substring berulang selesai
		if input[i] != input[i+1] {
			// Tambahkan substring yang berulang secara berurutan
			for j := startIndex; j <= i; j++ {
				substring := input[startIndex : j+1]
				substrings = append(substrings, substring)
			}
			// Perbarui indeks awal untuk karakter berikutnya
			startIndex = i + 1
		}
	}

	// Tambahkan substring terakhir jika ada
	for j := startIndex; j < len(input); j++ {
		substring := input[startIndex : j+1]
		substrings = append(substrings, substring)
	}

	return substrings
}

// Fungsi untuk menghitung total bobot dari sebuah string
func (ws *WeightedStrings) findTotalWeight(s string) int {
	total := 0
	for _, char := range s {
		total += ws.weights[char]
	}
	return total
}

// Fungsi untuk membuat WeightedStrings baru dengan bobot yang ditentukan
func NewWeightedStrings() *WeightedStrings {
	weights := make(map[rune]int)
	for i, char := range ALPHABET {
		weights[char] = i + 1
	}
	return &WeightedStrings{weights}
}

func main() {
	ws := NewWeightedStrings()

	var targetChar string
	fmt.Print("Masukkan target character: ")
	fmt.Scanln(&targetChar)
	fmt.Println()

	fmt.Println("Maksimum queries yang di perbolehkan adalah 4 number")
	lenQueries := 4
	queries := make([]int, lenQueries)
	for i := 0; i < lenQueries; i++ {
		fmt.Printf("Masukkan target query ke-%d: ", i+1)
		fmt.Scanln(&queries[i])
	}

	fmt.Println()
	fmt.Println("Berikut hasil dari Weighted Strings ada dibawah ini")
	fmt.Println(ws.doValidationQueriesOnChar(targetChar, queries))
}
