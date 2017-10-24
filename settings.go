package varis

type aFunction func(x float64) float64

var DEBUG bool = false

var ACTIVATION_FUNCTION aFunction = activation_sigmoid
var DEACTIVATION_FUNCTION aFunction = derivative_sigmoid
