package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func is_num(ch rune) bool {
	return ch >= 0 && ch <= 9
}

func reverse(s string) []rune {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return rns
}

func get_int(ch rune) int {
	n, _ := strconv.Atoi(string(ch))
	return n
}

func should_be_summarized(len int, matr [][]rune, first_id, last_id, row int) bool {
	dot_count := 1
	dots_cloud_nine := 8
	for i := 1; i < len; i++ {
		dots_cloud_nine += 2
	}
	if row == 0 || row == 139 {
		dots_cloud_nine -= 2 + len
	}
	if first_id < 0 {
		first_id = 0
	}
	if first_id == 0 {
		dots_cloud_nine -= 3
	}
	if last_id > 141 {
		last_id = 141
	}
	if last_id == 141 {
		dots_cloud_nine -= 3
		if row == 0 || row == 139 {
			dots_cloud_nine += 1
		}
	}
	st_row := row - 1
	en_row := row + 1
	if row == 0 {
		st_row = row
	} else if row == 139 {
		en_row = row
	}
	for i := first_id; i <= last_id; i++ {
		for j := st_row; j <= en_row; j++ {
			if matr[i][j] == -2 {
				dot_count++
			}
		}
	}
	return dot_count < dots_cloud_nine
}

func get_array_of_runes() [][]rune {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	lines := bufio.NewScanner(file)
	matrix := [][]rune{}
	for lines.Scan() {
		tmp := []rune{}
		for _, ch := range lines.Text() {
			tmp = append(tmp, ch-48)
		}
		matrix = append(matrix, tmp)
	}
	return matrix
}

func main() {
	matrix := get_array_of_runes()
	numbers := []int{}
	for i, v := range matrix {
		v = reverse(string(v))
		start_i := -1
		end_i := -1
		len := 0
		num := 0
		for in, vn := range v {
			if is_num(vn) {
				if start_i == -1 {
					start_i = in
					len = 1
					end_i = in
				}
				num += int(vn) * len
				len *= 10
				end_i = in
			} else {
				if start_i != -1 {
					len = end_i - start_i + 1
					if should_be_summarized(len, matrix, start_i, end_i, i) {
						numbers = append(numbers, num)
					}
					start_i = -1
					end_i = -1
					num = 0
					len = 0
				}
			}
		}
	}
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	fmt.Println(sum)
}
