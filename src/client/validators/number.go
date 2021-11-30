package validators

import (
	"errors"
	"strconv"
)

// Check if the given value is a number.
func Number(input string, size uint32) error {
	i, err := strconv.Atoi(input)

	if err != nil {

		return err
	}

	u32 := uint32(i + 1)

	if (u32 < 1) || (u32 > size) {
		err := errors.New("please enter a valid number")

		return err
	}

	return nil
}
