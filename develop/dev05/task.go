package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
var re *regexp.Regexp

var AFlag = flag.Int("A", 0, `"after" печатать +N строк после совпадения`)
var BFlag = flag.Int("B", 0, `"before" печатать +N строк до совпадения`)
var CFlag = flag.Int("C", 0, `"context" (A+B) печатать ±N строк вокруг совпадения`)
var cFlag = flag.Bool("c", false, `"count" (количество строк)`)
var iFlag = flag.Bool("i", false, `"ignore-case" (игнорировать регистр)`)
var vFlag = flag.Bool("v", false, `"invert" (вместо совпадения, исключать)`)
var FFlag = flag.Bool("F", false, `"fixed", точное совпадение со строкой, не паттерн`)
var nFlag = flag.Bool("n", false, `"line num", печатать номер строки`)

var before int
var after int

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func readLines(filename string) ([]string, error) {
	var result []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, scanner.Err()
}

func getLines(lines []string) (int, []bool, error) {
	count := 0
	mask := make([]bool, len(lines))
	for i := 0; i < len(mask); i++ {
		mask[i] = false
	}
	for index, line := range lines {
		if mask[index] {
			continue
		}
		if re.FindString(line) != "" {
			start := getMax(0, index-before)
			for i := start; i <= index+after; i++ {
				if !mask[i] {
					mask[i] = true
					count++
				}
			}
		}
	}
	return count, mask, nil
}

func main() {
	var pattern string

	flag.Parse()

	after = getMax(*CFlag, *AFlag)
	before = getMax(*CFlag, *BFlag)

	if *FFlag {
		pattern = "^" + flag.Args()[0] + "$"
	} else {
		pattern = flag.Args()[0]
	}

	if *iFlag {
		pattern = `(?i)` + pattern
	}

	re = regexp.MustCompile(pattern)

	lines, err := readLines("in.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	count, mask, err := getLines(lines)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	if *cFlag {
		fmt.Print(count)
		os.Exit(0)
	}

	for i := 0; i < len(mask); i++ {
		if mask[i] == (*vFlag) {
			continue
		}
		if *nFlag {
			fmt.Printf("%4d: ", i+1)
		}
		fmt.Println(lines[i])
	}

}
