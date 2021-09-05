package main

import (
	"fmt"
	"time"
)

func main() {
	// check the time
	// round down to nearest
	now := time.Now()
	s := timeToString(now)

	fmt.Println(s)

}

func timeToString(t time.Time) string {
	h := t.Hour()
	m := t.Minute()

	// 12 hour clock
	if h > 11 {
		h -= 12
	}

	// refer to next hour for all "to" times
	if m >= 25 {
		h += 1
	}

	fives := m / 5

	fmt.Println(fives)

	if fives == 0 {
		return fmt.Sprintf("%v", hText[h])
	}

	return fmt.Sprintf("%v %v", mText[fives], hText[h])
}

var mText = map[int]string{
	1:  "fem över",
	2:  "tio över",
	3:  "kvart över",
	4:  "tjugo över",
	5:  "fem i halv",
	6:  "halv",
	7:  "fem över halv",
	8:  "tjugo i",
	9:  "kvart i",
	10: "tio i",
	11: "fem i",
}

var hText = map[int]string{
	0:  "tolv",
	1:  "ett",
	2:  "två",
	3:  "tre",
	4:  "fyra",
	5:  "fem",
	6:  "sex",
	7:  "sju",
	8:  "åtta",
	9:  "nio",
	10: "tio",
	11: "elva",
	12: "tolv",
}
