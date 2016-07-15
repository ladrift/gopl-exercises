package tempconv

import "testing"

func TestCToF(t *testing.T) {
	var c Celsius = 37
	f := CToF(c)
	if f != Fahrenheit(98.6) {
		t.Errorf(`CToF(%g) == %g failed`, c, f)
	}
}

func TestFToC(t *testing.T) {
	var f Fahrenheit = 95
	c := FToC(f)
	if c != Celsius(35) {
		t.Errorf(`FToC(%g) == %g failed`, f, c)
	}
}

func TestCToK(t *testing.T) {
	var c Celsius = 37
	k := CToK(c)
	if k != Kelvin(310.15) {
		t.Errorf(`CToK(%g) == %g failed`, c, k)
	}
}

func TestKToC(t *testing.T) {
	var k Kelvin = 1000
	c := KToC(k)
	if c != Celsius(726.85) {
		t.Errorf(`KToC(%g) == %g failed`, k, c)
	}
}
