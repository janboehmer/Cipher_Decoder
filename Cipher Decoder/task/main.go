package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"unicode"
)

func main() {
	var b = 7
	var A int
	var B int
	var S int
	var answer string

	g, p, A := getUserInput()

	B = calcB(g, p, b)

	S = calcS(p, A, b)

	fmt.Printf("B is %d\n", B)

	fmt.Println(caesarCypherEncode("Will you marry me?", S))
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	answer = scanner.Text()
	answer = caesarCypherDecode(answer, S)

	if answer == "Yeah, okay!" {
		fmt.Println(caesarCypherEncode("Great!", S))
	} else if answer == "Let's be friends." {
		fmt.Println(caesarCypherEncode("What a pity!", S))
	}

}

func getUserInput() (int64, int64, int) {
	var g int64
	var p int64
	var A int

	fmt.Scanf("g is %d and p is %d\n", &g, &p)
	fmt.Println("OK")
	fmt.Scanf("A is %d\n", &A)

	return g, p, A
}

func caesarCypherEncode(text string, shift int) string {
	result := ""
	shift = shift % 26

	for _, char := range text {
		if unicode.IsLetter(char) {
			if unicode.IsUpper(char) {
				shifted := (char - 'A' + rune(shift)) % 26
				result += string('A' + shifted)
			} else if unicode.IsLower(char) {
				shifted := (char - 'a' + rune(shift)) % 26
				result += string('a' + shifted)
			}
		} else {
			result += string(char)
		}
	}
	return result
}

func caesarCypherDecode(text string, shift int) string {
	result := ""
	shift = shift % 26

	for _, char := range text {
		if unicode.IsLetter(char) {
			if unicode.IsUpper(char) {
				shifted := (char - 'A' - rune(shift) + 26) % 26
				result += string('A' + shifted)
			} else if unicode.IsLower(char) {
				shifted := (char - 'a' - rune(shift) + 26) % 26
				result += string('a' + shifted)
			}
		} else {
			result += string(char)
		}
	}
	return result
}

func calcB(g int64, p int64, b int) int {
	c := 1
	for i := 0; i < b; i++ {
		c = int(math.Mod(float64(c*int(g)), float64(p)))
	}
	return c
}

func calcS(p int64, A int, b int) int {
	c := 1
	for i := 0; i < b; i++ {
		c = int(math.Mod(float64(c*A), float64(p)))
	}
	return c
}
