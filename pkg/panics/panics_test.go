package panics_test

import (
	"errors"
	"runtime"
	"testing"

	"github.com/flare-foundation/go-flare-common/pkg/panics"
	"github.com/stretchr/testify/require"
)

func run(f func() error) (err error) {
	defer panics.Recover(&err)
	return f()
}

func TestRecover(t *testing.T) {
	errBoom := errors.New("boom")

	tests := []struct {
		name string
		f    func() error
		is   error  // errors.Is target
		as   any    // errors.As target
		msg  string // exact message
	}{
		{
			name: "error",
			f:    func() error { panic(errBoom) },
			is:   errBoom,
			msg:  "recovered panic: boom",
		},
		{
			name: "non error",
			f:    func() error { panic("boom") },
			msg:  "recovered panic non error: boom",
		},
		{
			name: "nil",
			f:    func() error { panic(nil) },
			as:   new(*runtime.PanicNilError),
		},
		{
			name: "runtime error",
			f: func() error {
				at := func(s []int, i int) int { return s[i] }
				_ = at(nil, 1)
				return nil
			},
			as: new(runtime.Error),
		},
		{
			name: "returned error kept",
			f:    func() error { return errBoom },
			is:   errBoom,
			msg:  "boom",
		},
		{
			name: "no error",
			f:    func() error { return nil },
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := run(tc.f)
			if tc.is == nil && tc.as == nil && tc.msg == "" {
				require.NoError(t, err)
			}
			if tc.is != nil {
				require.ErrorIs(t, err, tc.is)
			}
			if tc.as != nil {
				require.ErrorAs(t, err, tc.as)
			}
			if tc.msg != "" {
				require.EqualError(t, err, tc.msg)
			}
		})
	}
}

func TestRecoverNotDeferredDirectly(t *testing.T) {
	require.Panics(t, func() {
		var err error
		defer func() { panics.Recover(&err) }()
		panic("boom")
	})
}
