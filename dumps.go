package varis

import (
	"encoding/json"
)

func ToJSON(network Network) string {

	mapD := map[string]string{"apple": generate_uuid(), "lettuce": generate_uuid()}
	mapB, _ := json.Marshal(mapD)

	return string(mapB)
}
