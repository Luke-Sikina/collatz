package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestBaseCollatz(t *testing.T) {
	actual := Collatz(1)
	expected := []uint{1}
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Actual: " + toString(actual) + " Expected: " + toString(expected))
	}
}

func TestEvenCollatz(t *testing.T) {
	actual := Collatz(16)
	expected := []uint{16, 8, 4, 2, 1}
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Actual: " + toString(actual) + " Expected: " + toString(expected))
	}
}

func TestOddCollatz(t *testing.T) {
	actual := Collatz(3)
	expected := []uint{3, 10, 5, 16, 8, 4, 2, 1}
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Actual: " + toString(actual) + " Expected: " + toString(expected))
	}
}

func toString(slice []uint) string {
	return strings.Trim(fmt.Sprint(slice), "[]")
}
