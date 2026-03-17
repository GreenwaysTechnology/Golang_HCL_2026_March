package main

import "fmt"

type User struct {
	ID   int
	Name string
	Age  int
}

// users slice
var users []User

// create new user and append as part of slice
func createUser(name string, age int) {
	newID := len(users) + 1
	users = append(users, User{ID: newID, Name: name, Age: age})
}

// read All data from slice
func ReadAll() []User {
	return users
}

// read by id
func ReadByID(id int) (*User, bool) {
	for _, user := range users {
		if user.ID == id {
			return &user, true
		}
	}
	return nil, false
}

// remove user from slice by index
func DeleteByID(id int) bool {
	for i, _ := range users {
		if users[i].ID == id {
			//[:i] - pass index dynamically
			users = append(users[:i], users[i+1:]...)
			return true
		}
	}
	return false
}

func main() {
	createUser("John", 30)
	createUser("Jane", 30)
	all := ReadAll()
	fmt.Println(all)
	//user, ok := ReadByID(1)
	//fmt.Println(user, ok)

	//if user is available, the variable value would be user=data,ok=true
	user, ok := ReadByID(10)
	if ok {
		fmt.Println(user.Name, user.Age, ok)
	} else {
		fmt.Println("not found", ok, user)
	}
	//if user is not available, the variable value would be user=nil,ok=false
	//short cut
	if user, ok := ReadByID(1); ok {
		fmt.Println(user.Name, user.Age, ok)
	} else {
		fmt.Println("not found", ok, user)
	}
	//delete By id
	if value := DeleteByID(10); value {
		fmt.Println("The user has been deleted")
	} else {
		fmt.Printf("The user not deleted\n")
	}

}

//func main() {
//	slice := make([]int, 0)
//	fmt.Println("slice", slice)
//	slice = append(slice, 1)
//	fmt.Println("slice", slice)
//}
