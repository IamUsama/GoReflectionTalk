// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Example shows how to reflect over a struct type value.
package main

import (
	"fmt"
	"reflect"
)

// user represents a basic user in the system.
type user struct {
	name     string
	age      int
	building float32
	secure   bool
	roles    []string
}

func main() {

	// Create a struct type user value.
	u := user{
		name:     "Cindy",
		age:      27,
		building: 321.45,
		secure:   true,
		roles:    []string{"admin", "developer"},
	}

	// Display the information about the user value.
	v := reflect.ValueOf(u)
	fmt.Printf("Kind: %v\tType: %v\t\tNumFields: %v\n", v.Kind(), v.Type(), v.NumField())
}
