package add

//Addi adds up all the integer values sent as argument
//and returns the sum.
func Addi(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}

//Addf adds up all the float64 values sent as argument
//and returns the sum.
func Addf(xf ...float64) float64{
	var s float64
	for _, v := range xf {
		s += v
	}
	return s
}

