package main

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

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const escapeSymbol = '\\' // Обратная косая черта

// StringUnpack распаковывает строку с повторяющимися символами.
func StringUnpack(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}
	if unicode.IsDigit(rune(s[0])) {
		return "", errors.New("некорректная строка: строка не может начинаться с цифры")
	}

	var builder strings.Builder
	runes := []rune(s)
	i := 0

	for i < len(runes) {
		var runeToPrint rune = runes[i]

		if runeToPrint == escapeSymbol { // Если встретилась escape-последовательность
			i++
			if i < len(runes) {
				runeToPrint = runes[i] // Берем следующий символ
				// Если следующий символ - это цифра, то добавляем его как есть
				if unicode.IsDigit(runeToPrint) {
					builder.WriteRune(runeToPrint)
				} else {
					// В противном случае добавляем символ без изменений
					builder.WriteRune(runeToPrint)
				}
			} else {
				return "", errors.New("некорректная строка: незавершенная escape-последовательность")
			}
		} else {
			if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
				num, err := strconv.Atoi(string(runes[i+1]))
				if err != nil {
					return "", errors.New("не удалось преобразовать цифру в строку")
				}
				for j := 0; j < num; j++ {
					builder.WriteRune(runeToPrint)
				}
				i++ // Пропускаем цифру
			} else {
				builder.WriteRune(runeToPrint)
			}
		}
		i++
	}

	return builder.String(), nil
}

func main() {
	fmt.Println(StringUnpack("a4bc2d5e"))
	fmt.Println(StringUnpack("abcd"))
	fmt.Println(StringUnpack("45"))
	fmt.Println(StringUnpack(""))
}
