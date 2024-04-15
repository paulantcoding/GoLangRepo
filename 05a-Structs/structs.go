package main

import (
	"fmt"

	comStructs "example.com/structs/commonStructs"
)

func main() {

	userFirstName := getUserData("Please enter your first name : ")
	userLastName := getUserData("Please enter your last name : ")
	userBirthDate := getUserData("Please enter your date of birth ((DD/MM/YYYY)) : ")

	var userStruct *comStructs.User
	userStruct, err := comStructs.NewUser(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Print(err)
		return
	}

	adminStruct := comStructs.NewAdmin("dummy@email.com", "pasword")
	adminStruct.OutputUserDetails()

	userStruct.OutputUserDetails()
	userStruct.ClearUserName()
	userStruct.OutputUserDetails()
}

func getUserData(displayText string) string {
	fmt.Print(displayText)
	var uiInput string
	fmt.Scanln(&uiInput)
	return uiInput
}
