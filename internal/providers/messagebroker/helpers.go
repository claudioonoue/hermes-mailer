package messagebroker

import "encoding/json"

// toJSON is a method that will convert the message data structure to JSON bytes.
func toJSONBytes(message any) ([]byte, error) {
	return json.Marshal(message)
}

// fromJSON is a method that will convert the JSON Bytes to message data structure.
func fromJSONBytes(data []byte, body any) error {
	return json.Unmarshal(data, body)
}
