package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var decoder = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
}

var list = [...]rune{'I', 'V', 'X', 'L', 'C'}

func main() {
	calc()
}

func xor(a, b bool) bool {
	res := !a && b || a && !b
	return res
}

func calc() {
	for {
		fmt.Println("input:")
		input, err := input()
		if err != nil {
			log.Fatal(err)
		}
		splited := strings.Split(input, " ")
		if len(splited) != 3 {
			log.Println("invalid input format")
		}

		var isroman1 bool
		operand1, err := strconv.Atoi(splited[0])

		if err != nil {
			operand1 = fromRoman(splited[0])
			if operand1 == 0 {
				log.Fatal("Недопустимый операнд")
			}
			isroman1 = true
		}

		var isroman2 bool
		operand2, err := strconv.Atoi(splited[2])

		if err != nil {
			operand2 = fromRoman(splited[2])
			if operand2 == 0 {
				log.Fatal("Недопустимый операнд")
			}
			isroman2 = true
		}

		if operand1 > 10 || operand2 > 10 {
			log.Fatal("the number must be <= 10")
		}

		if xor(isroman1, isroman2) {
			log.Fatal("different number systems are used")
		}

		var res int

		switch splited[1] {
		case "+":
			res = operand1 + operand2
		case "-":
			res = operand1 - operand2
		case "*":
			res = operand1 * operand2
		case "/":
			res = operand1 / operand2
		default:
			log.Fatal("non-existent operator")
		}

		fmt.Println("output:")
		if isroman1 {
			if res < 1 {
				fmt.Println("there are no negative numbers and zero in the Roman system")
			} else {
				fmt.Println(toRoman(res))
			}
		} else {
			fmt.Println(res)
		}
	}
}

func toRoman(num int) string {
	i := len(list) - 1
	n := num
	var res string
	for n > 0 {
		for decoder[list[i]] > n {
			i--
		}
		res += string(list[i])
		n -= decoder[list[i]]
	}

	return res
}

func fromRoman(roman string) int {
	if len(roman) == 0 {
		return 0
	}
	first := decoder[rune(roman[0])]
	if len(roman) == 1 {
		return first
	}
	next := decoder[rune(roman[1])]
	if next > first {
		return (next - first) + fromRoman(roman[2:])
	}
	return first + fromRoman(roman[1:])
}

func input() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.TrimSpace(text)

	return text, nil
}
