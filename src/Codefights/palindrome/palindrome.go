package palindrome

func CheckPalindrome(inputString string) bool {
	var length = len(inputString)
	if length == 1 {
		return true
	} else {
		current := string(inputString[0]) == string(inputString[length-1])
		next := inputString[1 : length-1]
		return current && CheckPalindrome(next)
	}
}
