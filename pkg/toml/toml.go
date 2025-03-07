package toml

import (
	"fmt"
	"os"
	"reflect"

	"github.com/naoina/toml"
)

// ReadToml reads toml file from filePath and marshals it into struct of type T.
// If allowUnknownFields is set to false, any unknown field in toml file will trigger error.
func ReadToml[T any](filePath string, allowUnknownFields bool) (T, error) {
	var dest T

	file, err := os.ReadFile(filePath)
	if err != nil {
		return dest, fmt.Errorf("failed reading file %s with: %s", filePath, err)
	}

	tomlCfg := toml.DefaultConfig

	if allowUnknownFields {
		tomlCfg.MissingField = func(typ reflect.Type, key string) error { return nil }
	}

	err = tomlCfg.Unmarshal(file, &dest)
	if err != nil {
		return dest, fmt.Errorf("failed unmarshaling file %s with: %s", filePath, err)
	}

	return dest, nil
}

// ReadTomlTo reads toml file from filePath and marshals it into dest.
// If allowUnknownFields is set to false, any unknown field in toml file will trigger error.
func ReadTomlTo[T any](filePath string, dest *T, allowUnknownFields bool) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed reading file %s with: %s", filePath, err)
	}

	tomlCfg := toml.DefaultConfig
	if allowUnknownFields {
		tomlCfg.MissingField = func(typ reflect.Type, key string) error { return nil }
	}

	err = tomlCfg.Unmarshal(file, dest)
	if err != nil {
		return fmt.Errorf("failed unmarshaling file %s with: %s", filePath, err)
	}

	return nil
}
