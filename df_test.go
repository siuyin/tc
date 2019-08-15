package df

import "testing"

func TestFrameCount(t *testing.T) {
	dat := []struct {
		i string
		o int
	}{
		{i: "00:10:45:21", o: 19353},
		{i: "00:05:36:12", o: 10082},
		{i: "00:07:16:22", o: 13088},
		{i: "00:05:16:01", o: 9471},
		{i: "00:06:42:01", o: 12049},
	}
	for i, d := range dat {
		o, err := FrameCount(d.i)
		if err != nil {
			t.Error(err)
		}
		if o != d.o {
			t.Errorf("case %d: unexpected value: %d", i, o)
		}
	}
}

func TestFrameCountToDFTimeCode(t *testing.T) {
	dat := []struct {
		o string
		i int
	}{
		{o: "00:10:45;21", i: 19353},
		{o: "00:05:36;12", i: 10082},
		{o: "00:07:16;22", i: 13088},
		{o: "00:05:16;01", i: 9471},
		{o: "00:06:42;01", i: 12049},
	}
	for i, d := range dat {
		o := TimeCode(d.i)
		if o != d.o {
			t.Errorf("case %d: unexpected value: %s", i, o)
		}
	}
}
