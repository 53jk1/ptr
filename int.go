package ptr

// Int returns a pointer to the int value passed in.
func Int(i int) *int {
	return &i
}
