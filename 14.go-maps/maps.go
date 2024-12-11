package main

import "fmt"

func main() {
	//initializing map
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])
	//updating value
	websites["Google"] = "https://mail.google.com"
	//adding new key
	websites["Microsoft"] = "https://microsoft.com"
	fmt.Println(websites)

	//remove key from the map
	delete(websites, "Google")
	fmt.Println(websites)
}
