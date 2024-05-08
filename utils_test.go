package edsger

import "testing"

func TestSigned(t *testing.T) {
	if Signed[float32]() != true {
		t.Fatal("Invalid result for float32")
	}
	if Signed[int]() != true {
		t.Fatal("Invalid result for int")
	}
	if Signed[uint]() != false {
		t.Fatal("Invalid result for uint")
	}
}
