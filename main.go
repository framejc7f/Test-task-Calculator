package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var decoder = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
}

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
			operand1, err = fromRoman(splited[0])
			if err != nil {
				log.Fatal(err)
			}

			if operand1 == 0 {
				log.Fatal("invalid operand")
			}
			isroman1 = true
		}

		var isroman2 bool
		operand2, err := strconv.Atoi(splited[2])

		if err != nil {
			operand2, err = fromRoman(splited[2])

			if err != nil {
				log.Fatal(err)
			}

			if operand2 == 0 {
				log.Fatal("invalid operand")
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

func toRoman(arabic int) string {
	roman := ""
	values := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < len(values); i++ {
		for arabic >= values[i] {
			arabic -= values[i]
			roman += symbols[i]
		}
	}

	return roman
}

func fromRoman(roman string) (int, error) {
	arabic := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		currentSymbol := string(roman[i])
		currentValue, ok := decoder[currentSymbol]

		if !ok {
			return 0, errors.New("invalid roman numeral character")
		}

		if currentValue < prevValue {
			arabic -= currentValue
		} else {
			arabic += currentValue
		}

		prevValue = currentValue
	}

	if toRoman(arabic) != roman {
		return 0, errors.New("incorrect entry of the Roman number")
	}

	return arabic, nil
}

func input() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.TrimSpace(text)

	if text == "" {
		return "", errors.New("empty string")
	}

	return text, nil
}