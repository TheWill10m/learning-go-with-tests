package main

import "testing"

func TestStum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertCorrectNumber(t, got, want, numbers)
	})
}

func assertCorrectNumber(t testing.TB, got, want int, given []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d, given %v", got, want, given)
	}
}
