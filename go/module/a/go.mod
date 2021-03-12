module a

go 1.15

require (
	b v0.0.0 // v1.2.3 format
	d v0.0.1 // version number does not matter
	e v0.1.0 // as long as they follow format
	github.com/ronaudinho/erni v0.0.0-20200805101513-1e5abc7868e8
	x v1.0.0 // but major can only be 0 or 1
	x/v2 v2.0.0 // use /v{{number}} require format to go past v1
)

replace (
	b => ../b
	d => ../c
	e => ../e
	github.com/ronaudinho/erni => ../f
	x => ../d
	x/v2 => ../g
)
