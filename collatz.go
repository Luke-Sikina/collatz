package main

import (
	"os"
	"strconv"
	"strings"
	"fmt"
)

func main() {
	start, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err == nil {
		println(strings.Trim(fmt.Sprint(Collatz(uint(start))), "[]"))
	}
}

func Collatz(start uint) (steps []uint) {
	steps = []uint{}
	for start > 1 {
		steps = append(steps, start)
		if start % 2 == 1 {
			start = start*3 + 1
		} else {
			start = start/2
		}
	}
	steps = append(steps, start)
	return
}
