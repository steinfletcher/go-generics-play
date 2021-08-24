package main

import "fmt"

// Playground: https://go2goplay.golang.org/p/7lPwEGViFkh

// Ordered is a "type list" and is used as a constraint in min. It allows a constraint interface to have a list of types.
// I hope some useful constraints like this are exported in the std lib.
type Ordered interface {
	type int, int64, float32, float64
}

func min[T Ordered](s []T) T {
	if len(s) == 0 {
		panic("expected a non-empty input")
	}
	var min = s[0]
	for _, v := range s {
		if v < min { // < is alowed because of the Ordered constraint
			min = v
		}
	}
	return min
}

func minWithConstraintInlined[T interface{type int, int64, float32, float64}](s []T) T {
	if len(s) == 0 {
		panic("expected a non-empty input")
	}
	var min = s[0]
	for _, v := range s {
		if v < min { // < is alowed because of the Ordered constraint
			min = v
		}
	}
	return min
}

// "any" stands for no constraint (same as interface{}). "any" is only supported in the constraint position, but
// it might be opened up to use as a general type in future
func print[T any](v T) {
	fmt.Println(fmt.Sprintf("%+v", v))
}

func main() {
	print(min([]int{2, 1, 5}))

	// min([]string{"2", "A"}) doesn't compile as string doesn't satisfy the Ordered constraint

	// We can also inline constraints
	print(minWithConstraintInlined([]int{2, 1, 5}))

	// this instantiates a min function that works with ints only. This works because
	// the compiler generates a non-generic version of the min function, then separately
	// performs the invocation. In other words, the compiler "substitutes" the generic function
	// with an explicit version of the function with the int replacing the type argument.
	// This can be thought of as a two-step process -
	// 1) substitution/instantiation 2) function invocation (as usual)
	minner := min[int]
	print(minner([]int{31, 9}))
}