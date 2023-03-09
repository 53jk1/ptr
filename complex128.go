package ptr

// Complex128 returns a pointer to the complex128 value passed in.
func Complex128(c complex128) *complex128 {
	return &c
}
