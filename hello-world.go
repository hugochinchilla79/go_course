package main

import "fmt"
import "os"
/*
type Car struct {
	brand string
	model string
	color string
	price float32
}

func (car Car) String() string {
	return fmt.Sprintf("Brand: %s", car.brand)
}

func main(){
	var sum int = 5 + 9
	var name string = "Hello World"
	fmt.Println(name)
	fmt.Println("Sum is", sum)
	// time.Sleep(5 * time.Second)
	var car1 Car
	car1.brand = "Toyota"
	car1.model = "Corolla"
	car1.color = "Black"
	car1.price = 115000
	fmt.Println(car1)
}*/


func car(attributes ...string) string {
	return fmt.Sprintf("Brand: %s, Model: %s, Color: %s, Price: %s", attributes[0], attributes[1], attributes[2], attributes[3])
}

func main() {
	fmt.Println(car("Toyota", "Corolla", "Black", "115000"))	
	
	edad := 18
	
	if edad > 18 {
		fmt.Println("Eres mayor de edad")
	} else if (edad < 18) {
		fmt.Println("No tienes ni edad?")
	}else {
		fmt.Println("Eres menor de edad")
	}	

}

