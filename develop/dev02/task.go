package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func UnpackString(packed string) (string, error) {
	var builder strings.Builder
	var currentRune rune
	escape := false
	for index, r := range packed {
		if r == '\\' {
			if !escape {
				escape = true
				continue
			}
		}

		if '0' <= r && r <= '9' {
			if index == 0 {
				return "", errors.New("incorrect string")
			}
			if !escape {
				for i := r - '0'; i > 1; i-- {
					builder.WriteRune(currentRune)
				}
				continue
			}
		}

		escape = false
		currentRune = r
		builder.WriteRune(r)
	}
	return builder.String(), nil
}

func main() {
	fmt.Println(UnpackString("a4bc2d5e"))
	fmt.Println(UnpackString("abcd"))
	fmt.Println(UnpackString("45"))
	fmt.Println(UnpackString(""))
	fmt.Println(UnpackString("qwe\\4\\5"))
	fmt.Println(UnpackString("qwe\\45"))
	fmt.Println(UnpackString("qwe\\\\5"))
}
