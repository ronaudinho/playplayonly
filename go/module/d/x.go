package x

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
		Dir:     "../d",
		Module:  "x",
		Package: "x",
	})
	fmt.Println(string(b))
}
