package varis

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

// generateUUID  generate simple uuid.
func generateUUID() string {
	b := make([]byte, 16)
	rand.Read(b)
	uuid := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

// ToJSON dump and transform Perceptron to json string.
func ToJSON(network Perceptron) string {
	networkDump := make(map[string][]interface{})

	networkDump["Layers"] = []interface{}{}
	networkDump["Synapses"] = []interface{}{}

	cache := make(map[Neuron]string)

	for _, l := range network.layers {
		layerDump := []interface{}{}
		for _, n := range l {
			uuid := generateUUID()
			cache[n] = uuid
			layerDump = append(layerDump, map[string]float64{uuid: n.getCore().weight})
		}
		networkDump["Layers"] = append(networkDump["Layers"], layerDump)
	}
	for _, l := range network.layers {
		for _, n := range l {
			for _, os := range n.getCore().conn.outSynapses {
				synapseDump := map[string]interface{}{
					"in":     cache[os.inNeuron],
					"out":    cache[os.outNeuron],
					"weight": os.weight,
				}
				networkDump["Synapses"] = append(networkDump["Synapses"], synapseDump)
			}
		}
	}

	jsonNetwork, _ := json.Marshal(networkDump)
	return string(jsonNetwork)
}

// FromJSON load json string and create Perceptron.
func FromJSON(jsonString string) Perceptron {
	networkLoad := make(map[string][]interface{})

	json.Unmarshal([]byte(jsonString), &networkLoad)

	network := Perceptron{}
	network.input = make([]chan float64, 0)
	network.output = make([]chan float64, 0)

	cache := make(map[string]Neuron)

	for index, loadLayer := range networkLoad["Layers"] {
		layer := []Neuron{}
		normalizeLayer := loadLayer.([]interface{})
		for _, loadNeuron := range normalizeLayer {
			var neuron Neuron
			normalizedNeuron := loadNeuron.(map[string]interface{})
			for uuid, value := range normalizedNeuron {
				weight := value.(float64)
				switch index {
				case 0:
					channel := make(chan float64)
					neuron = INeuron(weight, channel)
					network.input = append(network.input, channel)
				case len(networkLoad["Layers"]) - 1:
					channel := make(chan float64)
					neuron = ONeuron(weight, channel)
					network.output = append(network.output, channel)
				default:
					neuron = HNeuron(weight)
				}
				layer = append(layer, neuron)
				cache[uuid] = neuron
			}
		}
		network.layers = append(network.layers, layer)
	}

	for _, s := range networkLoad["Synapses"] {
		normalizedSynapse := s.(map[string]interface{})
		inNeuronUUID := normalizedSynapse["in"].(string)
		outNeuronUUID := normalizedSynapse["out"].(string)
		weight := normalizedSynapse["weight"].(float64)
		ConnectNeurons(cache[inNeuronUUID], cache[outNeuronUUID], weight)
	}

	network.RunNeurons()

	return network
}
