package assertion

import "fmt"

func Require(b bool, msg string) {
	if !b {
		panic(msg)
	}
}

func Requiref(b bool, msg string, args ...any) {
	if !b {
		panic(fmt.Sprintf("precondition violation: "+msg, args))
	}
}

func Ensure(b bool, msg string) {
	if !b {
		panic(msg)
	}
}

func Ensuref(b bool, msg string, args ...any) {
	if !b {
		panic(fmt.Sprintf("postcondition violation: "+msg, args...))
	}
}

func Invariant(b bool, msg string) {
	if !b {
		panic(msg)
	}
}

func Invariantf(b bool, msg string, args ...any) {
	if !b {
		panic(fmt.Sprintf("invariant violation: "+msg, args...))
	}
}
