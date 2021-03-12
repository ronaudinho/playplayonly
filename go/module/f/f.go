package f

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
		Dir:     "../f",
		Module:  "f",
		Package: "f",
	})
	fmt.Println(string(b))
}
