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
