package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

# Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var fFlag = flag.String("f", "", `"fields" - выбрать поля (колонки)`)
var dFlag = flag.String("d", "\t", `"delimiter" - использовать другой разделитель`)
var sFlag = flag.Bool("s", true, `"separated" - только строки с разделителем`)
var columns []int
var fullString bool
var rightMost int

func writeLine(line string) ([]string, bool) {
	var result []string
	lineSeparated := strings.Split(line, *dFlag)
	fields := len(lineSeparated)

	if strings.Count(line, *dFlag) == 0 {
		if *sFlag {
			return nil, true
		} else {
			result = append(result, line)
			return result, false
		}
	}

	for _, index := range columns {
		if index-1 > fields {
			return result, false
		}

		result = append(result, lineSeparated[index-1])
	}
	if fullString {
		for i := rightMost - 1; i < fields; i++ {
			result = append(result, lineSeparated[i])
		}
	}
	return result, false
}

func main() {
	flag.Parse()

	for _, fields := range strings.Split(*fFlag, ",") {
		constrains := strings.Split(fields, "-")
		switch len(constrains) {
		case 1:
			index, err := strconv.Atoi(constrains[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(-1)
			}
			columns = append(columns, index)
		case 2:
			start, err := strconv.Atoi(constrains[0])
			if err != nil {
				start = 0
				os.Exit(-1)
			}
			end, err := strconv.Atoi(constrains[1])
			if err != nil {
				if start == 0 {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(-1)
				}
				rightMost = start
				fullString = true
				end = start
			}
			if end < start {
				os.Exit(-1)
			}
			for i := start; i <= end; i++ {
				columns = append(columns, i)
			}
		default:
			fmt.Fprintln(os.Stderr, "Wrong field declaration")
			os.Exit(-1)
		}
	}
	sort.Ints(columns)
	var builder strings.Builder
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		strs, empty := writeLine(scanner.Text())
		if !empty {
			builder.Reset()
			for _, str := range strs {
				builder.WriteString(str)
				builder.WriteString(" ")
			}
			fmt.Fprintln(os.Stdout, builder.String())
		}
	}

	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, scanner.Err())
		os.Exit(-1)
	}

}
