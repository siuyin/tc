// Package df provides NTSC dropframe calculations.
package df

import (
	"fmt"
	"strconv"
)

func validTC(tc string) bool {
	if len(tc) != 11 {
		return false
	}
	return true
}

// DropFrameCount returns the number of frames at 30 fps given a
// dropframe timecode tc.
func DropFrameCount(tc string) (int, error) {
	hh, err := strconv.Atoi(tc[0:2])
	if err != nil {
		return 0, err
	}
	mm, err := strconv.Atoi(tc[3:5])
	if err != nil {
		return 0, err
	}
	ss, err := strconv.Atoi(tc[6:8])
	if err != nil {
		return 0, err
	}
	ff, err := strconv.Atoi(tc[9:11])
	if err != nil {
		return 0, err
	}
	return hh*3600*30 + mm*60*30 + ss*30 + ff, nil
}

// FrameCount returns the number of frames at 29.97 (30000/1001)
// given a dropframe timecode tc.
func FrameCount(tc string) (int, error) {
	dfc, err := DropFrameCount(tc)
	if err != nil {
		return 0, err
	}
	mins := dfc / (60 * 30)
	tenMins := dfc / (10 * 60 * 30)
	return dfc - mins*2 + tenMins*2, nil
}

// TimeCode returns a dropframe timecode in the format
// hh:mm:ss;ff .
func TimeCode(fc int) string {
	mins := fc / (60 * 30)
	tenMins := fc / (10 * 60 * 30)
	dfc := fc + mins*2 - tenMins*2
	hh := dfc / (3600 * 30)
	mm := (dfc - hh*3600*30) / (60 * 30)
	ss := (dfc - hh*3600*30 - mm*60*30) / 30
	ff := dfc - hh*3600*30 - mm*60*30 - ss*30
	return fmt.Sprintf("%02d:%02d:%02d;%02d", hh, mm, ss, ff)
}
