//Package diff has func Diffi which returns difference between
//two consecutive elements in []int, []float32 and []float64
package diff

//Diffi calculates the difference between each consecutive element of []int
//and returns the value as []int of difference elements
func Diffi(xi ...int) []int {
	s := make([]int, 15)
	j := 1
	for i := 0; i < len(xi); i++ {
		s[i] = 0
	}
	for i := 0; i < len(xi); i++ {
		s[i] = xi[i] - xi[j]
		j++
		if j == len(xi) {
			break
		}
	}
	return s[:len(xi)-1]
}

//Difff32 calculates the difference between each consecutive element of []float32
//and returns the value as []float32 of difference elements
func Difff32(xi ...float32) []float32 {
	s := make([]float32, 15)
	j := 1
	for i := 0; i < len(xi); i++ {
		s[i] = 0
	}
	for i := 0; i < len(xi); i++ {
		s[i] = xi[i] - xi[j]
		j++
		if j == len(xi) {
			break
		}
	}
	return s[:len(xi)-1]
}

//Difff64 calculates the difference between each consecutive element of []float64
//and returns the value as []float64 of difference elements
func Difff64(xi ...float64) []float64 {
	s := make([]float64, 15)
	j := 1
	for i := 0; i < len(xi); i++ {
		s[i] = 0
	}
	for i := 0; i < len(xi); i++ {
		s[i] = xi[i] - xi[j]
		j++
		if j == len(xi) {
			break
		}
	}
	return s[:len(xi)-1]
}
