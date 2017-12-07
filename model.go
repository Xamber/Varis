package varis

import (
	"reflect"
)

type Values []interface{}
type ModelFunction func([]interface{}) []interface{}
type Field interface {
	toSignal(input interface{}) float64
	fromSignal(signal float64) interface{}
}

type BooleanField struct{}

func (f BooleanField) toSignal(input interface{}) float64 {

	value, ok := input.(bool)
	if !ok {
		panic("Value of BooleanField is not bool")
	}

	if value == true {
		return 1.00
	} else {
		return 0.00
	}
}

func (f BooleanField) fromSignal(signal float64) interface{} {
	if signal >= 0.5 {
		return true
	} else {
		return false
	}
}

type IntegerField struct {
	min int64
	max int64
}

type ArrayField struct{}

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

	var run ModelFunction = func(input []interface{}) []interface{} {

		in := Vector{}
		out := []interface{}{}

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
