# roman-numerals

## How to Install

````
go get github.com/fraidev/roman-numerals
````

## How to Use

````
package main

import (
	"fmt"
	"github.com/fraidev/roman-numerals"
)

func main() {
	num := roman.ToNumber("XLII")
	fmt.Printf("%d \n", num)
}

