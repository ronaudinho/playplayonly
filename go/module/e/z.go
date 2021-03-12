package z

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
		Dir:     "../e",
		Module:  "e",
		Package: "z",
	})
	fmt.Println(string(b))
}
