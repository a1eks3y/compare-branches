package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите строку: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		input = strings.TrimSpace(input)

		compareBrackets(input)
	}
}

var bracketsMap = map[rune]rune{
	'{': '}',
	'[': ']',
	'(': ')',
}

func compareBrackets(brackets string) {
	var stack []rune
	for i, char := range brackets {
		if bracketsMap[char] != 0 {
			stack = append(stack, bracketsMap[char])
		} else {
			length := len(stack)
			lastItem := stack[length-1]
			if lastItem != char {
				fmt.Println("Скобки несбалансированы")
				return
			}

			if i != len(brackets) {
				stack = stack[:length-1]
			}
		}
	}
	fmt.Println("Скобки сбалансированы")
}
