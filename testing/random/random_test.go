package random

import (
	"reflect"
	"testing"
)

// TestPrintHelloWorld is a simple test for testing the return of a function.
func TestReturnHelloWorld(t *testing.T) {
	result := ReturnHelloWorld()
	if result != "hello world" {
		t.Fatal("hello world was not returned")
	}
}

func TestReturnPassedString(t *testing.T) {
	result := ReturnPassedString("peter")
	if result != "peter" {
		t.Fatal("peter was not returned")
	}
	if reflect.TypeOf(result) != reflect.TypeOf("peter") {
		t.Fatal("type returned was not of type string")
	}
}

func TestReturnSomething(t *testing.T) {
	result := returnSomething()
	if result != "blah" {
		t.Fatal("blah not returned")
	}
}
