package uic_go_sdl2

func ifElseAssign[T interface{}](condition bool, true T, false T) T {
	if condition {
		return true
	} else {
		return false
	}
}

func nilOrPanic(v interface{}) {
	if v != nil {
		panic(v)
	}
}
