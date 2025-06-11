package ps

// Instruction functions operate on stacks and return true if executed.

func Pa(a, b *Stack) bool {
	v, ok := b.Pop()
	if !ok {
		return false
	}
	a.Push(v)
	return true
}

func Pb(a, b *Stack) bool {
	v, ok := a.Pop()
	if !ok {
		return false
	}
	b.Push(v)
	return true
}

func Sa(a *Stack) bool { return a.Swap() }
func Sb(b *Stack) bool { return b.Swap() }

func Ss(a, b *Stack) bool {
	sa := a.Swap()
	sb := b.Swap()
	return sa || sb
}

func Ra(a *Stack) bool { return a.Rotate() }
func Rb(b *Stack) bool { return b.Rotate() }

func Rr(a, b *Stack) bool {
	ra := a.Rotate()
	rb := b.Rotate()
	return ra || rb
}

func Rra(a *Stack) bool { return a.ReverseRotate() }
func Rrb(b *Stack) bool { return b.ReverseRotate() }

func Rrr(a, b *Stack) bool {
	rra := a.ReverseRotate()
	rrb := b.ReverseRotate()
	return rra || rrb
}
