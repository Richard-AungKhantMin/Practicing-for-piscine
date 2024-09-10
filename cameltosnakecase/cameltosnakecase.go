package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}

func CamelToSnakeCase(s string) string {

	var answer string = string(s[0])

	if len(s) == 0 {
		return ""
	}

	if !IsCamel(s) {

		return s

	}

	for i := 1; i < len(s); i++ {

		if IsCap(s[i]) {

			answer = answer + "_" + string(s[i])

		} else {

			answer = answer + string(s[i])

		}

	}

	return answer
}

func IsCamel(input string) bool {

	count := 0

	for i := 0; i < len(input); i++ {

		if i > 0 && IsCap(input[i]) && IsCap(input[i-1]) {

			return false

		}

		if IsNum(input[i]) {

			return false

		}

		if IsCap(input[i]) {

			count++

		}

		if i == len(input) && count == 0 {

			return false

		}

	}

	last := len(input) - 1

	return !IsCap(input[last])

}

func IsCap(a byte) bool {

	if a >= byte('A') && a <= byte('Z') {
		return true
	}

	return false
}

func IsNum(a byte) bool {

	if a >= byte('0') && a <= byte('9') {
		return true
	}

	return false
}
