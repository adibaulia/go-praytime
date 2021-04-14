package praytime

import (
	"fmt"
	"math"
)

// ---------------------- Compute Prayer Times -----------------------
// compute prayer times at given julian date
func computeTimes(times []float64) []float64 {

	t := dayPortion(times)

	Fajr := computeTime(180-methodParams[calcMethod][0], t[0])

	Sunrise := computeTime(180-0.833, t[1])

	Dhuhr := computeMidDay(t[2])
	Asr := computeAsr(1+float64(asrJuristic), t[3])
	Sunset := computeTime(0.833, t[4])

	Maghrib := computeTime(methodParams[calcMethod][2], t[5])
	Isha := computeTime(
		methodParams[calcMethod][4], t[6])

	CTimes := []float64{Fajr, Sunrise, Dhuhr, Asr, Sunset, Maghrib, Isha}

	return CTimes

}

// compute prayer times at given julian date
func computeDayTimes() []string {
	times := []float64{5, 6, 12, 13, 18, 18, 18} // default times

	for i := 1; i <= numIterations; i++ {
		times = computeTimes(times)
	}

	times = adjustTimes(times)
	times = tuneTimes(times)

	return adjustTimesFormat(times)
}

// adjust times in a prayer time array
func adjustTimes(times []float64) []float64 {
	for i := 0; i < len(times); i++ {
		times[i] += timeZone - lng/15
	}

	times[2] += float64(dhuhrMinutes) / 60 // Dhuhr
	if methodParams[calcMethod][1] == 1 {  // Maghrib
		times[5] = times[4] + methodParams[calcMethod][2]/60
	}
	if methodParams[calcMethod][3] == 1 { // Isha

		times[6] = times[5] + methodParams[calcMethod][4]/60
	}

	if adjustHighLats != None {
		times = adjustHighLatTimes(times)
	}

	return times
}

// convert times array to given time format
func adjustTimesFormat(times []float64) []string {

	result := []string{}

	if timeFormat == Floating {

		for _, time := range times {
			result = append(result, fmt.Sprintf("%f", time))
		}
		return result
	}

	for i := 0; i < 7; i++ {
		if timeFormat == Time12 {
			result = append(result, FloatToTime12(times[i], false))

		} else if timeFormat == Time12NS {
			result = append(result, FloatToTime12(times[i], true))
		} else {
			result = append(result, FloatToTime24(times[i]))
		}
	}
	return result
}

// adjust Fajr, Isha and Maghrib for locations in higher latitudes
func adjustHighLatTimes(times []float64) []float64 {
	nightTime := timeDiff(times[4], times[1]) // sunset to sunrise

	// Adjust Fajr
	FajrDiff := nightPortion(methodParams[calcMethod][0]) * nightTime

	if math.IsNaN(times[0]) || timeDiff(times[0], times[1]) > FajrDiff {
		times[0] = times[1] - FajrDiff
	}

	// Adjust Isha
	var IshaAngle, MaghribAngle float64
	if methodParams[calcMethod][3] == 0 {
		IshaAngle = methodParams[calcMethod][4]
	} else {
		IshaAngle = 18
	}

	IshaDiff := nightPortion(IshaAngle) * nightTime
	if math.IsNaN(times[6]) || timeDiff(times[4], times[6]) > IshaDiff {
		times[6] = times[4] + IshaDiff
	}

	// Adjust Maghrib

	if methodParams[calcMethod][1] == 0 {
		MaghribAngle = methodParams[calcMethod][2]
	} else {
		MaghribAngle = 4
	}

	MaghribDiff := nightPortion(MaghribAngle) * nightTime
	if math.IsNaN(times[5]) || timeDiff(times[4], times[5]) > MaghribDiff {
		times[5] = times[4] + MaghribDiff
	}

	return times
}

func nightPortion(angle float64) float64 {
	var calc float64 = 0

	if adjustHighLats == AngleBased {
		calc = (angle) / 60.0
	} else if adjustHighLats == MidNight {
		calc = 0.5
	} else if adjustHighLats == OneSeventh {
		calc = 0.14286
	}
	return calc
}

// convert hours to day portions
func dayPortion(times []float64) []float64 {
	for i := 0; i < 7; i++ {
		times[i] /= 24
	}
	return times
}

// Tune timings for adjustments
// Set time offsets
func Tune(offsetTimes []int) {

	for i := 0; i < len(offsetTimes); i++ { // offsetTimes length
		// should be 7 in order
		// of Fajr, Sunrise,
		// Dhuhr, Asr, Sunset,
		// Maghrib, Isha
		offsets[i] = offsetTimes[i]
	}
}

func tuneTimes(times []float64) []float64 {
	for i := 0; i < len(times); i++ {
		times[i] = times[i] + float64(offsets[i])/60.0
	}

	return times
}
