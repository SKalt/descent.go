# `descent`

[![Go Reference](https://pkg.go.dev/badge/github.com/skalt/descent.go.svg)](https://pkg.go.dev/github.com/skalt/descent.go)

Functions that iterate over trees of errors defined using
```go
interface { Unwrap()   error }
interface { Unwrap() []error }
```
These interfaces are how the [`errors`][errors] package implements error comparison via [`errors.Is`][errors.Is] and [`errors.As`][errors.As].




[errors]: https://pkg.go.dev/errors
[errors.Is]: https://pkg.go.dev/errors#Is
[errors.As]: https://pkg.go.dev/errors#As
