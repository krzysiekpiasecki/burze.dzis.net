package burzedzisnet

import "testing"
import "fmt"

func TestMyComplexTypeMiejscowoscIsSpec(t *testing.T) {
	crds := []MyComplexTypeMiejscowosc{
		MyComplexTypeMiejscowosc{X: 51.02, Y: 17.02},
		MyComplexTypeMiejscowosc{X: 0.0, Y: 12.34},
	}
	for _, v := range crds {
		if v.IsSpec() == false {
			t.Fatalf("MyComplexTypeMiejscowosc [%f, %f] is specified", v.X, v.Y)
		}
	}
}

func TestMyComplexTypeMiejscowoscIsNotSpec(t *testing.T) {
	crds := []MyComplexTypeMiejscowosc{
		MyComplexTypeMiejscowosc{X: 0, Y: 0},
	}
	for _, v := range crds {
		if v.IsSpec() == true {
			t.Fatalf("MyComplexTypeMiejscowosc [%f, %f] is not specified", v.X, v.Y)
		}
	}
}

func ExampleMyComplexTypeMiejscowosc_IsSpec() {
	l := MyComplexTypeMiejscowosc{0, 0}
	fmt.Print(l.IsSpec())
	// Output: false
}

func ExampleMyComplexTypeMiejscowosc_IsSpec_second() {
	l := MyComplexTypeMiejscowosc{52.07, 12.25}
	fmt.Print(l.IsSpec())
	// Output: true
}
