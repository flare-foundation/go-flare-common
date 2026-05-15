package logger

import "fmt"

// Nop discards Debug/Info/Warn/Error output. It satisfies any logger
// interface whose methods are a subset of the ones defined here.
//
// Panic and Fatal are NOT silent: they preserve the terminating-call
// contract of the real logger so that downstream code which writes
// `logger.Fatal("invariant"); return ok` does not silently continue
// when a Nop is substituted (e.g., in tests or as a default). Both
// methods panic — Fatal via panic rather than os.Exit so tests can
// recover, while still guaranteeing the call does not return.
type Nop struct{}

func (Nop) Debug(_ ...any)            {}
func (Nop) Debugf(_ string, _ ...any) {}

func (Nop) Info(_ ...any)            {}
func (Nop) Infof(_ string, _ ...any) {}

func (Nop) Warn(_ ...any)            {}
func (Nop) Warnf(_ string, _ ...any) {}

func (Nop) Error(_ ...any)            {}
func (Nop) Errorf(_ string, _ ...any) {}

func (Nop) Panic(args ...any)                 { panic(fmt.Sprint(args...)) }
func (Nop) Panicf(format string, args ...any) { panic(fmt.Sprintf(format, args...)) }

func (Nop) Fatal(args ...any)                 { panic(fmt.Sprint(args...)) }
func (Nop) Fatalf(format string, args ...any) { panic(fmt.Sprintf(format, args...)) }
