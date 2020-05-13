//Package sub contains functions that can find the difference
//of total no of elements in []int, []float32 and float64
package sub

//Subi returns the integer difference of total number of
//integers passed as arguement
func Subi(xi ...int) int {
	d := 0
	if xi == nil {
		return 0
	}
	if len(xi) == 1 {
		return xi[0]
	}
	for _, v := range xi {
		d -= v
	}
	return d
}

//Subf returns the float64 difference of total number of float64
//decimals passed as arguments.
func Subf(xf ...float64) float64 {
	var d float64
	if xf == nil {
		return 0.0
	}
	if len(xf) == 1 {
		return xf[0]
	}
	for _, v := range xf {
		d -= v
	}
	return d
}

//Subf32 returns the float32 difference of total number of float32
//decimals passed as arguments.
func Subf32(xf ...float32) float32 {
	var d float32
	if xf == nil {
		return 0.0
	}
	if len(xf) == 1 {
		return xf[0]
	}
	for _, v := range xf {
		d -= v
	}
	return d
}
