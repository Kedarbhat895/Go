package main


func main(){

	fruits := fruit{"apple","banana",newFruit()}
	fruits = append(fruits, "berries")
	fruits.print()

}

func newFruit() string{
	return "orange"
}