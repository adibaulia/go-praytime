package praytime

import "math"

// ---------------------- Trigonometric Functions -----------------------
//fixangle range reduce angle in degrees.
func fixangle(a float64) float64 {

	a = a - (360 * (math.Floor(a / 360.0)))

	if a < 0 {
		a = a + 24
	}

	return a
}

//fixhour range reduce hours to 0..23
func fixhour(a float64) float64 {
	a = a - 24.0*math.Floor(a/24.0)
	if a < 0 {
		a = a + 24
	}
	return a
}

//radiansToDegrees radian to degree
func radiansToDegrees(alpha float64) float64 {
	return ((alpha * 180.0) / math.Pi)
}

//degreesToRadians degree to radian
func degreesToRadians(alpha float64) float64 {
	return ((alpha * math.Pi) / 180.0)
}

//dsin degree sin
func dsin(d float64) float64 {
	return (math.Sin(degreesToRadians(d)))
}

// degree cos
func dcos(d float64) float64 {
	return (math.Cos(degreesToRadians(d)))
}

//dtan degree tan
func dtan(d float64) float64 {
	return (math.Tan(degreesToRadians(d)))
}

// degree arcsin
func darcsin(x float64) float64 {
	val := math.Asin(x)
	return radiansToDegrees(val)
}

// degree arccos
func darccos(x float64) float64 {
	val := math.Acos(x)
	return radiansToDegrees(val)
}

// degree arctan
func darctan(x float64) float64 {
	val := math.Atan(x)
	return radiansToDegrees(val)
}

// degree arctan2
func darctan2(y, x float64) float64 {
	val := math.Atan2(y, x)
	return radiansToDegrees(val)
}

// degree arccot
func darccot(x float64) float64 {
	val := math.Atan2(1.0, x)
	return radiansToDegrees(val)
}
