package main

import (
	"fmt"
	"json-magic/people"
)

func main()  {
	husbandjson :=people.CreateHusband(people.Husband{Name: "Alfonso", Age: 21}	)
	fmt.Println(string(husbandjson))
}