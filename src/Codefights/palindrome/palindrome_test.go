package palindrome

import "testing"

func Test_PALINDROME_ODD_OK(t *testing.T) {
	expected := false
	output := CheckPalindrome("aba")
	if expected != output {
		t.Errorf("Failed Test_PALINDROME_ODD_OK. \n Expected: %t \n Output: %t ", expected, output)
	}
}

func Test_PALINDROME_EVEN_OK(t *testing.T) {
	expected := false
	output := CheckPalindrome("aa")
	if expected != output {
		t.Errorf("Failed Test_PALINDROME_EVEN_OK. \n Expected: %t \n Output: %t ", expected, output)
	}
}

func Test_NO_PALINDROME_OK(t *testing.T) {
	expected := false
	output := CheckPalindrome("abc")
	if expected != output {
		t.Errorf("Failed Test_NO_PALINDROME_OK. \n Expected: %b \n Output: %b ", expected, output)
	}
}
