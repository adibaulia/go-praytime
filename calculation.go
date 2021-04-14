package praytime

import "math"

// ---------------------- Calculation Functions -----------------------
// References:
// http://www.ummah.net/astronomy/saltime
// http://aa.usno.navy.mil/faq/docs/SunApprox.html
// compute declination angle of sun and equation of time
func sunPosition(jd float64) []float64 {

	D := jd - 2451545
	g := fixangle(357.529 + 0.98560028*D)
	q := fixangle(280.459 + 0.98564736*D)
	L := fixangle(q + (1.915 * dsin(g)) + (0.020 * dsin(2*g)))

	// float64 R = 1.00014 - 0.01671 * [self dcos:g] - 0.00014 * [self dcos:
	// (2*g)];
	e := 23.439 - (0.00000036 * D)
	d := darcsin(dsin(e) * dsin(L))
	RA := (darctan2((dcos(e) * dsin(L)), (dcos(L)))) / 15.0
	RA = fixhour(RA)
	EqT := q/15.0 - RA
	sPosition := []float64{d, EqT}
	return sPosition
}

// compute equation of time
func equationOfTime(jd float64) float64 {
	eq := sunPosition(jd)[1]
	return eq
}

// compute declination angle of sun
func sunDeclination(jd float64) float64 {
	d := sunPosition(jd)[0]
	return d
}

// compute mid-day (Dhuhr, Zawal) time
func computeMidDay(t float64) float64 {
	T := equationOfTime(JDate + t)
	Z := fixhour(12 - T)
	return Z
}

// compute time for a given angle G
func computeTime(G, t float64) float64 {

	D := sunDeclination(JDate + t)
	Z := computeMidDay(t)
	Beg := -dsin(G) - dsin(D)*dsin(lat)
	Mid := dcos(D) * dcos(lat)
	V := darccos(Beg/Mid) / 15.0

	if G > 90 {
		V = -V
	}

	return Z + V
}

// compute the time of Asr
// Shafii: step=1, Hanafi: step=2
func computeAsr(step, t float64) float64 {
	D := sunDeclination(JDate + t)
	G := -darccot(step + dtan(math.Abs(lat-D)))
	return computeTime(G, t)
}
