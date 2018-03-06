package main

import "testing"

func Test_01(t *testing.T) {

	expected := "1:aaa/1:nnn/1:gg/2:ee/2:ff/2:ii/2:oo/2:rr/2:ss/2:tt"

	output := Mix(" In many languages", " there's a pair of functions")

	if expected != output {
		t.Errorf("Failed Test_01. \n Expected: %s \n Output: %s ", expected, output)
	}
}

func Test_02(t *testing.T) {

	expected := "2:eeeee/2:yy/=:hh/=:rr"

	output := Mix("Are they here", "yes, they are here")

	if expected != output {
		t.Errorf("Failed Test_02. \n Expected: %s \n Output: %s ", expected, output)
	}
}
