//Package diff has func Diffi which returns difference between
//two consecutive elements in []int, []float32 and []float64
package diff

import (
	"errors"
)

//Diffi calculates the difference between each consecutive element of []int
//and returns the value as []int of difference elements
func Diffi(xi ...int) ([]int, error) {
	if len(xi) == 0 {
		return []int{0}, errors.New("no values provided expected at least one element")
	}
	if len(xi) == 1 {
		return xi, nil
	}
	s := make([]int, len(xi)-1)
	for i := 0; i < len(xi)-1; i++ {
		s[i] = xi[i] - xi[i+1]
	}
	return s[:len(xi)-1], nil
}

//Difff32 calculates the difference between each consecutive element of []float32
//and returns the value as []float32 of difference elements
func Difff32(xi ...float32) ([]float32, error) {
	s := make([]float32, 15)
	if len(xi) == 0 {
		return []float32{0}, errors.New("The array is empty expect atleast one element")
	}
	if len(xi) == 1 {
		return xi, nil
	}
	for i := 0; i < len(xi)-1; i++ {
		s[i] = xi[i] - xi[i+1]
	}
	return s[:len(xi)-1], nil
}

//Difff64 calculates the difference between each consecutive element of []float64
//and returns the value as []float64 of difference elements
func Difff64(xi ...float64) ([]float64, error) {
	s := make([]float64, 15)
	if len(xi) == 0 {
		return []float64{0}, errors.New("The array is empty expect atleast one element")
	}
	if len(xi) == 1 {
		return xi, nil
	}
	for i := 0; i < len(xi)-1; i++ {
		s[i] = xi[i] - xi[i+1]
	}
	return s[:len(xi)-1], nil
}
