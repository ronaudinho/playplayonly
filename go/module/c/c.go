package c

import (
	"encoding/json"
	"fmt"
)

type info struct {
	Dir     string
	Module  string
	Package string
}

func Info() {
	b, _ := json.Marshal(info{
		Dir:     "../c",
		Module:  "c",
		Package: "c",
	})
	fmt.Println(string(b))
}
