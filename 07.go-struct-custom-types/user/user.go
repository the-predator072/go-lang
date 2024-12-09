package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
	userName  string
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin() *Admin {
	email := getUserData("Please Enter Your Email: ")
	password := getUserData("Please Enter Your password: ")
	user, err := New()
	if err != nil {
		panic(err)
	}
	admin := Admin{
		email,
		password,
		*user,
	}
	return &admin
}

func New() (*User, error) {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	if userFirstName == "" || userLastName == "" || userBirthdate == "" {
		return nil, errors.New("all fields are required")
	}
	return &User{
		userFirstName,
		userLastName,
		userBirthdate,
		time.Now(),
		"",
	}, nil
}

// method of attaching functions to structs
func (u *User) OutputUserDetails() {
	fmt.Printf("First Name: %v\n Last Name: %v\n Birth Date : %v\n", u.firstName, u.lastName, u.birthdate)
	fmt.Println("User Name:", u.userName)
}

func (u *User) AddUserName() {
	userNameStr := fmt.Sprintf("%v-%v", u.firstName, u.lastName)
	u.userName = userNameStr
}

func OutputUserDetailsWithPointer(u *User) {
	// pointer in structs are exception as we dont de reference it
	fmt.Printf("First Name: %v\n Last Name: %v\n Birth Date : %v\n", u.firstName, u.lastName, u.birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
