package main

import "testing"

func TestIsPalidrone(t *testing.T) {
	_ = isPalidrone("A man, a plan, a canal: Panama")

	_ = isPalidrone("race a car")

	_ = isPalidrone(" ")

	_ = isPalidrone(".,")
}
