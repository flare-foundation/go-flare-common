// Package toml provides utilities for reading and unmarshaling TOML configuration files. Wraps github.com/BurntSushi/toml.
package toml

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// Read reads toml file from filePath and marshals it into struct of type T.
// If allowUnknownFields is set to false, any unknown field in toml file will trigger error.
func Read[T any](filePath string, allowUnknownFields bool) (T, error) {
	var dest T

	err := ReadTo(filePath, &dest, allowUnknownFields)

	return dest, err
}

// ReadTo reads toml file from filePath and marshals it into dest.
// If allowUnknownFields is set to false, any unknown field in toml file will trigger error.
func ReadTo[T any](filePath string, dest *T, allowUnknownFields bool) error {
	md, err := toml.DecodeFile(filePath, dest)

	if err != nil {
		return fmt.Errorf("decoding file %s: %w", filePath, err)
	}

	if !allowUnknownFields && len(md.Undecoded()) > 0 {
		return fmt.Errorf("unknown field in toml %v", md.Undecoded()[0].String())
	}

	return nil
}
