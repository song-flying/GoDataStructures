package contract

import "fmt"

const On = true

func Require(b bool, msg string) {
	if On && !b {
		panic("pre-condition violation: " + msg)
	}
}

// Requiref can be used for debugging to print the actual values
func Requiref(b bool, msg string, args ...any) {
	if On && !b {
		panic(fmt.Sprintf("pre-condition violation: "+msg, args...))
	}
}

func Ensure(b bool, msg string) {
	if On && !b {
		panic("post-condition violation: " + msg)
	}
}

// Ensuref can be used for debugging to print the actual values
func Ensuref(b bool, msg string, args ...any) {
	if On && !b {
		panic(fmt.Sprintf("post-condition violation: "+msg, args...))
	}
}

func Invariant(b bool, msg string) {
	if On && !b {
		panic("invariant violation: " + msg)
	}
}

// Invariantf can be used for debugging to print the actual values
func Invariantf(b bool, msg string, args ...any) {
	if On && !b {
		panic(fmt.Sprintf("invariant violation: "+msg, args...))
	}
}

func Assert(b bool, msg string) {
	if On && !b {
		panic("assertion violation: " + msg)
	}
}

// Assertf can be used for debugging to print the actual values
func Assertf(b bool, msg string, args ...any) {
	if On && !b {
		panic(fmt.Sprintf("assertion violation: "+msg, args...))
	}
}
