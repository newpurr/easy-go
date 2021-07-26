package assert

func MustTrue(t bool, message string) {
	if !t {
		panic(message)
	}
}

func MustFalse(t bool, message string) {
	MustTrue(!t, message)
}
