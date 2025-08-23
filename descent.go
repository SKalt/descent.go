package descent

import (
	"iter"
)

func depthFirst(yield func(error) bool, errs ...error) (shouldContinue bool) {
	for _, err := range errs {
		if err != nil {
			shouldContinue = yield(err)
			if !shouldContinue {
				return
			}
			switch e := err.(type) {
			case interface{ Unwrap() error }:
				shouldContinue = depthFirst(yield, e.Unwrap())
				if !shouldContinue {
					return
				}
			case interface{ Unwrap() []error }:
				shouldContinue = depthFirst(yield, e.Unwrap()...)
				if !shouldContinue {
					return
				}
			}
		}
	}
	return true
}

// Iterates from the root error over the tree of errors defined by `Unwrap() error` and
// `Unwrap() []error` interfaces. Yields each error and its children before moving to the
// next sibling.
//
// This is the same order used by [errors.Is] [errors.As].
func DepthFirst(err error) iter.Seq[error] {
	return func(yield func(error) bool) {
		depthFirst(yield, err)
	}
}

func breadthFirst(yield func(error) bool, errs ...error) bool {
	children := []error{}
	for _, err := range errs {
		if err != nil {
			if !yield(err) {
				return false
			}
			switch e := err.(type) {
			case interface{ Unwrap() error }:
				children = append(children, e.Unwrap())
			case interface{ Unwrap() []error }:
				children = append(children, e.Unwrap()...)
			}

		}
	}
	if len(children) > 0 {
		return breadthFirst(yield, children...)
	} else {
		return true
	}
}

// Iterates from the root error over the tree of errors defined by `Unwrap() error` and
// `Unwrap() []error` interfaces. Yields each error's siblings before moving to their
// children.
//
// This is a different order than [errors.Is] [errors.As].
func BreadthFirst(err error) iter.Seq[error] {
	return func(yield func(error) bool) {
		breadthFirst(yield, err)
	}
}
