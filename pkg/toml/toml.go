package toml

import (
	"fmt"
	"os"
	"reflect"

	"github.com/naoina/toml"
)

// ReadToml reads toml file from filePath and marshals it into struct of type T.
func ReadToml[T any](filePath string, allowUnknownFields bool) (T, error) {
	var config T

	file, err := os.ReadFile(filePath)
	if err != nil {
		return config, fmt.Errorf("failed reading file %s with: %s", filePath, err)
	}

	tomlCfg := toml.DefaultConfig

	if allowUnknownFields {
		tomlCfg.MissingField = func(typ reflect.Type, key string) error { return nil }
	}

	err = tomlCfg.Unmarshal(file, &config)
	if err != nil {
		return config, fmt.Errorf("failed unmarshaling file %s with: %s", filePath, err)
	}

	return config, nil
}

// ReadToml reads toml file from filePath and marshals it into struct of type T.
func ReadTomlTo[T any](filePath string, dest *T) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed reading file %s with: %s", filePath, err)
	}

	err = toml.Unmarshal(file, dest)
	if err != nil {
		return fmt.Errorf("failed unmarshaling file %s with: %s", filePath, err)
	}

	return nil
}
