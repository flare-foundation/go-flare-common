// Package panics turns panics into errors.
package panics

import "fmt"

// Recover turns a panic into *err, wrapping the panic value if it is an error.
// It must be deferred directly, as in defer panics.Recover(&err); called from
// within another deferred function it recovers nothing.
func Recover(err *error) {
	if r := recover(); r != nil {
		if e, ok := r.(error); ok {
			*err = fmt.Errorf("recovered panic: %w", e)
		} else {
			*err = fmt.Errorf("recovered panic non error: %v", r)
		}
	}
}
