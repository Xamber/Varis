package varis

import "testing"

func TestVectorSum(t *testing.T) {
	vector := Vector{1.1, 2.2, 3.3}
	if vector.sum() != 6.6 {
		t.Error("Sums not equal")
	}

	zero := Vector{}
	if zero.sum() != 0.0 {
		t.Error("Sums not equal")
	}
}

func TestVectorEqual(t *testing.T) {
	vector := Vector{1.1, 2.2, 3.3}
	equal := Vector{1.1, 2.2, 3.3}
	notEqual := Vector{5.1, 9.2, 100.3}
	notSameLenght := Vector{5.1}

	if !vector.is(equal) {
		t.Error("Error in vector equal function")
	}

	if vector.is(notEqual) {
		t.Error("Error in vector equal function")
	}

	if vector.is(notSameLenght) {
		t.Error("Error in vector equal function")
	}
}

func TestGenerateUUID(t *testing.T) {
	uuid := generateUUID()
	if len(uuid) != 36 {
		t.Error("Len of uuid not a 36:", len(uuid))
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

func TestPanicVector(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	vector := Vector{1.1, 2.2, 3.3}
	vector.broadcast(make([]chan float64, 2))
}

func TestPanicNN(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	n := CreatePerceptron(2, 3, 1)
	n.Calculate(Vector{1.1, 2.2, 3.3})
}

func TestPanicBooleanFiled(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	field := BooleanField{}
	field.toSignal("asdasdasd")
}
