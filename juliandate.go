package praytime

import (
	"math"
	"time"
)

// ---------------------- Julian Date Functions -----------------------
// calculate julian date from a calendar date
func julianDate(year, month, day int) float64 {

	if month <= 2 {
		year -= 1
		month += 12
	}
	A := math.Floor(float64(year) / 100.0)

	B := 2 - A + math.Floor(A/4.0)

	JD := math.Floor(365.25*(float64(year)+4716)) + math.Floor(30.6001*(float64(month)+1)) + float64(day) + B - 1524.5

	return JD
}

// convert a calendar date to julian date (second method)
func calcJD(year, month, day int) float64 {
	J1970 := 2440588.0
	date := time.Date(year, time.Month(month-1), day, 0, 0, 0, 0, time.Local)

	ms := float64(date.UnixNano() / int64(time.Millisecond)) // # of milliseconds since midnight Jan 1, 1970
	days := math.Floor(ms / (1000.0 * 60.0 * 60.0 * 24.0))
	return J1970 + days - 0.5

}
