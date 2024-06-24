package division

import (
	"errors"
)

func Divide(val1, val2 int) (float64, error) {
	if val2 == 0 {

		return 0.0, errors.New("division by zero is not possible")
	}

	return float64(val1 / val2), nil

}
