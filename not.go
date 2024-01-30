package logic_gates

// Not returns true if a is false.
// Table:
// in | out
// 0  | 1
// 1  | 0
func Not(a bool) bool {
	return !a
}
