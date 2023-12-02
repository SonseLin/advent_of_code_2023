package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getTwoNumber(data string) (int, int) {
	number_string := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	digitRegex := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|\d)`)
	matches := digitRegex.FindAllString(data, -1)
	var data_int []int
	for _, v := range matches {
		if num, ok := number_string[v]; ok {
			data_int = append(data_int, num)
		} else if number_string[v] == 0 {
			n, _ := strconv.Atoi(v)
			data_int = append(data_int, n)
		}
	}
	digitRegexReversed := regexp.MustCompile(`(eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|\d)`)
	runes := []rune(data)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	reversed := string(runes)
	reverse := digitRegexReversed.FindAllString(reversed, -1)
	for _, v := range reverse {
		n, err := strconv.Atoi(v)
		if err == nil {
			data_int = append(data_int, n)
		} else {
			rune_word := []rune(v)
			for i, j := 0, len(rune_word)-1; i < j; i, j = i+1, j-1 {
				rune_word[i], rune_word[j] = rune_word[j], rune_word[i]
			}
			rev := string(rune_word)
			data_int = append(data_int, number_string[rev])
		}
	}
	if len(data_int) > 2 {
		return data_int[0], data_int[len(data_int)/2]
	} else {
		return data_int[0], data_int[len(data_int)-2]
	}
}

func main() {
	dat, err := os.Open("input.txt")
	check(err)
	defer dat.Close()
	file_line := bufio.NewScanner(dat)
	sum := 0
	for file_line.Scan() {
		x, y := getTwoNumber(file_line.Text())
		sum += x*10 + y
	}
	fmt.Println(sum)
}
