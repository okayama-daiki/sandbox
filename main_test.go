package main

import "testing"

func TestIsPrime(t *testing.T) {
	t.Run("is prime", func(t *testing.T) {
		got := isPrime(7)
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("is not prime", func(t *testing.T) {
		got := isPrime(10)
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}

func TestSieve(t *testing.T) {
	t.Run("sieve of 10", func(t *testing.T) {
		got := sieve(10)
		want := []int{2, 3, 5, 7}

		if len(got) != len(want) {
			t.Errorf("got %v want %v", got, want)
		}

		for i := range got {
			if got[i] != want[i] {
				t.Errorf("got %v want %v", got, want)
			}
		}
	})
}
