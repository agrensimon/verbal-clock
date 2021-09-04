package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	time, err := time.Parse(time.Kitchen, "12:00AM")
	if err != nil {
		t.Fatal()
	}

	s := timeToString(time)
	assert.Equal(t, "Klockan är tolv + 0", s)
}
