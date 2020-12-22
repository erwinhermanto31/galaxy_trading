package main

import "testing"

func TestFind(t *testing.T) {
	m := NewMain()

	rmc := []string{"glob", "prok", "pish", "tegj", "Silver", "Gold", "Iron"}
	val := "glob"

	f := m.Find(rmc, val)

	if f != true {
		t.Errorf("I should return true, but %v", f)
	}
}

func TestConvertToNumber(t *testing.T) {
	m := NewMain()

	numberString := "XI"

	c := m.ConvertToNumber(numberString)

	if c != 11 {
		t.Errorf("I should return 10, but %v", c)
	}
}

func TestConvertToRomawi(t *testing.T) {
	m := NewMain()

	sentence := []string{"glob", "prok", "Gold"}

	ctr := m.ConvertToRomawi(sentence)

	if ctr != "57800" {
		t.Errorf("I should return 57800, but %v", ctr)
	}
}

func TestSentenceParsing(t *testing.T) {
	m := NewMain()

	sentence := "how many Credits is glob prok Silver ?"

	_, rmw := m.SentenceParsing(sentence)

	if rmw != "68" {
		t.Errorf("I should return 68, but %v", rmw)
	}
}
