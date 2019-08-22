// Package tc provides timecode (including 29.97 dropframe) calculations.
package tc

import (
	"time"
)

const (
	tDF29_97 = time.Duration(1001 * 1000000000 / 30000)
	tN25     = time.Duration(1000000000 / 25)
)

// TC is the TimeCode interface.
type TC interface {
	Add(tc TC) TC
	Sub(tc TC) TC
	Dur() time.Duration
	FrameCount() (int, error)
	String() string
}

func validTC(tc string) bool {
	if len(tc) < 11 {
		return false
	}
	return true
}
