package composition

import "fmt"

type Address struct {
	City  string
	State string
}
type Person struct {
	FirstName string
	LastName  string
	//Structs Embbding
	Address Address
}

func PrintDetails() {
	p := Person{
		FirstName: "John",
		LastName:  "Smith",
		//embbeding
		Address: Address{
			City: "London",
		},
	}
	fmt.Println(p)
	fmt.Printf("%s %s %s %s\n", p.LastName, p.FirstName, p.Address.State, p.Address.City)
}
func PointerToStruct() {
	p := &Person{
		FirstName: "John",
		LastName:  "Smith",
		Address: Address{
			City: "London",
		},
	}
	fmt.Println(&p)
	fmt.Println(p.LastName, (*p).LastName)
}

// stack allocation
func CreatePerson(FirstName string, LastName string, Address Address) Person {
	p := Person{
		FirstName: FirstName,
		LastName:  LastName,
		Address:   Address,
	}
	return p
}

//factory function to initalize the structs with parameters

// constructors : newFunctionName(args) return Type, since go does not support
// constructors like other languages, constructors are specified via factory functions
//constructos factory functions should not called outside the current structs declaration

func newPerson(FirstName string, LastName string, Address Address) *Person {
	//p := Person{
	//	FirstName: FirstName,
	//	LastName:  LastName,
	//	Address:   Address,
	//}
	//return &p
	return &Person{
		FirstName: FirstName,
		LastName:  LastName,
		Address:   Address,
	}
}

func CreatePersonFactory() {
	p := newPerson("Subramanian", "Murugan", Address{
		State: "TN",
		City:  "Coimbatore",
	})
	fmt.Println(p)
}
