package tc

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// DF29_97 represents a dropframe timecode
type DF29_97 struct {
	tc string
}

// NewDF29_97 returns a reference to a new dropframe timecode.
func NewDF29_97(tc string) *DF29_97 {
	if !validTC(tc) {
		return &DF29_97{"00:00:00;00"}
	}
	l := len(tc)
	tc = tc[0:l-3] + ";" + tc[l-2:l]
	//tc = tc[0:8] + ";" + tc[9:11]
	return &DF29_97{tc}
}

func (c *DF29_97) String() string {
	return c.tc
}

// DropFrameCount returns the number of frames at 30 fps given a
// dropframe timecode tc.
func (c *DF29_97) DropFrameCount() (int, error) {
	c1 := strings.Index(c.tc, ":")
	hh, err := strconv.Atoi(c.tc[0:c1])
	if err != nil {
		return 0, err
	}
	mm, err := strconv.Atoi(c.tc[c1+1 : c1+3])
	if err != nil {
		return 0, err
	}
	ss, err := strconv.Atoi(c.tc[c1+4 : c1+6])
	if err != nil {
		return 0, err
	}
	ff, err := strconv.Atoi(c.tc[c1+7 : c1+9])
	if err != nil {
		return 0, err
	}
	// hh, err := strconv.Atoi(c.tc[0:2])
	// if err != nil {
	// 	return 0, err
	// }
	// mm, err := strconv.Atoi(c.tc[3:5])
	// if err != nil {
	// 	return 0, err
	// }
	// ss, err := strconv.Atoi(c.tc[6:8])
	// if err != nil {
	// 	return 0, err
	// }
	// ff, err := strconv.Atoi(c.tc[9:11])
	// if err != nil {
	// 	return 0, err
	// }
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

// NewDF29_97FrameCount returns a dropframe timecode from frame count.
func NewDF29_97FrameCount(fc int) *DF29_97 {
	const (
		fc10Mins = 10 * 60 * 30000 / 1001
		fc1Min   = 1 * 60 * 30000 / 1001
	)
	tenMins := fc / fc10Mins
	mins := fc / fc1Min
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
