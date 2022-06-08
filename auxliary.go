package uic_go_sdl2

import "strconv"

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

func parseInt32(s string) int32 {
	n, err := strconv.ParseInt(s, 10, 64)
	nilOrPanic(err)
	return int32(n)
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	nilOrPanic(err)
	return n
}

func parseUint8(s string) uint8 {
	n, err := strconv.Atoi(s)
	nilOrPanic(err)
	return uint8(n)
}
