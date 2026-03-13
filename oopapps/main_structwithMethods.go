package main

import "fmt"

type Product struct {
	ID       string
	Name     string
	Price    float32
	Quantity int
}

// methods of Struct
// p Product is called receiver - value reciver
func (p Product) calculate() float32 {
	//type conversion (int to float)
	return p.Price * float32(p.Quantity)
}

func main() {
	//create instance with defaults
	var pro Product
	//later we can initalize
	pro.Name = "Subramanian Murugan"

	//create instance and initalize the values
	p := Product{
		ID:       "123",
		Name:     "Iphone",
		Price:    42.0,
		Quantity: 10,
	}
	fmt.Println(p.Name, p.Price, p.Quantity, p.ID)
	fmt.Println(p.calculate())
}
