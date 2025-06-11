package main

import (
	"bufio"
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
	a := ps.NewStack(nums)
	b := ps.NewStack(nil)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inst := scanner.Text()
		if inst == "" {
			continue
		}
		if err := ps.Execute(inst, a, b); err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
	}

	if ps.IsSorted(a.Data()) && b.Len() == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
