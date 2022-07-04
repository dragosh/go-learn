package main

import (
	"testing"
)

func TestGreet(t *testing.T) {
	actual := Greet("Dragosh")
	expected := "Hello Dragosh!\n"
	if actual != expected {
		t.Errorf("Test failed. Actual: '%q', Expected : '%q'", actual, expected)
	}
}

func TestDepart(t *testing.T) {
	actual := depart("Dragosh")
	expected := "Good bye Dragosh!\n"
	if actual != expected {
		t.Errorf("Test failed. Actual: '%q', Expected : '%q'", actual, expected)
	}
}

func TestGreetTableDriven(t *testing.T) {
	suite := []struct {
		input    string
		expected string
	}{
		{input: "Dragosh", expected: "Hello Dragosh!\n"},
		{input: "Suzana", expected: "Hello Suzana!\n"},
	}
	for _, s := range suite {
		actual := Greet(s.input)
		if actual != s.expected {
			t.Errorf("Test failed. Actual: '%q', Expected : '%q'", actual, s.expected)
		}
	}

}

func TestSum(t *testing.T) {
	actual := 2 + 2
	expected := 4
	if actual != expected {
		t.Errorf("Test Failed: Actual: '%v', Expected : '%v'", actual, expected)
	}
}

func TestSub(t *testing.T) {
	actual := 4 - 2
	expected := 2
	if actual != expected {
		t.Errorf("Test Failed: Actual: '%v', Expected : '%v'", actual, expected)
	}
}
