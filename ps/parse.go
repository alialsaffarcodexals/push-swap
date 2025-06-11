package ps

import (
	"errors"
	"strconv"
	"strings"
)

// ParseNumbers parses arguments that may include quoted lists and returns
// the numbers as a slice. It verifies duplicates and returns an error on
// invalid input.
func ParseNumbers(args []string) ([]int, error) {
	var fields []string
	for _, arg := range args {
		fields = append(fields, strings.Fields(arg)...)
	}
	nums := make([]int, len(fields))
	seen := make(map[int]struct{})
	for i, f := range fields {
		n, err := strconv.Atoi(f)
		if err != nil {
			return nil, errors.New("not an integer")
		}
		if _, ok := seen[n]; ok {
			return nil, errors.New("duplicate")
		}
		seen[n] = struct{}{}
		nums[i] = n
	}
	return nums, nil
}
