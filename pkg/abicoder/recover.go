package abicoder

import "fmt"

// recoverError turns a panic into *err. It must be deferred directly.
func recoverError(err *error) {
	if r := recover(); r != nil {
		if e, ok := r.(error); ok {
			*err = fmt.Errorf("recovered panic: %w", e)
		} else {
			*err = fmt.Errorf("recovered panic non error: %v", r)
		}
	}
}
