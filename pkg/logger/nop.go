package logger

// Nop discards all log output. It satisfies any logger interface
// whose methods are a subset of the ones defined here.
type Nop struct{}

func (Nop) Debugf(_ string, _ ...any) {}
func (Nop) Debug(_ ...any)            {}
func (Nop) Infof(_ string, _ ...any)  {}
func (Nop) Info(_ ...any)             {}
func (Nop) Warnf(_ string, _ ...any)  {}
func (Nop) Errorf(_ string, _ ...any) {}
