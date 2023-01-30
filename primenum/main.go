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
    factor_list := []int{}

	fmt.Print("Insert a number: ")
	fmt.Scan(&usr_in)
	usr_num, err := strconv.Atoi(usr_in)
	if err != nil {
		os.Exit(1)
	}
	for usr_num-pos_fac != -1 {
        if usr_num%pos_fac == 0 {
			factors += 1
            factor_list = append(factor_list, pos_fac)
        }  
        if factors == 2 && factor_list[1] == usr_num && factor_list[0] == 1 {
            fmt.Println("Is prime")
            break
        } else if factors > 2 {
            fmt.Println("Not prime")
        }
        pos_fac += 1
    }
}
