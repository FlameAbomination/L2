package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

# Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

# Дополнительное

# Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var kFlag = flag.Int("k", -1, "указание колонки для сортировки")
var nFlag = flag.Bool("n", false, "сортировать по числовому значению")
var rFlag = flag.Bool("r", false, "сортировать в обратном порядке")
var uFlag = flag.Bool("u", false, "не выводить повторяющиеся строки")
var mFlag = flag.Bool("M", false, "сортировать по названию месяца")
var bFlag = flag.Bool("b", false, "игнорировать хвостовые пробелы")
var cFlag = flag.Bool("c", false, "проверять отсортированы ли данные")
var hFlag = flag.Bool("h", false, "сортировать по числовому значению с учётом суффиксов")

type SortString struct {
	source   string
	keyStr   string
	keyInt   int
	keyFloat float32
}

var monthsMap map[string]int
var suffixMap map[string]int
var re *regexp.Regexp

type SortByColumn []SortString

func (a SortByColumn) Len() int           { return len(a) }
func (a SortByColumn) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByColumn) Less(i, j int) bool { return a[i].keyStr < a[j].keyStr }

type SortByMonth []SortString

func (a SortByMonth) Len() int      { return len(a) }
func (a SortByMonth) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByMonth) Less(i, j int) bool {
	return monthsMap[a[i].keyStr] < monthsMap[a[j].keyStr]
}

type SortByNumber []SortString

func (a SortByNumber) Len() int      { return len(a) }
func (a SortByNumber) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByNumber) Less(i, j int) bool {
	return a[i].keyInt < a[j].keyInt
}

type SortBySuffix []SortString

func (a SortBySuffix) Len() int      { return len(a) }
func (a SortBySuffix) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortBySuffix) Less(i, j int) bool {
	if a[i].keyStr != a[j].keyStr {
		return suffixMap[a[i].keyStr] < suffixMap[a[j].keyStr]
	} else {
		return a[i].keyFloat < a[j].keyFloat
	}
}

func splitColumns(str SortString) []string {
	if *bFlag {
		return strings.Fields(str.source)
	} else {
		return re.FindAllString(str.source, -1)
	}
}

func readLines(filename string, column int) ([]SortString, error) {
	var result []SortString

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var str SortString
		str.source = scanner.Text()
		if column != -1 {
			if column > len(splitColumns(str)) {
				return nil, errors.New("wrong column index")
			}
			str.keyStr = splitColumns(str)[column]
		} else {
			str.keyStr = str.source
		}

		if *nFlag {
			str.keyInt, err = strconv.Atoi(strings.Trim(str.keyStr, " "))
			if err != nil {
				return nil, err
			}
		} else if *hFlag {
			var suffix string
			str.keyStr = splitColumns(str)[column]
			_, err = fmt.Sscanf(str.keyStr, "%f%s", str.keyFloat, &suffix)
			if err != nil {
				return nil, err
			}
		} else if *mFlag {
			str.keyStr = strings.ToUpper(str.keyStr)
		}
		result = append(result, str)
	}

	return result, scanner.Err()
}

func writeLines(filename string, lines []SortString) error {
	var known map[string]struct{}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	if *uFlag {
		known = make(map[string]struct{})
		for _, line := range lines {
			if _, ok := known[line.source]; ok {
				known[line.source] = struct{}{}
				_, err := writer.WriteString(line.source)
				if err != nil {
					return err
				}
				err = writer.Flush()
				if err != nil {
					return err
				}
			}
		}
	} else {
		for _, line := range lines {
			_, err := writer.WriteString(line.source)
			if err != nil {
				return err
			}
			err = writer.Flush()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	flag.Parse()
	re = regexp.MustCompile(`[^ ]+([ ]+|$)`)
	monthsMap = make(map[string]int)
	monthsMap["JAN"] = 1
	monthsMap["FEB"] = 2
	monthsMap["MAR"] = 3
	monthsMap["APR"] = 4
	monthsMap["MAY"] = 5
	monthsMap["JUN"] = 6
	monthsMap["JUL"] = 7
	monthsMap["AUG"] = 8
	monthsMap["SEP"] = 9
	monthsMap["OCT"] = 10
	monthsMap["NOV"] = 11
	monthsMap["DEC"] = 12

	suffixMap = make(map[string]int)
	suffixMap["B"] = 1
	suffixMap["K"] = 2
	suffixMap["M"] = 3
	suffixMap["G"] = 4
	suffixMap["T"] = 5
	suffixMap["P"] = 6
	suffixMap["E"] = 7
	suffixMap["Z"] = 8
	suffixMap["Y"] = 9

	lines, err := readLines("in.txt", *kFlag)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	var data sort.Interface

	if *nFlag {
		data = SortByNumber(lines)
	} else if *mFlag {
		data = SortByMonth(lines)
	} else if *hFlag {
		data = SortBySuffix(lines)
	} else {
		data = SortByColumn(lines)
	}

	if *rFlag {
		data = sort.Reverse(data)
	}

	if *cFlag {
		fmt.Println("Is array sorted? ", sort.IsSorted(data))
	} else {
		sort.Sort(data)
	}
	writeLines("out.txt", lines)
}
