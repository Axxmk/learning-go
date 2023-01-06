package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook  string `json:"facebook,omitempty"`
	Twitter   string `json:"twitter,omitempty"`
	Instagram string `json:"instagram,omitempty"`
}

func main() {
	//* Open JSON file
	jsonFile, err := os.Open("./playground/users.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully Opened users.json")

	//* Defer the closing of JSON file so that we can parse it later on
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Successfully Closed users.json")
	}(jsonFile)

	//* Read the opened jsonFile as a byte array
	byteValue, _ := io.ReadAll(jsonFile)

	//* Unmarshal our byteArray which contains the jsonFile's content into Users
	var users Users
	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users.Users {
		fmt.Println("User Name: ", user.Name)
		fmt.Println("User Type: ", user.Type)
		fmt.Println("User Age: ", user.Age)
		fmt.Println("User FaceBook Url: ", user.Social.Facebook)
	}
}
