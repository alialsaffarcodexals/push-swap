package ps

import "fmt"

// Execute applies an instruction string to the stacks.
func Execute(inst string, a, b *Stack) error {
	switch inst {
	case "pa":
		Pa(a, b)
	case "pb":
		Pb(a, b)
	case "sa":
		Sa(a)
	case "sb":
		Sb(b)
	case "ss":
		Ss(a, b)
	case "ra":
		Ra(a)
	case "rb":
		Rb(b)
	case "rr":
		Rr(a, b)
	case "rra":
		Rra(a)
	case "rrb":
		Rrb(b)
	case "rrr":
		Rrr(a, b)
	default:
		return fmt.Errorf("invalid instruction")
	}
	return nil
}
