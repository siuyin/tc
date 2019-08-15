// Package df provides NTSC dropframe calculations.
package df

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func validTC(tc string) bool {
	if len(tc) != 11 {
		return false
	}
	return true
}

// TC represents a dropframe timecode
type TC struct {
	tc string
}

// NewTC returns a reference to a new dropframe timecode.
func NewTC(tc string) *TC {
	if !validTC(tc) {
		return &TC{"00:00:00;00"}
	}
	tc = tc[0:8] + ";" + tc[9:11]
	return &TC{tc}
}

func (c *TC) String() string {
	return c.tc
}

// DropFrameCount returns the number of frames at 30 fps given a
// dropframe timecode tc.
func (c *TC) DropFrameCount() (int, error) {
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
func (c *TC) FrameCount() (int, error) {
	dfc, err := c.DropFrameCount()
	if err != nil {
		return 0, err
	}
	mins := dfc / (60 * 30)
	tenMins := dfc / (10 * 60 * 30)
	return dfc - mins*2 + tenMins*2, nil
}

// NewTCFrameCount returns a dropframe timecode from frame count.
func NewTCFrameCount(fc int) *TC {
	mins := fc / (60 * 30)
	tenMins := fc / (10 * 60 * 30)
	dfc := fc + mins*2 - tenMins*2
	hh := dfc / (3600 * 30)
	mm := (dfc - hh*3600*30) / (60 * 30)
	ss := (dfc - hh*3600*30 - mm*60*30) / 30
	ff := dfc - hh*3600*30 - mm*60*30 - ss*30
	return NewTC(fmt.Sprintf("%02d:%02d:%02d;%02d", hh, mm, ss, ff))
}

var frameTimeNS = int(math.Trunc(float64(1001) / float64(30000) * float64(1000000000)))

// Dur returns the duration of the timecode.
func (c *TC) Dur() time.Duration {
	fc, _ := c.FrameCount()
	return time.Duration(fc * frameTimeNS)
}

// Add sums timecodes c+tc.
func (c *TC) Add(tc *TC) *TC {
	fc := (c.Dur() + tc.Dur()) / time.Duration(frameTimeNS)
	return NewTCFrameCount(int(fc))
}

// Sub returns c-tc as timecode.
func (c *TC) Sub(tc *TC) *TC {
	fc := (c.Dur() - tc.Dur()) / time.Duration(frameTimeNS)
	return NewTCFrameCount(int(fc))
}
