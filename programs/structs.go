package main


import "fmt"

type fruit struct{
	price float64
	color string
	name string
}


type basket struct{
	numberOfFruits int
	fruit fruit
}


func main(){

	apple := fruit{
		price : 10.0,
		color : "red",
		name: "apple",
	}

	newBasket := basket{
		numberOfFruits : 3,
		fruit : apple,
	}

	newBasket.updatePrice(20.0) //pass by value (passing the same memory location)

	fmt.Printf("%+v", newBasket)


	updatePrice2(&newBasket, 30.0)


	fmt.Println(apple) //price of apple will remain same even tough it is changed in basket (some weird behaviour)
	fmt.Printf("%+v", newBasket)
	
}


func (pointedBasket *basket) updatePrice(newPrice float64){
	(*pointedBasket).fruit.price = newPrice
}

func updatePrice2(b *basket, p float64){
	b.fruit.price = p
}