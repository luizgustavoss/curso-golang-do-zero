package main

import "fmt"

func main() {

	headers := map[string]string{
		"origin" : "localhost",
		"user" : "Jon Doe",
		"id" : "123-1234-1234-5543-3445",
	}

	fmt.Println(headers)

	fmt.Println(headers["origin"])
	fmt.Println(headers["user"])
	fmt.Println(headers["id"])

	fmt.Println(len(headers))

	delete(headers, "user")

	fmt.Println(headers)
}
