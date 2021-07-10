package util

func DefaultString(val, def string) string {
	if val == "" {
		return def
	}

	return val
}

func If(b bool, fn func()) {
	if b {
		fn()
	}
}
