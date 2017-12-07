package main

import (
	"fmt"
	"github.com/xamber/Varis"
	"reflect"
)

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

		boolF, ok := field.Interface().(BooleanField)
		if ok {
			fmt.Println(boolF)
		}

		fmt.Printf("%s:\tValue: %v\t Direction: %s\n", typeField.Name, field.Interface(), direction)
	}

	fmt.Println(inputs, output)

	var run ModelFunction = func(input []interface{}) []interface{} {

		return []interface{}{}
	}

	return run

}

func main() {

	n := varis.CreatePerceptron(2, 3, 1)

	f := &Model{
		Network: &n,
		X1:      BooleanField{},
		X2:      BooleanField{},
		O:       BooleanField{},
	}

	GenerateRunFunction(f)
	//run := GenerateRunFunction(f)
}
