package ptr

// String returns a pointer to the string value passed in.
func String(s string) *string {
	return &s
}
