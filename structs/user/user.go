package user

import (
	"errors"
	"fmt"
	"time"
)

// ====> This is Struct
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// =====> Here is two Methods
func (u *User) OutPutUser() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// ====> This is a Constructor function
func NewUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("FirstName, LastName and Birthdate are required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
