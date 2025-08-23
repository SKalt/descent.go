package descent_test

import (
	"errors"
	"fmt"

	descent "github.com/skalt/descent.go"
)

func ExampleBreadthFirst() {
	err := errors.Join(
		errors.New("a"),
		fmt.Errorf("b: %w",
			errors.Join(
				fmt.Errorf("d: %w", errors.New("f")),
				nil,
				errors.New("e"),
			),
		),
		errors.New("c"),
	)
	i := 0
	for e := range descent.BreadthFirst(err) {
		fmt.Printf("%d %T(%#v)\n", i, e, e.Error())
		if e.Error() == "e" {
			break
		}
		i += 1
	}
	// Output:
	// 0 *errors.joinError("a\nb: d: f\ne\nc")
	// 1 *errors.errorString("a")
	// 2 *fmt.wrapError("b: d: f\ne")
	// 3 *errors.errorString("c")
	// 4 *errors.joinError("d: f\ne")
	// 5 *fmt.wrapError("d: f")
	// 6 *errors.errorString("e")
}
