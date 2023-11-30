package main

import "fmt"



//2 different methods which depicts the difference between the function recevier and the typical function

type fruit []string


//This is function receiver which will be a method of type fruit

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


//Whereas this is a normal method and is not a method type fruit.

func getOnPlate(f fruit, size int) fruit{

	return f[:size]
}