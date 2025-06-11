package ps

import "testing"

func TestSortInstructions(t *testing.T) {
	nums := []int{4, 3, 2, 1}
	ops := SortInstructions(nums)
	a := NewStack(nums)
	b := NewStack(nil)
	for _, op := range ops {
		if err := Execute(op, a, b); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}
	if !IsSorted(a.Data()) || b.Len() != 0 {
		t.Fatalf("stack not sorted after executing ops")
	}
}
