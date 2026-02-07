package main

import "fmt"

type userStruct struct {
	fName string
	lName string
	bDate string
}

// makeing a function into a mthod , or , struct type method , there is no oop in go so, when
// if you not send any pointer in reciver argument, then it wont change struct value , it will change a copy of a struct value
func (u *userStruct) outPutUserDetails() {
	fmt.Println(u.fName, u.lName, u.bDate)
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// // var appUser userStruct
	// // appUser = userStruct{
	// // 	fName: userFirstName,
	// // 	lName: userLastName,
	// // 	bDate: userBirthdate,
	// // }

	appUser := userStruct{
		fName: userFirstName,
		lName: userLastName,
		bDate: userBirthdate,
	}

	appUser.outPutUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
