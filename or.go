package logic_gates

// Or returns true if either a or b is true.
// Table:
// in | in | out
// 0  | 0  | 0
// 0  | 1  | 1
// 1  | 0  | 1
// 1  | 1  | 1
func Or(a, b bool) bool {
	return a || b
}
