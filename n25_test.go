package tc

import (
	"testing"
	"time"
)

func TestFrameCountN25(t *testing.T) {
	dat := []struct {
		i string
		o int
	}{
		{i: "00:10:45:21", o: 16146},
	}
	for i, d := range dat {
		tc := NewN25(d.i)
		o, err := tc.FrameCount()
		if err != nil {
			t.Error(err)
		}
		if o != d.o {
			t.Errorf("case %d: unexpected value: %d", i, o)
		}
	}
}
func TestSubN25(t *testing.T) {
	dat := []struct {
		a, b string
		o    string
	}{
		{"10:10:45:00", "00:00:00:01", "10:10:44:24"},
		{"10:10:45:01", "00:00:00:01", "10:10:45:00"},
	}
	for i, d := range dat {
		o := NewN25(d.a).Sub(NewN25(d.b))
		if o.String() != NewN25(d.o).String() {
			t.Errorf("case %d: Unexpected value: %s", i, o)
		}
	}
}

func TestDurN25(t *testing.T) {
	dat := []struct {
		a string
		o time.Duration
	}{
		{"12:34:56:24", (12*3600*25 + 34*60*25 + 56*25 + 24) * tN25},
	}
	for i, d := range dat {
		o := NewN25(d.a)
		if o.Dur() != d.o {
			t.Errorf("case %d: Unexpected value: %d", i, o.Dur())
		}
	}
}
