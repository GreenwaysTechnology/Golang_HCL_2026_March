package main

import "oopapps/composition"

func getPointer() *int {
	x := 10
	return &x
}
func CreatePersonReturnPointer(FirstName string, LastName string, Address composition.Address) *composition.Person {
	//p := Person{
	//	FirstName: FirstName,
	//	LastName:  LastName,
	//	Address:   Address,
	//}
	//return &p
	return &composition.Person{
		FirstName: FirstName,
		LastName:  LastName,
		Address:   Address,
	}
}

func main() {
	//composition.PrintDetails()
	//composition.PointerToStruct()
	//p := composition.CreatePerson("Subramanian", "Murugan", composition.Address{
	//	State: "TN",
	//	City:  "Coimbatore",
	//})
	//println(p.FirstName, p.LastName, p.Address.State, p.Address.City)
	composition.CreatePersonFactory()
}
