package b

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
		Dir:     "../b",
		Module:  "b",
		Package: "b",
	})
	fmt.Println(string(b))
}
