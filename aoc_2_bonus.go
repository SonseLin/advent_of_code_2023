package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func delim(r rune) bool {
	return r == ':' || r == ',' || r == ' '
}

func main() {
	dat, err := os.Open("input.txt")
	check(err)
	defer dat.Close()
	file_line := bufio.NewScanner(dat)
	id_sum := 0
	for file_line.Scan() {
		s := strings.Replace(file_line.Text(), ":", "", -1)
		s = strings.Replace(s, ",", "", -1)
		s = strings.Replace(s, ";", "", -1)
		lexems := strings.Split(s, " ")
		var cubes int
		r_max, b_max, g_max := 0, 0, 0
		for ind, v := range lexems[2:] {
			if ind%2 != 0 {
				switch v {
				case "red":
					if cubes > r_max {
						r_max = cubes
					}
				case "blue":
					if cubes > b_max {
						b_max = cubes
					}
				case "green":
					if cubes > g_max {
						g_max = cubes
					}
				}
			} else {
				cubes, _ = strconv.Atoi(v)
			}
		}
		id_sum += r_max * g_max * b_max
	}
	fmt.Println(id_sum)
}
