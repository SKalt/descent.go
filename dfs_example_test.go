package descent_test

import (
	"errors"
	"fmt"

	descent "github.com/skalt/descent.go"
)

func ExampleDepthFirst() {
	err := errors.Join(
		errors.New("a"),
		fmt.Errorf("b: %w",
			errors.Join(
				fmt.Errorf("c: %w", errors.New("d")),
				nil,
				errors.New("e"),
			),
		),
		errors.New("f"),
	)
	i := 0
	for e := range descent.DepthFirst(err) {
		fmt.Printf("%d %T(%#v)\n", i, e, e.Error())
		i += 1
	}
	// Output:
	// 0 *errors.joinError("a\nb: c: d\ne\nf")
	// 1 *errors.errorString("a")
	// 2 *fmt.wrapError("b: c: d\ne")
	// 3 *errors.joinError("c: d\ne")
	// 4 *fmt.wrapError("c: d")
	// 5 *errors.errorString("d")
	// 6 *errors.errorString("e")
	// 7 *errors.errorString("f")
}
