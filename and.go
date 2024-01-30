package logic_gates

// And returns true if both a and b are true.
// Table:
// in | in | out
// 0  | 0  | 0
// 0  | 1  | 0
// 1  | 0  | 0
// 1  | 1  | 1
func And(a, b bool) bool {
	return a && b
}
