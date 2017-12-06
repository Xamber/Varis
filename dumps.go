package varis

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

func generate_uuid() string {
	b := make([]byte, 16)
	rand.Read(b)
	uuid := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

type NetworkDump struct {
	Neurons  [][]neuronDump
	Synapses []synapseDump
}

type neuronDump struct {
	UUID   string
	Weight float64
}

type synapseDump struct {
	UUID      string
	Weight    float64
	InNeuron  string
	OutNeuron string
}

func (network *Perceptron) Dump() NetworkDump {
	dump := NetworkDump{}

	neuronsUUIDs := make(map[*Neuron]string)

	for _, l := range network.layers {
		layerDump := []neuronDump{}
		for _, n := range l {
			uuid := generate_uuid()
			neuronsUUIDs[n] = uuid

			neuronDump := neuronDump{uuid, n.weight}
			layerDump = append(layerDump, neuronDump)
		}
		dump.Neurons = append(dump.Neurons, layerDump)
	}
	for _, l := range network.layers {
		for _, n := range l {
			for _, os := range n.conn.outSynapses {
				synapseDump := synapseDump{
					Weight:    os.weight,
					InNeuron:  neuronsUUIDs[os.inNeuron],
					OutNeuron: neuronsUUIDs[os.outNeuron],
				}
				dump.Synapses = append(dump.Synapses, synapseDump)
			}
		}
	}

	return dump
}

func (load NetworkDump) Load() Perceptron {
	cache := make(map[string]*Neuron)

	network := Perceptron{output: make([]chan float64, 0)}
	for index, loadLayer := range load.Neurons {
		layer := []*Neuron{}
		for _, n := range loadLayer {
			var neuron *Neuron
			var channel chan float64

			switch index {
			case 0:
				neuron, channel = CreateNeuron(InputNeuron, n.Weight)
				network.input = append(network.input, channel)
			case len(load.Neurons) - 1:
				neuron, channel = CreateNeuron(OutputNeuron, n.Weight)
				network.output = append(network.output, channel)
			default:
				neuron, _ = CreateNeuron(HiddenNeuron, n.Weight)
			}
			layer = append(layer, neuron)
			cache[n.UUID] = neuron
		}
		network.layers = append(network.layers, layer)
	}

	for _, s := range load.Synapses {
		ConnectNeurons(cache[s.InNeuron], cache[s.OutNeuron], s.Weight)
	}

	network.RunNeurons()

	return network
}

func ToJSON(network Perceptron) string {
	dump := network.Dump()
	jsonString, _ := json.Marshal(dump)
	return string(jsonString)
}

func FromJSON(jsonString string) Perceptron {
	var load NetworkDump
	json.Unmarshal([]byte(jsonString), &load)
	network := load.Load()
	return network
}
