package main

import (
	"encoding/json"
	"fmt"
)

type exa struct {
	ID    int64
	Name  string
	Hobby []string
}

func main() {
	var m exa
	b := []byte(`{"ID" : 1, "Name" : "Denile Guo", "Hobby" : ["swimming", "playing basketball"]}`)
	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m.Name)
}
