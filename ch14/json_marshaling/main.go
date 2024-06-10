package main

import (
	"bytes"
	"encoding/json"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

type Data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func marshalData(data Data) ([]byte, error) {
	buffer := bufferPool.Get().(*bytes.Buffer)
	defer bufferPool.Put(buffer)
	buffer.Reset()

	err := json.NewEncoder(buffer).Encode(data)
	if err != nil {
		return nil, err
	}

	result := make([]byte, buffer.Len())
	copy(result, buffer.Bytes())
	return result, nil
}

func main() {
	data := Data{Name: "John Doe", Age: 30}
	jsonBytes, err := marshalData(data)
	if err != nil {
		println("Error marshaling JSON:", err.Error())
	} else {
		println("JSON output:", string(jsonBytes))
	}
}
