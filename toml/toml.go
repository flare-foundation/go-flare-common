package toml

import (
	"fmt"
	"os"

	"github.com/naoina/toml"
)

// ReadToml reads toml file from filePath and marshals it into struct of type T.
func ReadToml[T any](filePath string) (T, error) {
	var config T

	file, err := os.ReadFile(filePath)
	if err != nil {
		return config, fmt.Errorf("failed reading file %s with: %s", filePath, err)
	}

	err = toml.Unmarshal(file, &config)
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
