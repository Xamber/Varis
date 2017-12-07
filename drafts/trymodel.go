package main

import (
	"fmt"
	"github.com/xamber/Varis"
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
		return 0.00
	} else {
		return 1.00
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

type Model struct {
	Network *varis.Perceptron
	X1      BooleanField `direction:"input"`
	X2      BooleanField `direction:"input"`
	O       BooleanField `direction:"output"`
}

func GenerateRunFunction(f interface{}) ModelFunction {
	model := reflect.ValueOf(f).Elem()

	var network *varis.Perceptron

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

		nn, ok := field.Interface().(*varis.Perceptron)
		if ok {
			network = nn
		}

		fmt.Printf("%s:\tValue: %v\t Direction: %s\n", typeField.Name, field.Interface(), direction)
	}

	fmt.Println(inputs, output)

	var run ModelFunction = func(input []interface{}) []interface{} {

		in := varis.Vector{}
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

func main() {

	n := varis.CreatePerceptron(2, 3, 1)

	dataset := varis.Dataset{
		{varis.Vector{0.0, 0.0}, varis.Vector{1.0}},
		{varis.Vector{1.0, 0.0}, varis.Vector{0.0}},
		{varis.Vector{0.0, 1.0}, varis.Vector{0.0}},
		{varis.Vector{1.0, 1.0}, varis.Vector{1.0}},
	}

	varis.BackPropagation(&n, dataset, 10000)

	type Model struct {
		Network *varis.Perceptron
		X1      BooleanField `direction:"input"`
		X2      BooleanField `direction:"input"`
		O       BooleanField `direction:"output"`
	}

	f := Model{Network: &n}

	calculate := GenerateRunFunction(f)
	output := calculate([]interface{}{true, "sdfsdf"})
	fmt.Println(output[0])
	//run := GenerateRunFunction(f)
}
