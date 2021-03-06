//Package add has functions that can sum []int, []float32 and []float64
//elements and return the values
package add

//Addi adds up all the integer values sent as argument
//and returns the sum.
func Addi(xi ...int) int {
	s := 0
	if xi == nil {
		return 0
	}
	if len(xi) == 1 {
		return xi[0]
	}
	for _, v := range xi {
		s += v
	}
	return s
}

//Addf adds up all the float64 values sent as argument
//and returns the sum.
func Addf(xf ...float64) float64 {
	var s float64
	if xf == nil {
		return 0
	}
	if len(xf) == 1 {
		return xf[0]
	}
	for _, v := range xf {
		s += v
	}
	return s
}

//Addf32 adds up all the float32 values sent as argument
//and returns the sum.
func Addf32(xf ...float32) float32 {
	var s float32
	if xf == nil {
		return 0
	}
	if len(xf) == 1 {
		return xf[0]
	}
	for _, v := range xf {
		s += v
	}
	return s
}
