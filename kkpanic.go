package kkpanic

// P to determine whether the err is nil. If not nil, panic.
func P(err error) {
	if err != nil {
		panic(err)
	}
}
