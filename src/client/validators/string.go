package validators

import "errors"

// Validates that the given string is not empty.
func String(input string, size uint32) error {
	if input == "" {
		err := errors.New("string is empty")

		return err
	}

	return nil
}
