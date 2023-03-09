package ptr

// Bool returns a pointer to the bool value passed in.
func Bool(b bool) *bool {
	return &b
}
