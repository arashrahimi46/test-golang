package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	person1 := func(firstName string ,lastName string) Person{
		fmt.Println("we are calling anynymous function")
		return Person{
			FirstName: firstName,
			LastName:  lastName,
			Age:       20,
		}
	}
	person2:= person1("Arash" , "rahimi")
	fmt.Println(person2.FirstName)
	fmt.Println("we are developing new features")

}



func (p *Person)GetName()string {
	return p.FirstName + " " + p.LastName
}


func GetNewPerson(createNewPerson func(firstName string, lastName string) Person) Person {
	return createNewPerson("arash" , "rahimi")
}