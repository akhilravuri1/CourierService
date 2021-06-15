package main

import "testing"

func TestDiscountWithValidData(t *testing.T) {
	if DiscountWithValidData() != true {
		t.Error("Data is not equal")
	}
}
