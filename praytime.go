//--------------------- Copyright Block ----------------------
/*

PrayTime.Go: Prayer Times Calculator (ver 1.0)
Copyright (C) 2007-2010 PrayTimes.org

Go Code By: Adib Aulia Rahmansyah
Original JS Code By: Hamid Zarrabi-Zadeh

License: GNU LGPL v3.0

TERMS OF USE:
	Permission is granted to use this code, with or
	without modification, in any website or application
	provided that credit is given to the original work
	with a link back to PrayTimes.org.

This program is distributed in the hope that it will
be useful, but WITHOUT ANY WARRANTY.

PLEASE DO NOT REMOVE THIS COPYRIGHT BLOCK.


User's Manual:
http://praytimes.org/manual

Calculation Formulas:
http://praytimes.org/calculation


*/

package praytime

var (
	// ---------------------- Global Variables --------------------
	calcMethod     int     // caculation method
	asrJuristic    int     // Juristic method for Asr
	dhuhrMinutes   int     // minutes after mid-day for Dhuhr
	adjustHighLats int     // adjusting method for higher latitudes
	timeFormat     int     // time format
	lat            float64 // latitude
	lng            float64 // longitude
	timeZone       float64 // time-zone
	JDate          float64 // Julian date
	// ------------------------------------------------------------

	// Calculation Methods
	Jafari  int // Ithna Ashari
	Karachi int // University of Islamic Sciences, Karachi
	ISNA    int // Islamic Society of North America (ISNA)
	MWL     int // Muslim World League (MWL)
	Makkah  int // Umm al-Qura, Makkah
	Egypt   int // Egyptian General Authority of Survey
	Custom  int // Custom Setting
	Tehran  int // Institute of Geophysics, University of Tehran

	// Juristic Methods
	Shafii int // Shafii (standard)
	Hanafi int // Hanafi

	// Adjusting Methods for Higher Latitudes
	None       int // No adjustment
	MidNight   int // middle of night
	OneSeventh int // 1/7th of night
	AngleBased int // angle/60th of night

	// Time Formats
	Time24   int // 24-hour format
	Time12   int // 12-hour format
	Time12NS int // 12-hour format with no suffix
	Floating int // floating point number

	// Time Names
	timeNames   []string
	InvalidTime string // The string used for invalid times

	// --------------------- Technical Settings --------------------
	numIterations int // number of iterations needed to compute times

	// ------------------- Calc Method Parameters --------------------
	methodParams map[int][]float64

	/*
	 * methodParams[methodNum] = new Array(fa, ms, mv, is, iv)
	 *
	 * fa : fajr angle ms : maghrib selector (0 = angle 1 = minutes after
	 * sunset) mv : maghrib parameter value (in angle or minutes) is : isha
	 * selector (0 = angle 1 = minutes after maghrib) iv : isha parameter value
	 * (in angle or minutes)
	 */

	prayerTimesCurrent []float64
	offsets            []int
)

func PrayTimeInit() {
	// Initialize vars
	calcMethod = 0
	asrJuristic = 0
	dhuhrMinutes = 0
	adjustHighLats = 1
	timeFormat = 0

	// Calculation Methods
	Jafari = 0  // Ithna Ashari
	Karachi = 1 // University of Islamic Sciences, Karachi
	ISNA = 2    // Islamic Society of North America (ISNA)
	MWL = 3     // Muslim World League (MWL)
	Makkah = 4  // Umm al-Qura, Makkah
	Egypt = 5   // Egyptian General Authority of Survey
	Tehran = 6  // Institute of Geophysics, University of Tehran
	Custom = 7  // Custom Setting

	// Juristic Methods
	Shafii = 0 // Shafii (standard)
	Hanafi = 1 // Hanafi

	// Adjusting Methods for Higher Latitudes
	None = 0       // No adjustment
	MidNight = 1   // middle of night
	OneSeventh = 2 // 1/7th of night
	AngleBased = 3 // angle/60th of night

	// Time Formats
	Time24 = 0   // 24-hour format
	Time12 = 1   // 12-hour format
	Time12NS = 2 // 12-hour format with no suffix
	Floating = 3 // floating point number

	// Time Names

	timeNames = append(timeNames, "Fajr")
	timeNames = append(timeNames, "Sunrise")
	timeNames = append(timeNames, "Dhuhr")
	timeNames = append(timeNames, "Asr")
	timeNames = append(timeNames, "Sunset")
	timeNames = append(timeNames, "Maghrib")
	timeNames = append(timeNames, "Isha")

	InvalidTime = "-----" // The string used for invalid times

	// --------------------- Technical Settings --------------------

	numIterations = 1 // number of iterations needed to compute times

	// ------------------- Calc Method Parameters --------------------

	// Tuning offsets {fajr, sunrise, dhuhr, asr, sunset, maghrib, isha}
	offsets = []int{0, 0, 0, 0, 0, 0, 0}

	//Jafari
	Jvalues := []float64{16, 0, 4, 0, 14}

	//Karachi
	Kvalues := []float64{18, 1, 0, 0, 18}

	//ISNA
	Ivalues := []float64{15, 1, 0, 0, 15}

	//MWL
	MWLvalues := []float64{18, 1, 0, 0, 17}

	//Makkah
	MKvalues := []float64{18.5, 1, 0, 1, 90}

	//Egypt
	Evalues := []float64{19.5, 1, 0, 0, 17.5}

	//Tehran
	Tvalues := []float64{17.7, 0, 4.5, 0, 14}

	//Custom
	Cvalues := []float64{18, 1, 0, 0, 17}

	/*
	 *
	 * fa : fajr angle ms : maghrib selector (0 = angle; 1 = minutes after
	 * sunset) mv : maghrib parameter value (in angle or minutes) is : isha
	 * selector (0 = angle; 1 = minutes after maghrib) iv : isha parameter
	 * value (in angle or minutes)
	 */
	methodParams = map[int][]float64{
		Jafari:  Jvalues,
		Karachi: Kvalues,
		ISNA:    Ivalues,
		MWL:     MWLvalues,
		Makkah:  MKvalues,
		Egypt:   Evalues,
		Tehran:  Tvalues,
		Custom:  Cvalues,
	}
}
