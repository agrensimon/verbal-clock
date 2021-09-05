package main

import (
	"fmt"
	"time"
)

const (
	hourPinOffset = 12
)

func main() {
	// check the time
	// round down to nearest
	now := time.Now()
	s := timeToString(now)

	fmt.Println(s)

}

func enumerate(t time.Time) (h, m int) {
	h = t.Hour()
	minutes := t.Minute()

	// 12 hour clock
	if h > 11 {
		h -= 12
	}

	// refer to next hour for all "to" times
	if minutes >= 25 {
		h = (h + 1) % 12
	}

	m = minutes / 5
	return
}

func timeToString(t time.Time) string {
	hours, fives := enumerate(t)

	if fives == 0 {
		return fmt.Sprintf("%v", hText[hours])
	}

	return fmt.Sprintf("%v %v", mText[fives], hText[hours])
}

func timeToPinOut(t time.Time) (out int64) {
	h, fives := enumerate(t)

	for _, elem := range mPinOut[fives] {
		out = out | 1<<(elem)
	}

	out = out | 1<<(h+hourPinOffset)

	return

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
}

var mPinOut = map[int][]int{
	1:  {2, 7},    // fem över
	2:  {3, 7},    // tio över
	3:  {4, 7},    // kvart över
	4:  {5, 7},    // tjugo över
	5:  {2, 8, 6}, // fem i halv
	6:  {6},       // halv
	7:  {2, 7, 6}, // fem över halv
	8:  {5, 11},   // tjugo i
	9:  {4, 10},   //kvart i
	10: {3, 9},    // tio i
	11: {2, 8},    // fem i
}
