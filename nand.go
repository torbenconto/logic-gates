package logic_gates

// Nand returns true if both a and b are false.
// Table:
// in | in | out
// 0  | 0  | 1
// 0  | 1  | 1
// 1  | 0  | 1
// 1  | 1  | 0
func Nand(a, b bool) bool {
	return !(a && b)
}
