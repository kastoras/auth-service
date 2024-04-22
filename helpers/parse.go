package helpers

import (
	"fmt"
	"strconv"
)

func ParseQueryVarToInt(stringVal string) (int, error) {
	if stringVal == "" {
		return -1, fmt.Errorf("empty string value")
	}

	i, err := strconv.Atoi(stringVal)
	if err != nil {
		return -1, err
	}

	return i, nil
}
