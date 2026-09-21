package mstruct

import (
	"fmt"

	"com.example/investement-calculator/model"
)




func Run() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *model.User
	
	appUser, err := model.New(userFirstName, userLastName, userBirthdate)
	
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	appUser.ClearUserName()
	appUser.OutputUserDetails()

	userEmail := getUserData("Please enter your email: ")
	userPassword := getUserData("Please enter your password: ")
	appAdmin, err := model.NewAdmin(userEmail, userPassword)

	if err != nil {
		fmt.Println("Error creating admin:", err)
		return
	}
	appAdmin.SetEmail(userEmail)
	appAdmin.SetPassword(userPassword)
	appAdmin.OutputUserDetails()
	fmt.Println("Admin email:", appAdmin.Email(userEmail))
	fmt.Println("Admin password:", appAdmin.Password(userPassword))
}

func getUserData(prompt string) string {
	fmt.Print(prompt)
	var value string
	//fmt.Scan(&value)
	fmt.Scanln(&value)
	return value
}