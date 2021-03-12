package g

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
		Dir:     "../g",
		Module:  "g",
		Package: "g",
	})
	fmt.Println(string(b))
}
