package numbers

import "testing"

func TestCountPrimes(t *testing.T) {
	tests := []IntergerTest{
		{Name: "test data", Input: 8, Output: 4},
	}
	for _, test := range tests {
		result := CountPrime(test.Input)
		if result != test.Output {
			t.Fatalf("test %s failed", test.Name)
		}
	}
}

func TestCountPrimesFast(t *testing.T) {
	tests := []IntergerTest{
		{Name: "test data", Input: 8, Output: 4},
	}
	for _, test := range tests {
		result := CountPrimeFast(test.Input)
		if result != test.Output {
			t.Fatalf("test %s failed", test.Name)
		}
	}
}
