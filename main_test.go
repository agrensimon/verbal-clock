package main

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeToString(t *testing.T) {

	type test struct {
		input string // time i time.Kitchen format
		want  string
	}
	tests := []test{
		{input: "00:00AM", want: "tolv"},
		{input: "00:06AM", want: "fem över tolv"},
		{input: "00:12AM", want: "tio över tolv"},
		{input: "00:23AM", want: "tjugo över tolv"},
		{input: "00:29AM", want: "fem i halv ett"},
		{input: "00:30AM", want: "halv ett"},
		{input: "00:36AM", want: "fem över halv ett"},
		{input: "00:42AM", want: "tjugo i ett"},
		{input: "00:53AM", want: "tio i ett"},
		{input: "00:59AM", want: "fem i ett"},
		{input: "11:59PM", want: "fem i tolv"},
	}

	for _, tc := range tests {
		time, err := time.Parse(time.Kitchen, tc.input)
		if err != nil {
			t.Fatal()
		}

		s := timeToString(time)
		assert.Equal(t, tc.want, s)
	}

	for _, tc := range tests {
		time1, err := time.Parse(time.Kitchen, tc.input)
		if err != nil {
			t.Fatal()
		}

		s1 := timeToString(time1)

		time2, err := time.Parse(time.Kitchen, tc.input[:len(tc.input)-2]+"PM")
		if err != nil {
			t.Fatal()
		}

		s2 := timeToString(time2)

		assert.Equal(t, s1, s2)
	}
}

func TestTimeToPint(t *testing.T) {
	type test struct {
		input string // time i time.Kitchen format
		want  string
	}
	tests := []test{
		{input: "00:00AM", want: "1000000000000"},
		{input: "00:06AM", want: "1000010000100"},
		{input: "00:12AM", want: "1000010001000"},
		{input: "00:23AM", want: "1000010100000"},
		{input: "00:29AM", want: "10000101000100"},
		{input: "00:30AM", want: "10000001000000"},
		{input: "00:36AM", want: "10000011000100"},
		{input: "00:42AM", want: "10100000100000"},
		{input: "00:53AM", want: "10001000001000"},
		{input: "00:59AM", want: "10000100000100"},
	}

	for _, tc := range tests {
		time, err := time.Parse(time.Kitchen, tc.input)
		if err != nil {
			t.Fatal()
		}

		pins := timeToPinOut(time)
		pinsString := strconv.FormatInt(pins, 2)
		if err != nil {
			t.Fatal()
		}

		assert.Equal(t, tc.want, pinsString)

	}
}
