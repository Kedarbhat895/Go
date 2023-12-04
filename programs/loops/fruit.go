package main

func main() {

	fruits := fruit{"apple", "banana", newFruit()}
	fruits = append(fruits, "greenApple")

	fruits.print()

	fruitOnPlate := getOnPlate(fruits, 3)
	fruitOnPlate.print()

	veggies := fruit{"1", "2"}

	veggies.print()

}

func newFruit() string {
	return "orange"
}
