package tc

import (
	"fmt"
	"strconv"
	"time"
)

// N25 represent 25 fps Non-drop frame (ie. PAL rate)
type N25 struct {
	tc string
}

// NewN25 returns a new 25fps non-drop frame timecode.
func NewN25(tc string) *N25 {
	if !validTC(tc) {
		return &N25{"00:00:00:00"}
	}
	tc = tc[0:8] + ":" + tc[9:11]
	return &N25{tc}
}

// FrameCount returns an N25 timecode frame count.
func (c *N25) FrameCount() (int, error) {
	hh, err := strconv.Atoi(c.tc[0:2])
	if err != nil {
		return 0, err
	}
	mm, err := strconv.Atoi(c.tc[3:5])
	if err != nil {
		return 0, err
	}
	ss, err := strconv.Atoi(c.tc[6:8])
	if err != nil {
		return 0, err
	}
	ff, err := strconv.Atoi(c.tc[9:11])
	if err != nil {
		return 0, err
	}
	return hh*3600*25 + mm*60*25 + ss*25 + ff, nil
}

// NewN25FrameCount returns a new timecode given a frame count.
func NewN25FrameCount(fc int) *N25 {
	hh := fc / (3600 * 25)
	mm := (fc - hh*3600*25) / (60 * 25)
	ss := (fc - hh*3600*25 - mm*60*25) / 25
	ff := fc - hh*3600*25 - mm*60*25 - ss*25
	return NewN25(fmt.Sprintf("%02d:%02d:%02d:%02d", hh, mm, ss, ff))
}

// Add adds timecodes c + tc.
func (c *N25) Add(tc TC) TC {
	fc1, _ := c.FrameCount()
	fc2, _ := tc.FrameCount()
	return NewN25FrameCount(fc1 + fc2)
}

// Sub subtracts timecodes c - tc.
func (c *N25) Sub(tc TC) TC {
	fc1, _ := c.FrameCount()
	fc2, _ := tc.FrameCount()
	return NewN25FrameCount(fc1 - fc2)
}

// Dur return the duration of the timecode.
func (c *N25) Dur() time.Duration {
	fc, _ := c.FrameCount()
	return time.Duration(fc) * tN25
}

// String fulfills the Stringer interface
func (c *N25) String() string {
	return c.tc
}
