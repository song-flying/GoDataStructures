package assertion

import "fmt"

func Require(b bool, msg string) {
	if !b {
		panic("pre-condition violation: " + msg)
	}
}

func Requiref(b bool, msg string, args ...any) {
	if !b {
		panic(fmt.Sprintf("pre-condition violation: "+msg, args))
	}
}

func Ensure(b bool, msg string) {
	if !b {
		panic("post-condition violation: " + msg)
	}
}

func Ensuref(b bool, msg string, args ...any) {
	if !b {
		panic(fmt.Sprintf("post-condition violation: "+msg, args...))
	}
}

func Invariant(b bool, msg string) {
	if !b {
		panic("invariant violation: " + msg)
	}
}

func Invariantf(b bool, msg string, args ...any) {
	if !b {
		panic(fmt.Sprintf("invariant violation: "+msg, args...))
	}
}

func Check(b bool, msg string) {
	if !b {
		panic("assertion violation: " + msg)
	}
}
