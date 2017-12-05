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

	neuronsUUIDs := make(map[*Neuron]string)

	for _, l := range network.Layers {
		layerDump := layerDump{}
		for _, n := range l {
			uuid := generate_uuid()
			neuronsUUIDs[n] = uuid

			neuronDump := neuronDump{uuid, n.weight}
			layerDump = append(layerDump, neuronDump)
		}
		dump.Neurons = append(dump.Neurons, layerDump)
	}
	for _, l := range network.Layers {
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

func (load networkDump) Load() Network {
	cache := make(map[string]*Neuron)

	network := Network{Output: make([]chan float64, 0)}
	for index, loadLayer := range load.Neurons {
		layer := []*Neuron{}
		for _, n := range loadLayer {
			neuron := &Neuron{weight: n.Weight}
			neuron.callbackFunc = neuron.conn.broadcastSignals
			switch index {
			case 0:
				neuron.callbackFunc = nil
			case len(load.Neurons) - 1:
				outputChan := make(chan float64)
				neuron.callbackFunc = func(value float64) {
					outputChan <- value
				}
				network.Output = append(network.Output, outputChan)
			}
			layer = append(layer, neuron)
			cache[n.UUID] = neuron
		}
		network.AddLayer(layer)
	}

	for _, s := range load.Synapses {
		ConnectNeurons(cache[s.InNeuron], cache[s.OutNeuron], s.Weight)
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
