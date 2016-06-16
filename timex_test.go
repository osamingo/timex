package timex

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNowTruncate(t *testing.T) {

	now := time.Now()
	tm := NowTruncate(time.Hour)

	assert.True(t, tm.Before(now))
	assert.Empty(t, tm.Nanosecond())
	assert.Empty(t, tm.Second())
	assert.Empty(t, tm.Minute())
}

func TestBetween(t *testing.T) {

	now := time.Now()
	after := now.Add(Week)
	before := now.Add(-Week)

	assert.True(t, Between(now, before, after))
	assert.False(t, Between(now, after, before))
}
