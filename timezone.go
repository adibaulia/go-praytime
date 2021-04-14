package praytime

import "time"

// ---------------------- Time-Zone Functions -----------------------
// compute local time-zone for a specific date
func getTimeZone() float64 {
	location, _ := time.LoadLocation("Local")
	_, tzOffset := time.Now().In(location).Zone()
	hoursDiff := (tzOffset / 1000.0) / 3600
	return float64(hoursDiff)
}

// ---------------------- Misc Functions -----------------------
// compute the difference between two times
func timeDiff(time1, time2 float64) float64 {
	return fixhour(time2 - time1)
}

// detect daylight saving in a given date
// func detectDaylightSaving() float64 {
// 	location, _ := time.LoadLocation("Local")
// 	_, tzOffset := time.Now().In(location).Zone()
// 	hoursDiff := timez.getDSTSavings()
// 	return hoursDiff
// }
