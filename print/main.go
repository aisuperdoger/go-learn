package main

import "fmt"


type Person struct {
	Name string
	Age int
}

func (p Person) String() string{
	return fmt.Sprintf("%s %d",p.Name,p.Age)
}

func main(){

	person := []Person{
		{"1",1},
		{"2",2},
		{"3",3},		
	}

	fmt.Printf("%v \n",person)
}