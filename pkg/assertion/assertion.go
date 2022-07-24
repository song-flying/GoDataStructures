package assertion

import "fmt"

func Require(b bool, msg string) {
	if !b {
		panic("pre-condition violation: " + msg)
	}
}

// Requiref can be used for debugging to print the actual values
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

// Ensuref can be used for debugging to print the actual values
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

// Invariantf can be used for debugging to print the actual values
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

// Checkf can be used for debugging to print the actual values
func Checkf(b bool, msg string, args ...any) {
	if !b {
		panic(fmt.Sprintf("assertion violation: "+msg, args...))
	}
}
