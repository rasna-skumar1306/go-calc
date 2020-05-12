package add

import (
	"fmt"
	"testing"
)

func TestAddi(t *testing.T) {
	n := Addi([]int{2, 3, 5, 7}...)
	if n != 17 {
		t.Error("got", n, "expected", 17)
	}
}

func TestAddf(t *testing.T) {
	n := Addf([]float64{2.1, 3.2, 5.3, 7.4, 8.5}...)
	if n != 26.5 {
		t.Error("got", n, "expected", 26.5)
	}
}

func ExampleAddi() {
	fmt.Println(Addi([]int{2, 3, 5, 7}...))
	//Output:
	//17
}
func ExampleAddf() {
	fmt.Println(Addf([]float64{2.1, 3.2, 5.3, 7.4, 8.5}...))
	//Output:
	//26.5
}

func BenchmarkAddi(b *testing.B) {
	c := []int{2, 3, 5, 7}
	for i := 0; i < b.N; i++ {
		Addi(c...)
	}
}
func BenchmarkAddf(b *testing.B) {
	c := []float64{2.1, 3.2, 5.3, 7.4}
	for i := 0; i < b.N; i++ {
		Addf(c...)
	}
}
