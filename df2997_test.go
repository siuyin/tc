package tc

import "testing"

func TestFrameCountDF29_97(t *testing.T) {
	dat := []struct {
		i string
		o int
	}{
		{i: "00:10:45:21", o: 19353},
		{i: "00:05:36:12", o: 10082},
		{i: "00:07:16:22", o: 13088},
		{i: "00:05:16:01", o: 9471},
		{i: "00:06:42:01", o: 12049},
		{i: "01:00:00;00", o: 107892},
		{i: "02:00:00;00", o: 215784},
		{i: "10:00:00;00", o: 1078920},
		{i: "20:00:00;00", o: 2157840},
		{i: "30:00:00;00", o: 3236760},
	}
	for i, d := range dat {
		tc := NewDF29_97(d.i)
		o, err := tc.FrameCount()
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
		{o: "01:00:00;00", i: 107892},
		{o: "02:00:00;00", i: 215784},
		{o: "10:00:00;00", i: 1078920},
		{o: "20:00:00;00", i: 2157840},
		{o: "30:00:00;00", i: 3236760},
	}
	for i, d := range dat {
		o := NewDF29_97FrameCount(d.i)
		if o.String() != d.o {
			t.Errorf("case %d: unexpected value: %s", i, o)
		}
	}
}

func TestSub(t *testing.T) {
	dat := []struct {
		a, b string
		o    string
	}{
		{"10:10:45:21", "10:00:00:00", "00:10:45:21"},
		{"10:16:32:11", "10:10:55:27", "00:05:36:12"},
		{"10:23:59:19", "10:16:42:29", "00:07:16:22"},
		{"10:29:26:01", "10:24:10:00", "00:05:16:01"},
		{"10:36:18:01", "10:29:36:00", "00:06:42:01"},
		{"10:44:20:00", "10:36:28:00", "00:07:52:00"},
	}
	for i, d := range dat {
		o := NewDF29_97(d.a).Sub(NewDF29_97(d.b))
		if o.String() != NewDF29_97(d.o).String() {
			t.Errorf("case %d: Unexpected value: %s", i, o)
		}
	}
}

func TestAddDF29_97(t *testing.T) {
	t1 := NewDF29_97("10:00:00:00")
	t2 := NewDF29_97("10:00:00:00")
	t3 := t1.Add(t2)
	if t3.String() != "20:00:00;00" {
		t.Errorf("unexpected value: %s", t3)
	}
}
func TestOverflowF29_97(t *testing.T) {
	t1 := NewDF29_97("40:00:00:00")
	t2 := NewDF29_97("80:00:00:00")
	t3 := t1.Add(t2)
	if t3.String() != "120:00:00;00" {
		t.Errorf("unexpected value: %s", t3)
	}
}
