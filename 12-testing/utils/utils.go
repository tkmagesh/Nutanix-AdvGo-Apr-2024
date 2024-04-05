package utils

var Count int = 0

func init() {
	Count = 0
}

func IsPrime_1(no int) bool {
	Count++
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func IsPrime_2(no int) bool {
	Count++
	for i := 2; i <= (no - 1); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
