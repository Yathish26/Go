package main

import "fmt"


func main(){
	wilson("Yathish",24)
	price("Laptop",50000)
}

func wilson(name string, age int32){
	fmt.Printf("Hi my name is %v and i am %v years old",name,age)
}

func price(product string, price int32){
	fmt.Printf("The price of %v is %v",product,price)
}