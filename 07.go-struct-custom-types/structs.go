package main

import (
	"example.com/structs/user"
)

func main() {
	// appUser, err := user.New()
	// if err != nil {
	// 	panic(err)
	// }
	// appUser.AddUserName()
	// appUser.OutputUserDetails()
	// user.OutputUserDetailsWithPointer(appUser)
	admin := user.NewAdmin()
	admin.OutputUserDetails()
	admin.AddUserName()
	admin.OutputUserDetails()
}
