package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var usr_in string
	factors := 0
	pos_fac := 1

	fmt.Print("Insert a number: ")
	fmt.Scan(&usr_in)
	usr_num, err := strconv.Atoi(usr_in)
	if err != nil {
		os.Exit(1)
	}
	for usr_num-pos_fac != -1 {
		if usr_num%pos_fac == 0 {
			factors += 1
			pos_fac += 1
		} else {
			pos_fac += 1
		}
	}
	if factors == 2 {
		fmt.Println("Is prime")
	} else {
		fmt.Println("Not Prime")
	}
}
