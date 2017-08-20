package burzedzisnet

import "testing"
import "fmt"

func TestMyComplexTypeMiejscowosc_IsSpec(t *testing.T) {
	crds := []MyComplexTypeMiejscowosc{
		{X: 51.02, Y: 17.02},
		{X: 0.0, Y: 12.34},
	}
	for _, v := range crds {
		if v.IsSpec() == false {
			t.Fatalf("MyComplexTypeMiejscowosc{%f, %f}.isSpec() returned true - expected false", v.X, v.Y)
		}
	}
}

func TestMyComplexTypeMiejscowosc_IsSpec2(t *testing.T) {
	crds := []MyComplexTypeMiejscowosc{
		MyComplexTypeMiejscowosc{X: 0, Y: 0},
	}
	for _, v := range crds {
		if v.IsSpec() == true {
			t.Fatalf("MyComplexTypeMiejscowosc{%f, %f}.isSpec() returned false - expected true", v.X, v.Y)
		}
	}
}

func TestMyComplexTypeOstrzezenia_IsSafe(t *testing.T) {
	warns := []MyComplexTypeOstrzezenia{
		{Mroz: 1},
		{Upal: 1},
		{Wiatr: 1},
		{Opad: 1},
		{Traba: 1},
	}
	for _, v := range warns {
		if v.IsSafe() == true {
			t.Fatalf("%v.isSafe() returned true - expected false", v)
		}
	}
}

func TestMyComplexTypeOstrzezenia_IsSafe2(t *testing.T) {
	warn := MyComplexTypeOstrzezenia{}
	if warn.IsSafe() == false {
		t.Fatalf("%v.isSafe() returned false - expected true", warn)
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
