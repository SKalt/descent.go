package descent_test

import (
	"errors"
	"fmt"
	"testing"

	descent "github.com/skalt/descent.go"
)

type noChildren string

func (err noChildren) Error() string { return string(err) }
func (noChildren) Unwrap() []error   { return nil }

func ExampleBreadthFirst_noChildren() {
	err := errors.Join(
		errors.New("a"),
		noChildren("b"),
		errors.Join(errors.New("c")),
	)
	for e := range descent.BreadthFirst(err) {
		switch t := e.(type) {
		case interface{ Unwrap() []error }:
			{
				children := t.Unwrap()
				fmt.Printf("%q: %d children\n", e.Error(), len(children))
				for i, child := range children {
					fmt.Printf("  %d. %q\n", i, child.Error())
				}
			}
		case interface{ Unwrap() error }:
			{
				fmt.Printf("%q: 1 child\n", e.Error())
				child := t.Unwrap()
				fmt.Printf("  0. %q\n", child.Error())
			}
		default:
			fmt.Printf("%q (no children)\n", e.Error())
		}
	}
	// Output:
	// "a\nb\nc": 3 children
	//   0. "a"
	//   1. "b"
	//   2. "c"
	// "a" (no children)
	// "b": 0 children
	// "c": 1 children
	//   0. "c"
	// "c" (no children)
}

func TestDepthFirst_break(t *testing.T) {
	err := /*0*/ errors.Join(
		/*1*/ errors.New("a"),
		/*2*/ noChildren("b"),
		/*3*/ errors.Join( /*4*/ errors.New("c")),
		/*5*/ fmt.Errorf("d: %w" /*6*/, errors.New("e")),
		/*7*/ errors.New("f"),
		/*8*/ errors.New("g"), // should not be reached
	)
	expected := []string{
		"a\nb\nc\nd: e\nf\ng", // 0
		"a",                   // 1
		"b",                   // 2
		"c",                   // 3
		"c",                   // 4
		"d: e",                // 5
		"e",                   // 6
		"f",                   // 7
		"g",
	}
	check := func(cutoff int) {
		_expected := expected[:cutoff]
		actual := []string{}
		i := 0
		for e := range descent.DepthFirst(err) {
			actual = append(actual, e.Error())
			i += 1
			if i >= cutoff {
				break
			}
		}
		if len(actual) != len(_expected) {
			t.Fatalf("expected %d results, got %d", len(_expected), len(actual))
		}
		for i := range _expected {
			if expected[i] != actual[i] {
				t.Errorf("at index %d, expected %q, got %q", i, expected[i], actual[i])
			}
		}
	}
	check(3)
	check(7)

}
