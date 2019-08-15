// Package tc provides timecode (including 29.97 dropframe) calculations.
package tc

import (
	"fmt"
	"strconv"
	"time"
)

const (
	tDF29_97 = time.Duration(1001 * 1000000000 / 30000)
	tN25     = time.Duration(1000000000 / 25)
)

type TC interface {
	Add(tc TC) TC
	Sub(tc TC) TC
	Dur() time.Duration
	FrameCount() (int, error)
	String() string
}

func validTC(tc string) bool {
	if len(tc) != 11 {
		return false
	}
	return true
}

// N25 represent 25 fps Non-drop frame (ie. PAL rate)
type N25 struct {
	tc string
}

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

// DF29_97 represents a dropframe timecode
type DF29_97 struct {
	tc string
}

// NewTC returns a reference to a new dropframe timecode.
func NewDF29_97(tc string) *DF29_97 {
	if !validTC(tc) {
		return &DF29_97{"00:00:00;00"}
	}
	tc = tc[0:8] + ";" + tc[9:11]
	return &DF29_97{tc}
}

func (c *DF29_97) String() string {
	return c.tc
}

// DropFrameCount returns the number of frames at 30 fps given a
// dropframe timecode tc.
func (c *DF29_97) DropFrameCount() (int, error) {
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
	return hh*3600*30 + mm*60*30 + ss*30 + ff, nil
}

// FrameCount returns the number of frames at 29.97 (30000/1001)
// given a dropframe timecode tc.
func (c *DF29_97) FrameCount() (int, error) {
	dfc, err := c.DropFrameCount()
	if err != nil {
		return 0, err
	}
	mins := dfc / (60 * 30)
	tenMins := dfc / (10 * 60 * 30)
	return dfc - mins*2 + tenMins*2, nil
}

// NewTCFrameCount returns a dropframe timecode from frame count.
func NewDF29_97FrameCount(fc int) *DF29_97 {
	mins := fc / (60 * 30)
	tenMins := fc / (10 * 60 * 30)
	dfc := fc + mins*2 - tenMins*2
	hh := dfc / (3600 * 30)
	mm := (dfc - hh*3600*30) / (60 * 30)
	ss := (dfc - hh*3600*30 - mm*60*30) / 30
	ff := dfc - hh*3600*30 - mm*60*30 - ss*30
	return NewDF29_97(fmt.Sprintf("%02d:%02d:%02d;%02d", hh, mm, ss, ff))
}

// Dur returns the duration of the timecode.
func (c *DF29_97) Dur() time.Duration {
	fc, _ := c.FrameCount()
	return time.Duration(fc) * tDF29_97
}

// Add sums timecodes c+tc.
func (c *DF29_97) Add(tc TC) TC {
	fc1, _ := c.FrameCount()
	fc2, _ := tc.FrameCount()
	return NewDF29_97FrameCount(fc1 + fc2)
}

// Sub returns c-tc as timecode.
func (c *DF29_97) Sub(tc TC) TC {
	fc1, _ := c.FrameCount()
	fc2, _ := tc.FrameCount()
	return NewDF29_97FrameCount(fc1 - fc2)
}
