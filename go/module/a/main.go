// import as require
// use as replace
package main

import (
	"b"
	"d"
	"e"
	"github.com/ronaudinho/erni"
	XD "x"
	"x/v2"
)

func main() {
	bInfo()
	dInfo()
	eInfo()
	erniInfo()
	XDxInfo()
	xv2Info()
}

// require non-existent remote package
// // imported as "b"
// replace with local package
// // shares the same module name
// probably the only real-life use case of this
func bInfo() {
	b.Info()
}

// require non-existent remote package
// // imported as "d"
// replace with non-related local package
// // using function existing in said required
func dInfo() {
	c.Info()
}

// require non-existent remote package
// // imported as "e"
// replace with local package
// // module named "e"
// // root package named "z"
func eInfo() {
	z.Info()
}

// require existing remote package
// // imported as "github.com/ronaudinho/erni"
// replace with non-related local package
// // using function not existing in said required
func erniInfo() {
	f.Info()
}

// XD
// require non-existent remote package
// // imported as "x"
// replace with local package
// // shares the same module name
func XDxInfo() {
	XD.Info()
}

// require non-existent remote package
// // imported as "x/v2"
// // major version upgrade
// replace with local package
func xv2Info() {
	g.Info()
}
