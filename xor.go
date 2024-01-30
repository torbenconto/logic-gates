package logic_gates

// Xor returns the logical XOR of two boolean values.
// Table:
// in | in | out
// 0  | 0  | 0
// 0  | 1  | 1
// 1  | 0  | 1
// 1  | 1  | 0
func Xor(a, b bool) bool {
	return a != b
}
