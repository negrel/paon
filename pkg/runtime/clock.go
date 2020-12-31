package runtime

import (
	"time"
)

// Clock is responsible for triggering screen update.
// Default value is 16ms (60 fps).
var Clock = time.NewTicker(time.Millisecond * 16)
