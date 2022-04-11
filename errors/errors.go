package errors

import (
	"errors"
	"fmt"
	"os"
)

var (
	New    = errors.New
	Errorf = fmt.Errorf

	Is     = errors.Is
	As     = errors.As
	Unwrap = errors.Unwrap
)

func IsTimeout(err error) bool {
	for err != nil {
		if os.IsTimeout(err) {
			return true
		}

		err = Unwrap(err)
	}

	return false
}
