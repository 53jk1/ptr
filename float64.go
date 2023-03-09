package ptr

// Float64 returns a pointer to the float64 value passed in.
func Float64(f float64) *float64 {
	return &f
}
