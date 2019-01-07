package strsim

import (
	"reflect"
	"testing"
)

func TestCompare(t *testing.T) {
	for _, test := range testCases1 {
		observed := Compare(test.a, test.b)
		if observed != test.expected {
			t.Fatalf("Compare(%s, %s) = %f, want %f",
				test.a, test.b, observed, test.expected)
		}
	}
}

func TestFindBestMatch(t *testing.T) {
	for _, test := range testCases2 {
		observed, err := FindBestMatch(test.s, test.targets)
		if err != nil && err.Error() != test.err.Error() {
			t.Fatalf("Unexpected error from FindBestMatch: %v", err)
		}

		if !reflect.DeepEqual(observed, test.expected) {
			t.Fatalf("FindBestMatch(%s, %v) = %+v, want %+v", test.s, test.targets, observed, test.expected)
		}
	}
}

func TestSortedByScore(t *testing.T) {
	for _, test := range testCases3 {
		observed := test.actual
		observed.SortedByScore()
		if !reflect.DeepEqual(observed, test.expected) {
			t.Fatalf("SortedByScore() = %+v, want %+v", observed, test.expected)
		}
	}
}

func BenchmarkCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases1 {
			Compare(test.a, test.b)
		}
	}
}

func BenchmarkFindBestMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases2 {
			FindBestMatch(test.s, test.targets)
		}
	}
}

func BenchmarkSortedByScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases3 {
			observed := test.actual
			observed.SortedByScore()
		}
	}
}
