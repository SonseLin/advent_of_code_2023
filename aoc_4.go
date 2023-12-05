package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_ticket_value(winnings, ticket_values []int) int {
	worth := 0
	used_nums := []int{}
	for _, w := range winnings {
		for _, tv := range ticket_values {
			if w == tv {
				if worth == 0 {
					worth = 1
					used_nums = append(used_nums, tv)
					break
				} else {
					already_used := 0
					for _, n := range used_nums {
						if n == tv {
							already_used = 1
							break
						}
					}
					if already_used == 0 {
						worth *= 2
						used_nums = append(used_nums, tv)
					}
					break
				}
			}
		}
	}
	for _, d := range used_nums {
		fmt.Printf("%d ", d)
	}
	fmt.Println()
	return worth
}

func print_arr(arr []int) {
	for _, d := range arr {
		fmt.Printf("%d ", d)
	}
	fmt.Println()
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		s := strings.ReplaceAll(scanner.Text(), ":", "")
		arr := strings.Split(s, " ")
		id := arr[1]
		res := []int{}
		for _, d := range arr[2:] {
			if num, err := strconv.Atoi(d); err == nil {
				res = append(res, num)
			} else if d == "|" {
				break
			}
		}
		pipe_found := 0
		ticket_values := []int{}
		for _, d := range arr {
			if pipe_found == 1 {
				if num, err := strconv.Atoi(d); err == nil {
					ticket_values = append(ticket_values, num)
				}
			} else if d == "|" {
				pipe_found = 1
			}
		}

		if worth := get_ticket_value(res, ticket_values); worth != 0 {
			fmt.Println("array of winnings")
			print_arr(res)
			fmt.Println("array of ticket")
			print_arr(ticket_values)
			fmt.Println("String itself")
			fmt.Println(s)
			fmt.Printf("id [%s] has worth of %d\n", id, worth)
			total += worth
		}
	}
	fmt.Println(total)
}
