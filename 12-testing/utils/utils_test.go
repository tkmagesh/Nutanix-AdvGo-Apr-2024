package utils

import (
	"testing"
)

/* func Test_IsPrime_11(t *testing.T) {
	// arrange
	var no = 11
	var expected_result = true

	// act
	actual_result := IsPrime(no)

	// assert
	if actual_result != expected_result {
		t.Error("11 is a prime number")
	}
	fmt.Println("count :", Count)
}

func Test_IsPrime_12(t *testing.T) {
	// arrange
	var no = 12
	var expected_result = false

	// act
	actual_result := IsPrime(no)

	// assert
	if actual_result != expected_result {
		t.Error("12 is a prime number")
	}
	fmt.Println("count :", Count)
} */

/* func Test_IsPrime(t *testing.T) {
	var primeTestData []struct {
		no       int
		expected bool
	} = []struct {
		no       int
		expected bool
	}{
		{no: 11, expected: true},
		{no: 12, expected: false},
		{no: 13, expected: true},
		{no: 17, expected: true},
		{no: 19, expected: true},
	}
	for _, testData := range primeTestData {
		t.Run(fmt.Sprintf("Testing_IsPrime_%d", testData.no), func(t *testing.T) {
			// arrange
			var no = testData.no
			var expected_result = testData.expected

			// act
			actual_result := IsPrime(no)

			// assert
			if actual_result != expected_result {
				t.Errorf("testing isPrime(%d), expected = %v, actual = %v\n", no, expected_result, actual_result)
			}
		})
	}
}
*/

func BenchmarkIsPrime_1_997(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime_1(997)
	}
}

func BenchmarkIsPrime_2_997(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime_2(997)
	}
}
