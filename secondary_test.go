package varis

import "testing"

func TestVectorSum(t *testing.T) {
	vector := Vector{1.1, 2.2, 3.3}
	summa := vector.sum()
	if summa != 6.6 {
		t.Error("Sums not equal")
	}
}

func TestVectorZeroSum(t *testing.T) {
	vector := Vector{}
	summa := vector.sum()
	if summa != 0.0 {
		t.Error("Sums not equal")
	}
}

func TestGenerateUUID(t *testing.T) {
	uuid := generateUUID()
	if len(uuid) != 36 {
		t.Error("Len of uuid not a same:", len(uuid))
	}
}

func TestDeactivationFunction(t *testing.T) {
	if DEACTIVATION(0.123123123) != 0.24905493220237876 {
		t.Error("Wrong deactivation function")
	}
}

func TestActivationFunction(t *testing.T) {
	if ACTIVATION(0.123123123) != 0.5307419550064931 {
		t.Error("Wrong activation function")
	}
}
