package toml

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
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
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed reading file %s with: %s", filePath, err)
	}

	dec := toml.NewDecoder(file)

	if !allowUnknownFields {
		dec = dec.DisallowUnknownFields()
	}

	err = dec.Decode(&dest)
	if err != nil {
		return fmt.Errorf("failed unmarshaling file %s with: %s", filePath, err)
	}

	err = file.Close()
	return err
}
