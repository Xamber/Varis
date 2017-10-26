package varis

import (
	"encoding/json"
)

type neuronDump struct {
	UUID   string  `json:"uuid"`
	Weight float64 `json:"weight"`
}

type synapseDump struct {
	UUID    string  `json:"uuid"`
	Weight  float64 `json:"weight"`
	InUUID  string  `json:"in"`
	OutUUID string  `json:"out"`
}

type networkDump struct {
	Neurons  [][]neuronDump `json:"neurons"`
	Synapses []synapseDump  `json:"synapses"`
}

func ToJSON(network Network) string {
	dump := networkDump{}
	for _, l := range network.Layers {
		var layerDump []neuronDump
		for _, n := range l.getNeurons() {
			layerDump = append(layerDump, neuronDump{n.getUUID(), n.getWeight()})
			for _, os := range n.getConnection().outSynapses {
				dump.Synapses = append(dump.Synapses, synapseDump{UUID: os.uuid, Weight: os.weight, InUUID: os.inputUUID, OutUUID: os.outputUUID})
			}
		}
		dump.Neurons = append(dump.Neurons, layerDump)
	}

	jsonString, _ := json.Marshal(dump)
	return string(jsonString)
}

func FromJSON(jsonString string) *Network {
	var load networkDump
	cache := make(map[string]Neuroner)

	json.Unmarshal([]byte(jsonString), &load)

	network := Network{Output: make([]chan float64, 0)}
	for index, loadLayer := range load.Neurons {
		layer := &layer{}
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
		createSynapse(cache[s.InUUID], cache[s.OutUUID], s.UUID, s.Weight)
	}

	network.RunNeurons()

	return &network
}
