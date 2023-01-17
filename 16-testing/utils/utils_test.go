package utils

import "testing"

func TestIsPrime_11(t *testing.T) {
	//Arrange
	no := 11
	expectedResult := false

	//Act
	actualResult := IsPrime(no)

	//Assert
	if actualResult != expectedResult {
		t.Errorf("Expected %t but got %t\n", expectedResult, actualResult)
	}
}

func TestIsPrime(t *testing.T) {
	testData := []struct {
		input          int
		name           string
		expectedResult bool
		actualResult   bool
	}{
		{input: 7, name: "IsPrime_7", expectedResult: true},
		{input: 9, name: "IsPrime_9", expectedResult: false},
		{input: 11, name: "IsPrime_11", expectedResult: true},
		{input: 13, name: "IsPrime_13", expectedResult: true},
		{input: 10, name: "IsPrime_10", expectedResult: false},
	}
	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			td.actualResult = IsPrime(td.input)
			if td.actualResult != td.expectedResult {
				t.Errorf("Expected %t but got %t\n", td.expectedResult, td.actualResult)
			}
		})
	}
}
