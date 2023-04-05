package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func getKey(str string) string {
	runeSlice := []rune(str)
	sort.Slice(runeSlice, func(i, j int) bool {
		return runeSlice[i] < runeSlice[j]
	})
	return string(runeSlice)
}

func getDictionary(words []string) map[string][]string {
	tempDictionary := make(map[string][]string)
	dictionary := make(map[string][]string)

	for _, word := range words {
		word = strings.ToLower(word)
		tempDictionary[getKey(word)] = append(tempDictionary[getKey(word)], word)
	}

	for _, value := range tempDictionary {
		if len(value) > 1 {
			valueSorted := value[1:]
			sort.Strings(valueSorted)
			dictionary[value[0]] = valueSorted
		}
	}

	return dictionary
}

func main() {
	dict := []string{
		"пятак", "пятка", "тяпка",
		"листок", "слиток", "столик",
	}
	fmt.Print(getDictionary(dict))
}
