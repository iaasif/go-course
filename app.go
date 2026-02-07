package main

import "fmt"

type userStruct struct {
	fName string
	lName string
	bDate string
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// var appUser userStruct
	// appUser = userStruct{
	// 	fName: userFirstName,
	// 	lName: userLastName,
	// 	bDate: userBirthdate,
	// }

	appUser := userStruct{
		fName: userFirstName,
		lName: userLastName,
		bDate: userBirthdate,
	}

	outPutUserDetails(appUser)

}

func outPutUserDetails(u userStruct) {
	fmt.Println(u.fName, u.lName, u.bDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
