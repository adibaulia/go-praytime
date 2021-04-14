package praytime

import (
	"fmt"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	latitude := -37.823689
	longitude := 145.121597
	timezone := 10.0
	// Test Prayer times here
	PrayTimeInit()

	timeFormat = Time24
	calcMethod = Jafari
	asrJuristic = Shafii
	adjustHighLats = AngleBased
	offsets := []int{0, 0, 0, 0, 0, 0, 0}
	Tune(offsets)
	now := time.Now()

	p := getPrayerTimes(now, latitude, longitude, timezone)
	ps := timeNames

	for i := 0; i < len(p); i++ {
		fmt.Println(ps[i], " - ", p[i])
	}

}
