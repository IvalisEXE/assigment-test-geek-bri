package main

import (
	"fmt"
	"strings"
)

// Fungsi untuk memeriksa apakah sebuah angka palindrom atau tidak
func (hp *HighestPalindrom) isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	if s[0] != s[len(s)-1] {
		return false
	}
	return hp.isPalindrome(s[1 : len(s)-1])
}

// Fungsi untuk mengembalikan bilangan palindrom terbesar dengan melakukan paling banyak k penggantian karakter
func (hp *HighestPalindrom) findHighestPalindromeRecursive(s string, k int) string {
	trimInput := strings.TrimSpace(s)
	fieldsInput := strings.Fields(trimInput)
	s = strings.Join(fieldsInput, " ")

	runes := []rune(s)

	// Looping melalui string dan mengganti karakter yang tidak sama di awal dan akhir
	var compareDifferentStartAndEnd func([]rune, int, int, int) int
	start := 0
	end := len(s) - 1
	compareDifferentStartAndEnd = func(runes []rune, start, end, k int) int {
		if start < end {
			if runes[start] != runes[end] {
				// Pilih karakter yang lebih besar untuk diganti dengan karakter yang sama di posisi yang berlawanan
				if runes[start] > runes[end] {
					runes[end] = runes[start]
				} else {
					runes[start] = runes[end]
				}
				k-- // Kurangi jumlah penggantian yang tersedia
			}
			start++
			end--
			return compareDifferentStartAndEnd(runes, start, end, k)
		} else {
			return k
		}
	}

	k = compareDifferentStartAndEnd(runes, start, end, k)

	start = 0
	end = len(s) - 1
	// Jika masih tersisa penggantian, coba ganti karakter di tengah string menjadi '9' untuk membuat palindrom terbesar
	var compareRemainingReplacment func([]rune, int, int, int, string)
	compareRemainingReplacment = func(runes []rune, start2, end2, k int, s string) {
		if k > 0 && start <= end {
			if start == end {
				runes[start] = '9'
				compareRemainingReplacment(runes, start, end, k, s)
			}
			if runes[start] != '9' {
				// Ganti karakter di posisi yang sama dengan '9'
				runes[start] = '9'
				runes[end] = '9'
				k-- // Kurangi jumlah penggantian yang tersedia
			}
			start++
			end--

			compareRemainingReplacment(runes, start2, end2, k, s)
		}
	}

	compareRemainingReplacment(runes, start, end, k, s)

	if !hp.isPalindrome(string(runes)) {
		return "-1"
	}
	return string(runes)
}

type HighestPalindrom struct{}

func NewHighestPalindrom() *HighestPalindrom {
	return &HighestPalindrom{}
}

func main() {
	var input string
	var k int

	fmt.Println("Masukkan string:")
	fmt.Scanln(&input)
	fmt.Println("Masukkan nilai k:")
	fmt.Scanln(&k)
	fmt.Println()

	hp := NewHighestPalindrom()
	result := hp.findHighestPalindromeRecursive(input, k)

	fmt.Println("Output:", result)
}
