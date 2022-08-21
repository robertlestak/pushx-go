package main

import (
	"bytes"
	"encoding/json"

	"github.com/robertlestak/pushx-go/pkg/pushx"
)

type ExampleData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Work string `json:"work"`
}

func main() {
	data := ExampleData{
		Name: "John Doe",
		Age:  42,
		Work: "Software Engineer",
	}
	jd, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	args := []string{
		"-driver",
		"redis-list",
		"-redis-host",
		"localhost",
		"-redis-port",
		"6379",
		"-redis-key",
		"test-redis-list",
	}
	if err := pushx.Pushx(bytes.NewReader(jd), args); err != nil {
		panic(err)
	}
}
