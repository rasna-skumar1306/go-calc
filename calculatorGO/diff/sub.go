package diff

//Diffi returns the difference between each element of []int
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
