package main

import (
	"fmt"
	"time"
)

// fem över
// tio över
// kvar över
// tjugo över

// fem i halv 		(tjugofem över)
// halv
// fem över halv	(tjugofem i)

// tjugo i
// kvart i
// tio i
// fem i

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
	if m >= 30 {
		h += 1
	}

	fives := m / 5 * 5

	var preposition string
	switch {
	case fives == 0:
		return fmt.Sprintf("Klockan är %v + %v", hText[h], m%5)
	case fives == 30:
		return fmt.Sprintf("Klockan är %v %v + %v", mText[fives], hText[h], m%5)
	case fives > 30:
		preposition = "i"
	case fives < 30:
		preposition = "över"
	}

	return fmt.Sprintf("Klockan är %v %v %v + %v", mText[fives], preposition, hText[h], m%5)
}

var mText = map[int]string{
	5:  "fem",
	10: "tio",
	15: "kvart",
	20: "tjugo",
	25: "tjugofem",
	30: "halv",
	35: "tjugofem",
	40: "tjugo",
	45: "kvart",
	50: "tio",
	55: "fem",
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
