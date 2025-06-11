package ps

// SortInstructions returns a list of instructions to sort numbers using
// the push-swap operations. The algorithm is not optimal but strives to
// use few rotations by always moving the maximal element from stack b
// back to stack a.
func SortInstructions(nums []int) []string {
	a := NewStack(nums)
	b := NewStack(nil)
	ops := []string{}

	// Push everything to stack b
	for a.Len() > 0 {
		if Pb(a, b) {
			ops = append(ops, "pb")
		}
	}

	for b.Len() > 0 {
		// find index of max element in b
		maxIdx := 0
		maxVal := b.data[0]
		for i, v := range b.data {
			if v > maxVal {
				maxVal = v
				maxIdx = i
			}
		}
		// Bring max to top using the shortest rotation
		if maxIdx <= b.Len()/2 {
			for i := 0; i < maxIdx; i++ {
				if Rb(b) {
					ops = append(ops, "rb")
				}
			}
		} else {
			for i := 0; i < b.Len()-maxIdx; i++ {
				if Rrb(b) {
					ops = append(ops, "rrb")
				}
			}
		}
		if Pa(a, b) {
			ops = append(ops, "pa")
		}
	}
	return ops
}
