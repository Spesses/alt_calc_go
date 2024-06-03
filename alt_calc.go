package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var strNil string
var operators = map[string]func(string, string) string{
	"+": func(a, b string) string { return a + b },
	"-": func(a, b string) string { return strings.TrimSuffix(a, b) },
	"*": func(a, b string) string { return multiplier(a, b) },
	"/": func(a, b string) string { return divider(a, b) },
}

const (
	GREETINGS        = "Вас приветствует приложение Строковый калькулятор!"
	MENTION          = "Выражение должно быть формата: \n\"Строка\" оператор \"Строка\" \n\"Строка\" оператор число \nЧисло пишется без скобок, иначе воспринимается как строка "
	BV1              = "Первое значение должно быть строкой до 10 символов формата \"Строка\", без пробелов после и перед кавычками"
	BV2              = "Второе значение должно быть строкой до 10 символов формата \"Строка\" или числом от 1 до 10"
	BADOP            = "Калькулятор поддерживает следующие операторы:\"+\", \"-\", \"/\", \"*\". "
	BADOPER          = "Калькулятор умеет выполнять операции сложения строк, вычитания строки из строки, умножения строки на число и деления строки на число: \"a\" + \"b\", \"a\" - \"b\", \"a\" * b, \"a\" / b. "
	BV2T             = "При операции умножения и деления должно использоваться число"
	bkt       string = "\""
)

func multiplier(a, b string) string {
	var result string
	var i, _ = strconv.Atoi(b)
	for j := 0; j < i; j++ {
		result += a
	}
	return result
}

func divider(a, b string) string {
	var result string
	var i, _ = strconv.Atoi(b)
	ln := len(a)
	z := ln / i
	for j := 0; j < z; j++ {
		result += string(a[j])
	}
	return result
}

func executor(input string) (string, []string) {
	var operatorsData string
	var operandsData []string
	var operandsResult []string
	operandsData = strings.Split(input, bkt)
	for idx, val := range operandsData {
		if idx == 1 || idx == 3 {
		} else {
			operandsData[idx] = strings.TrimSpace(val)
		}
		//fmt.Println(len(operandsData[idx]), idx, operandsData[idx])
	}
	if len(operandsData) < 2 {
		panic(MENTION)
	}
	for idx := range operators {
		for _, val := range operandsData[2] {
			if idx == string(val) {
				operatorsData += idx
				operandsData[2] = strings.Trim(operandsData[2], operatorsData)
				operandsData[2] = strings.TrimSpace(operandsData[2])
				var lastCheck, _ = strconv.Atoi(operandsData[2])
				if operatorsData == "*" || operatorsData == "/" {
					if operandsData[2] == strNil {
						panic(BV2T)
					}
				}
				if lastCheck > 10 {
					//fmt.Println(operandsData[2])
					panic(BV2)
				}
			}
		}
	}
	switch {
	case len(operatorsData) != 1:
		panic(BADOP)
	}
	//fmt.Println(operatorsData, "\n", operandsData)
	for idx, val := range operandsData {
		if val != strNil {
			operandsResult = append(operandsResult, operandsData[idx])
		}
	}
	//fmt.Println(operandsResult)
	return operatorsData, operandsResult
}

func validator(operands []string) (string, string) {
	//var data []string = strings.Split(operands[0], bkt)
	var data = operands
	if len(data) != 2 {
		//fmt.Println(true)
	}
	if len(data[0]) > 10 {
		panic(BV1)
	}

	//var val1 = data[1]
	//var val2 string
	//var dataSec = data[3]
	if len(operands[1]) == 1 {
		result, _ := strconv.Atoi(operands[1])
		if result <= 10 {
			if result >= 1 {
				//fmt.Println("Результат: ", result)
			} else {
				panic(BV2)
			}
		} else {
			panic(BV2)
		}
	} else {
		if len(operands[1]) > 10 {
			panic(BV2)
		}
	}

	return operands[0], operands[1]
}

func stringCalc(operator, val1, val2 string) string {
	return operators[operator](val1, val2)
}

func resultEndPrcess(result string) {
	var a = result
	if len(result) > 40 {
		result = strNil
		for i := 0; i < 40; i++ {
			result += string(a[i])
		}
		result += "..."
	}

	fmt.Println("\"" + result + "\"")
}

func main() {
	fmt.Println(GREETINGS)
	reader := bufio.NewReader(os.Stdin)
	for {
		console, err := reader.ReadString('\n')

		if err != nil {
			panic("ReadString Error")
		}

		console = strings.TrimSpace(console)
		operator, operands := executor(console)

		var op1, op2 = validator(operands)
		var result = stringCalc(operator, op1, op2)
		resultEndPrcess(result)
		//fmt.Println("\"" + result + "\"")
	}
}
