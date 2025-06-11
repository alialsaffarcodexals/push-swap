package main

import (
	"fmt"
	"os"

	"github.com/yourname/push-swap/ps"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	nums, err := ps.ParseNumbers(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}
	if ps.IsSorted(nums) {
		return
	}
	ops := ps.SortInstructions(nums)
	for i, op := range ops {
		if i == len(ops)-1 {
			fmt.Print(op)
		} else {
			fmt.Println(op)
		}
	}
}
