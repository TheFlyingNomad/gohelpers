package numbers

import (
	"errors"
	"math/rand"
	"time"
)

// Random -
func Random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// ReadNumericFromInterface -
func ReadNumericFromInterface(data interface{}) (int, error) {
	var value = -1
	switch t := data.(type) {
	case int:
		value = int(t)
	case int8:
		value = int(t)
	case int16:
		value = int(t)
	case int32:
		value = int(t)
	case int64:
		value = int(t)
	case float32:
		value = int(t)
	case float64:
		value = int(t)
	case uint8:
		value = int(t)
	case uint16:
		value = int(t)
	case uint32:
		value = int(t)
	case uint64:
		value = int(t)
	default:
		return value, errors.New("No numeric found")
	}

	return value, nil
}
