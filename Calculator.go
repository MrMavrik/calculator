package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Roman bool = false

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Пример ввода: 3 + 4")
	//тело цикла
	for {
		fmt.Println("Введите значение ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		outTextArray := strings.Split(text, " ")

		if len(outTextArray) != 3 {
			err := errors.New("Введено не корректное кол-во значений!")
			fmt.Println(err)
			continue // return для завершения программы
		}
		symbol := outTextArray[1]

		n1, n2, err := checkVariablesInt(outTextArray)
		if err != nil {
			fmt.Println(err)
		} else {
			sum, err := calculateInt(n1, n2, symbol)
			if err != nil {
				fmt.Println(err)
				continue // return для завершения программы
			}
			if Roman == true {
				if sum < 1 {
					fmt.Printf("Исключение, результат вычислений римских цифр не может быть отрицательным или равен 0, результат равен %d", sum)
				}
				fmt.Println(integerToRoman(sum))
			} else {
				fmt.Println(sum)
			}
		}
	}
}

func checkVariablesInt(c []string) (int, int, error) {
	//fmt.Println(c)
	isNum1, isNum2 := true, true

	num1, err := strconv.Atoi(c[0]) // перевод в int
	if err != nil {
		isNum1 = false
	}
	num2, err := strconv.Atoi(c[2])
	if err != nil {
		isNum2 = false
	}
	if isNum1 || isNum2 {
		Roman = false
	} else {
		Roman = true
	}

	if isNum1 != isNum2 {
		return 0, 0, errors.New("Вывод ошибки, так как используются одновременно разные системы счисления.")
	}

	if isNum1 == true && isNum2 == true {
		if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
			return 0, 0, errors.New("Вывод ошибки, так как арабские цифры должны быть от 1 до 10 включительно.")
		}
	}

	if isNum1 == false && isNum2 == false { //ввели римские цифры или float
		var x = strings.Split(c[0], "")
		var y = strings.Split(c[2], "")
		r1, err := romanNumerals(x)
		r2, err2 := romanNumerals(y)

		if err != nil {
			return 0, 0, err
		}
		if err2 != nil {
			return r1, r2, err2
		}
		num1, num2 = r1, r2
	}
	return num1, num2, nil
}

func romanNumerals(r []string) (int, error) {
	roman := map[string](int){"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	var pred, num int
	for i := len(r) - 1; i >= 0; i-- {
		_, ok := roman[r[i]]
		if !ok {
			return 0, errors.New("Вывод ошибки, так как в римской системе нет такого значение или значение не int.")
		}
		x := roman[r[i]]
		if x < pred {
			num -= x
		} else {
			num += x
			pred = x
		}
		pred = x
	}
	return num, nil
}

func calculateInt(a, b int, symbol string) (int, error) {

	switch symbol {

	case "+":
		return a + b, nil

	case "-":
		return a - b, nil

	case "*":
		return a * b, nil

	case "/":
		return a / b, nil

	default:
		return 0, errors.New("Вывод ошибки, так как введён не правильный арифметический символ.(+, -, /, *)")
	}
}

func integerToRoman(number int) string {

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}
	return roman.String()
}
