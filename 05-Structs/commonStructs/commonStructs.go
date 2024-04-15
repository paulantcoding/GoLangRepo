package commonStructs

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// embedded struct
type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}

func NewUser(userFirstName, userLastName, userBirthDate string) (*User, error) {
	if userFirstName == "" || userLastName == "" || userBirthDate == "" {
		return nil, errors.New("FirstName, LastName and BirthDate are required fields")
	} else {

		return &User{
			firstName: userFirstName,
			lastName:  userLastName,
			birthDate: userBirthDate,
			createdAt: time.Now(),
		}, nil
	}
}

func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
