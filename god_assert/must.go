package god_assert

// MustNil if err != nil, panic
func MustNil(err error) {
	if err != nil {
		panic(err)
	}
}
