package main

import "testing"

func TestSum(t *testing.T) {
	var calculator Arithmetic = Calculator{}
	sum := calculator.Add(5,5)
	if sum != 10 {
		t.Errorf("Add() incorrect, got: %d, wanted: %d", sum, 10)
	}
}

func TestSub(t *testing.T) {
	var calculator Arithmetic = Calculator{}
	diff := calculator.Sub(5,5)
	if diff != 0 {
		t.Errorf("Sub() incorrect, got: %d, wanted: %d", diff, 0)
	}
}

func TestMult(t *testing.T) {
	var calculator Arithmetic = Calculator{}
	prod := calculator.Mult(5,5)
	if prod != 25 {
		t.Errorf("Mult() incorrect, got: %d, wanted: %d", prod, 25)
	}
}

func TestDiv(t *testing.T) {
	var calculator Arithmetic = Calculator{}
	quo := calculator.Div(5,5)
	if quo != 1 {
		t.Errorf("Div() incorrect, got: %d, wanted: %d", quo, 1)
	}
}