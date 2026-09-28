package logger

import (
	"fmt"
	"os"
)

// Nop discards output from Debug, Info, Warn and Error calls and their variants.
//
// Panic calls still panic, and Fatal calls still exit with status 1.
type Nop struct{}

func (Nop) Debug(_ ...any)            {}
func (Nop) Debugf(_ string, _ ...any) {}
func (Nop) Debugw(_ string, _ ...any) {}

func (Nop) Info(_ ...any)            {}
func (Nop) Infof(_ string, _ ...any) {}
func (Nop) Infow(_ string, _ ...any) {}

func (Nop) Warn(_ ...any)            {}
func (Nop) Warnf(_ string, _ ...any) {}
func (Nop) Warnw(_ string, _ ...any) {}

func (Nop) Error(_ ...any)            {}
func (Nop) Errorf(_ string, _ ...any) {}
func (Nop) Errorw(_ string, _ ...any) {}

func (Nop) Panic(args ...any)                 { panic(fmt.Sprint(args...)) }
func (Nop) Panicf(format string, args ...any) { panic(fmt.Sprintf(format, args...)) }
func (Nop) Panicw(msg string, _ ...any)       { panic(msg) }

func (Nop) Fatal(_ ...any)            { os.Exit(1) }
func (Nop) Fatalf(_ string, _ ...any) { os.Exit(1) }
func (Nop) Fatalw(_ string, _ ...any) { os.Exit(1) }
