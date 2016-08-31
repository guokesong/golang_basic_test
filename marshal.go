package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type exa struct {
	name  string
	id    int64
	hobby []string
}

func main() {
	type exa1 struct {
		Name  string
		ID    int
		Hobby []string
	}
	myjson := exa1{
		Name:  "Deniel Guo",
		ID:    1,
		Hobby: []string{"swimming", "basketball"},
	}
	b, err := json.Marshal(myjson)
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout.Write(b)
}
