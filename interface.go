package praytime

import (
	"fmt"
	"math"
	"time"
)

// -------------------- Interface Functions --------------------
// return prayer times for a given date
func getDatePrayerTimes(year, month, day int, latitude, longitude, tZone float64) []string {
	lat = latitude
	lng = longitude
	timeZone = tZone
	JDate = julianDate(year, month, day)
	lonDiff := longitude / (15.0 * 24.0)
	JDate = JDate - lonDiff
	return computeDayTimes()
}

// return prayer times for a given date
func getPrayerTimes(date time.Time, latitude, longitude, tZone float64) []string {

	year, month, day := date.Date()

	return getDatePrayerTimes(year, int(month)+1, day, latitude, longitude, tZone)
}

// set custom values for calculation parameters
func setCustomParams(params []float64) {

	for i := 0; i < 5; i++ {
		if params[i] == -1 {
			params[i] = methodParams[calcMethod][i]
			methodParams[Custom] = params
		} else {
			methodParams[Custom][i] = params[i]
		}
	}
	calcMethod = Custom
}

//SetFajrAngle set the angle for calculating Fajr
func SetFajrAngle(angle float64) {
	params := []float64{angle, -1, -1, -1, -1}
	setCustomParams(params)
}

//SetMaghribAngle set the angle for calculating Maghrib
func SetMaghribAngle(angle float64) {
	params := []float64{-1, 0, angle, -1, -1}
	setCustomParams(params)
}

//SetIshaAngle set the angle for calculating Isha
func SetIshaAngle(angle float64) {
	params := []float64{-1, -1, -1, 0, angle}
	setCustomParams(params)
}

//SetMaghribMinutes set the minutes after Sunset for calculating Maghrib
func SetMaghribMinutes(minutes float64) {
	params := []float64{-1, 1, minutes, -1, -1}
	setCustomParams(params)
}

//SetIshaMinutes set the minutes after Maghrib for calculating Isha
func SetIshaMinutes(minutes float64) {
	params := []float64{-1, -1, -1, 1, minutes}
	setCustomParams(params)
}

//FloatToTime24 convert double hours to 24h format
func FloatToTime24(time float64) string {

	var result string

	if math.IsNaN(time) {
		return InvalidTime
	}

	time = fixhour(time + 0.5/60.0) // add 0.5 minutes to round
	hours := int(math.Floor(time))
	minutes := math.Floor((time - float64(hours)) * 60.0)

	if (hours >= 0 && hours <= 9) && (minutes >= 0 && minutes <= 9) {
		result = fmt.Sprint("0", hours, ":0", math.Round(minutes))
	} else if hours >= 0 && hours <= 9 {
		result = fmt.Sprint("0", hours, ":", math.Round(minutes))
	} else if minutes >= 0 && minutes <= 9 {
		result = fmt.Sprint(hours, ":0", math.Round(minutes))
	} else {
		result = fmt.Sprint(hours, ":", math.Round(minutes))
	}
	return result
}

// convert double hours to 12h format
func FloatToTime12(time float64, noSuffix bool) string {

	if math.IsNaN(time) {
		return InvalidTime
	}

	time = fixhour(time + 0.5/60) // add 0.5 minutes to round
	hours := int(math.Floor(time))
	minutes := math.Floor((time - float64(hours)) * 60)
	var suffix, result string
	if hours >= 12 {
		suffix = "pm"
	} else {
		suffix = "am"
	}
	hours = ((((hours + 12) - 1) % (12)) + 1)
	/*hours = (hours + 12) - 1;
	  int hrs = (int) hours % 12;
	  hrs += 1;*/
	if noSuffix == false {
		if (hours >= 0 && hours <= 9) && (minutes >= 0 && minutes <= 9) {
			result = fmt.Sprint("0", hours, ":0", math.Round(minutes), " ", suffix)
		} else if hours >= 0 && hours <= 9 {
			result = fmt.Sprint("0", hours, ":", math.Round(minutes), " ", suffix)
		} else if minutes >= 0 && minutes <= 9 {
			result = fmt.Sprint(hours, ":0", math.Round(minutes), " ", suffix)
		} else {
			result = fmt.Sprint(hours, ":", math.Round(minutes), " ", suffix)
		}

	} else {
		if (hours >= 0 && hours <= 9) && (minutes >= 0 && minutes <= 9) {
			result = fmt.Sprint("0", hours, ":0", math.Round(minutes))
		} else if hours >= 0 && hours <= 9 {
			result = fmt.Sprint("0", hours, ":", math.Round(minutes))
		} else if minutes >= 0 && minutes <= 9 {
			result = fmt.Sprint(hours, ":0", math.Round(minutes))
		} else {
			result = fmt.Sprint(hours, ":", math.Round(minutes))
		}
	}
	return result

}

// convert double hours to 12h format with no suffix
func FloatToTime12NS(time float64) string {
	return FloatToTime12(time, true)
}
