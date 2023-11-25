package main

import "fmt"

type fruit []string


func (f fruit) print(){

	count := len(f)

	fmt.Println("---in terms of count---")
	for i := 0; i < count; i++ {
		fmt.Println(f[i])
		
	}

	fmt.Println("-- in terms of range")
	for i, fru := range f{
		fmt.Println(i,fru)
	}

}
