package authentication

import (
	"cli_todo/response"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const USERS_DATABASE_FILE = "users.json"

type User struct {
	Id       int
	User     string
	Password string
	Cookie   string
}

func (user *User) Register() any {

	var oldUsers []User
	responseMsg := response.Response{}

	fmt.Print("Enter username: ")
	fmt.Scanln(&user.User)

	fmt.Print("Enter password: ")
	fmt.Scanln(&user.Password)

	readUsersFile, _ := os.ReadFile(USERS_DATABASE_FILE)
	json.Unmarshal(readUsersFile, &oldUsers)

	for _, value := range oldUsers {
		fmt.Println(value)
		if value.User == user.User {
			responseMsg.SetContent("User is available")
			responseMsg.SetStatus(400)
			return responseMsg.BuildResponse()
		}
	}

	oldUsers = append(oldUsers, *user)

	openFile, openFileErr := os.OpenFile(USERS_DATABASE_FILE, os.O_RDWR, 0777)

	if openFileErr != nil {
		log.Fatal("Error in opening users file")
	}

	marshalUser, _ := json.Marshal(oldUsers)

	_, _ = openFile.Write(marshalUser)

	responseMsg.SetContent("You have registered successfully")
	responseMsg.SetStatus(200)
	return responseMsg.BuildResponse()

}