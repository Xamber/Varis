package varis

import (
	"reflect"
)

// ModelFunction type of calculate model function
type ModelFunction func(values) values

type values []interface{}

// Field store two methods: toSignal and fromSignal for convert values to signals
type Field interface {
	toSignal(input interface{}) float64
	fromSignal(signal float64) interface{}
}

// BooleanField implement bool type of Model field.
type BooleanField struct{}

func (f BooleanField) toSignal(input interface{}) float64 {
	value, ok := input.(bool)
	if !ok {
		panic("Value of BooleanField is not bool")
	}

	if value == true {
		return 1.00
	}
	return 0.00
}

func (f BooleanField) fromSignal(signal float64) interface{} {
	if signal >= 0.5 {
		return true
	}
	return false
}

// GenerateRunFunction generate RunFunction from model.
func GenerateRunFunction(f interface{}) ModelFunction {
	model := reflect.ValueOf(f)

	var network *Perceptron

	inputs := []Field{}
	output := []Field{}

	for i := 0; i < model.NumField(); i++ {
		field := model.Field(i)
		typeField := model.Type().Field(i)

		direction := typeField.Tag.Get("direction")

		modelField, isField := field.Interface().(Field)
		if isField && direction == "input" {
			inputs = append(inputs, modelField)
		}
		if isField && direction == "output" {
			output = append(output, modelField)
		}

		nn, ok := field.Interface().(*Perceptron)
		if ok {
			network = nn
		}
	}

	var run ModelFunction = func(input values) values {

		in := Vector{}
		out := values{}

		for i, v := range input {
			in = append(in, inputs[i].toSignal(v))
		}

		output := network.Calculate(in)

		for i, v := range output {
			out = append(out, inputs[i].fromSignal(v))
		}

		return out
	}

	return run
}
