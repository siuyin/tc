package tc

import (
	"fmt"
)

func ExampleTC_df2997() {
	var a TC
	a = NewDF29_97("00:00:10:02")
	var b TC
	b = NewDF29_97("01:00:00:00")
	fmt.Println(a.Dur())
	fmt.Println(a.Add(b))
}
func ExampleTC_n25() {
	var a TC
	a = NewN25("00:00:10:02")
	var b TC
	b = NewN25("01:00:00:00")
	fmt.Println(a.Dur())
	fmt.Println(a.Add(b))
}
