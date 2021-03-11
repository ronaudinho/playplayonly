// variable scope playplay
package main

import (
	"fmt"
)

func main() {
	nen()
	enn()
	nenen()
}

// never leh cher
func never() error {
	return nil
}

// always liddat ah how
func always() error {
	return fmt.Errorf("error leh")
}

// nil err nil
func nen() {
	var err error
	err = never()
	fmt.Println(err)
	// scope, not related to err above
	if err := always(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(err)
}

// err nil nil
func enn() {
	var err error
	// scope, not related to err above
	if err := always(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(err)
	err = always()
	// scope, not related to err above
	if err := never(); err == nil {
		fmt.Println(err)
	}
}

// nil err nil err nil err
func nenen() {
	if err := never(); err == nil {
		fmt.Printf("n")
		if err := always(); err != nil {
			fmt.Printf("e")
			if err := never(); err == nil {
				fmt.Printf("n")
			}
			if err != nil {
				fmt.Printf("e")
			}
		}
		if err == nil {
			fmt.Println("n")
		}
	}
}
