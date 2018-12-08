package numbers

import (
	"math/rand"
	"time"
)

// Random -
func Random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
