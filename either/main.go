package main

import "fmt"

type Either[L, R any] struct {
	Left         L
	Right        R
	isLeftBiased bool
}

func Right[L any, R any](value R) Either[L, R] {
	return Either[L, R]{Right: value, isLeftBiased: false}
}

func Left[L any, R any](value L) Either[L, R] {
	return Either[L, R]{Left: value, isLeftBiased: true}
}

func (e Either[L, R]) Map(fn func(val L) L) Either[L, R] {
	if !e.isLeftBiased {
		return e
	}
	return Either[L, R]{
		Left:         fn(e.Left),
		Right:        e.Right,
		isLeftBiased: e.isLeftBiased,
	}
}

func main() {
	l := Left[string, error]("some result")

	updated := l.Map(func(name string) string {
		return name + " updated"
	})

	fmt.Println(updated)
}

