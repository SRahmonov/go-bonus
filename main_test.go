package main

import "testing"

func Test_amountOfBonus(t *testing.T) {

	tests := []struct {
		name  string
		sales [] int
		want  int
	}{
		// TODO: Add test cases.
		{"Have bonus", []int{12_000, 8_000, 15_000, 9_000}, 350},
		{"No bonus", []int{10_000}, 0},
	}
	for _, tests := range tests {
		got := amountOfBonus(tests.sales)
		if got != tests.want {
			t.Error("Test fail", tests.name, "got", got, "want", tests.want)
		}
	}
}
