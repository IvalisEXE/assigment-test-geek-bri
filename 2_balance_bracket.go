package main

import (
	"fmt"
	"strings"
)

const (
	SUCCESS_RESULT_2 = "YES"
	FAILED_RESULT_2  = "N0"
)

var (
	BracketOpen1  rune = '('
	BracketClose1 rune = ')'
	BracketOpen2  rune = '['
	BracketClose2 rune = ']'
	BracketOpen3  rune = '{'
	BracketClose3 rune = '}'
)

var allowedBracketCharacters = map[rune]bool{
	BracketOpen1:  true,
	BracketClose1: true,
	BracketOpen2:  true,
	BracketClose2: true,
	BracketOpen3:  true,
	BracketClose3: true,
}

func (bb *BalanceBracket) isBalancedBracket(input string) string {
	trimInput := strings.TrimSpace(input)
	fieldsInput := strings.Fields(trimInput)
	input = strings.Join(fieldsInput, " ")

	if inputNok := input == ""; inputNok {
		fmt.Println("Penjelasan -> input tidak boleh kosong")
		return FAILED_RESULT_2
	}

	inputLen := len(input)
	if inputBalance := inputLen%2 == 0; !inputBalance {
		fmt.Println("Penjelasan -> tidak seimbang untuk karakter yang diapit")
		return FAILED_RESULT_2
	}

	for i := 0; i < inputLen/2; i++ {
		start := input[i]
		end := input[inputLen-1-i]
		if allowedBracketCharacters := bb.findAllowedBracketCharacters(rune(start)) &&
			bb.findAllowedBracketCharacters(rune(end)); !allowedBracketCharacters {

			fmt.Println("Penjelasan -> Input hanya boleh { [ ( ) ] }")
			return FAILED_RESULT_2
		}

		bb.findBalanceBracket(rune(start), rune(end))

		bb.findCountAllBucketOpenAndClose(rune(start))
		bb.findCountAllBucketOpenAndClose(rune(end))
	}

	if bb.isBracketBalance {
		fmt.Println("Penjelasan -> Setiap bracket seimbang, antara bracket buka dan bracket tutup.")
		return SUCCESS_RESULT_2
	}

	if bb.findIrregularBalanceBracket() {
		fmt.Println("Penjelasan -> Setiap bracket seimbang, antara bracket buka dan bracket tutup, meskipun struktur bracket tidak beraturan.")
		return SUCCESS_RESULT_2
	}

	return FAILED_RESULT_2 // Unknown rule
}

func (bb *BalanceBracket) findAllowedBracketCharacters(bracket rune) bool {
	return allowedBracketCharacters[bracket] && bracket != ' '
}

func (bb *BalanceBracket) findCountAllBucketOpenAndClose(bracket rune) {
	switch bracket {
	case BracketOpen1:
		bb.countBracketOpen1++
	case BracketOpen2:
		bb.countBracketOpen2++
	case BracketOpen3:
		bb.countBracketOpen3++
	case BracketClose1:
		bb.countBracketClose1++
	case BracketClose2:
		bb.countBracketClose2++
	case BracketClose3:
		bb.countBracketClose3++
	default:
	}
}

func (bb *BalanceBracket) findBalanceBracket(start, end rune) bool {
	balanceBracket1 := start == BracketOpen1 && end == BracketClose1
	balanceBracket2 := start == BracketOpen2 && end == BracketClose2
	balanceBracket3 := start == BracketOpen3 && end == BracketClose3
	bb.isBracketBalance = balanceBracket1 || balanceBracket2 || balanceBracket3
	return bb.isBracketBalance
}

func (bb *BalanceBracket) findIrregularBalanceBracket() bool {
	bb.isIrregularBracketBalance = bb.countBracketOpen1 == bb.countBracketClose1 &&
		bb.countBracketOpen2 == bb.countBracketClose2 &&
		bb.countBracketOpen3 == bb.countBracketClose3
	return bb.isIrregularBracketBalance
}

type BalanceBracket struct {
	isBracketBalance          bool // Setiap bracket seimbang, antara bracket buka dan bracket tutup.
	isIrregularBracketBalance bool // Setiap bracket seimbang, antara bracket buka dan bracket tutup, meskipun struktur bracket tidak beraturan.
	countBracketOpen1         int
	countBracketClose1        int
	countBracketOpen2         int
	countBracketClose2        int
	countBracketOpen3         int
	countBracketClose3        int
}

func NewBalanceBracket() *BalanceBracket {
	return &BalanceBracket{}
}

func main() {
	var bracketChar string
	fmt.Print("Masukkan bracket character: ")
	fmt.Scanln(&bracketChar)
	fmt.Println()

	result := NewBalanceBracket()
	fmt.Println("hasil -> ", result.isBalancedBracket(bracketChar))
}
