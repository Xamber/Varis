package varis

import (
	"encoding/json"
)

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

type layerDump []neuronDump

type networkDump struct {
	Neurons  []layerDump
	Synapses []synapseDump
}

func (network *Network) Dump() networkDump {
	dump := networkDump{}

	for _, l := range network.Layers {
		layerDump := layerDump{}
		for _, n := range l {
			neuronDump := neuronDump{
				n.getUUID(),
				n.getWeight(),
			}
			layerDump = append(layerDump, neuronDump)
			for _, os := range n.getConnection().outSynapses {
				synapseDump := synapseDump{
					UUID:      os.uuid,
					Weight:    os.weight,
					InNeuron:  os.inNeuron.getUUID(),
					OutNeuron: os.outNeuron.getUUID(),
				}
				dump.Synapses = append(dump.Synapses, synapseDump)
			}
		}
		dump.Neurons = append(dump.Neurons, layerDump)
	}
	return dump
}

func (load networkDump) Load() Network {
	cache := make(map[string]Neuroner)

	network := Network{Output: make([]chan float64, 0)}
	for index, loadLayer := range load.Neurons {
		layer := Layer{}
		for _, n := range loadLayer {
			var neuron Neuroner
			switch index {
			case 0:
				neuron = network.createInputNeuron(n.UUID, n.Weight)
			case len(load.Neurons) - 1:
				neuron = network.createOutputNeuron(n.UUID, n.Weight)
			default:
				neuron = network.createHiddenNeuron(n.UUID, n.Weight)
			}
			layer.AddNeuron(neuron)
			cache[neuron.getUUID()] = neuron
		}
		network.AddLayer(layer)
	}

	for _, s := range load.Synapses {
		ConnectNeurons(cache[s.InNeuron], cache[s.OutNeuron], s.UUID, s.Weight)
	}

	network.RunNeurons()

	return network
}

func ToJSON(network Network) string {
	dump := network.Dump()
	jsonString, _ := json.Marshal(dump)
	return string(jsonString)
}

func FromJSON(jsonString string) Network {
	var load networkDump

	json.Unmarshal([]byte(jsonString), &load)
	network := load.Load()

	return network
}
