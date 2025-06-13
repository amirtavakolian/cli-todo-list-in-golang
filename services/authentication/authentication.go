package authentication

import (
	"cli_todo/services/response"
	"cli_todo/helper"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	user.Password = string(hashedPassword)

	readUsersFile, _ := os.ReadFile(USERS_DATABASE_FILE)
	json.Unmarshal(readUsersFile, &oldUsers)

	user.Id = len(oldUsers) + 1

	for _, value := range oldUsers {
		fmt.Println(value)
		if value.User == user.User {
			responseMsg.SetContent("User is available")
			responseMsg.SetStatus(400)
			return responseMsg.BuildResponse()
		}
	}

	oldUsers = append(oldUsers, *user)

	openUsersFile, openUsersFileErr := os.OpenFile(USERS_DATABASE_FILE, os.O_RDWR, 0777)

	if openUsersFileErr != nil {
		log.Fatal("Error in opening users file")
	}

	marshalUser, _ := json.Marshal(oldUsers)

	openUsersFile.Write(marshalUser)

	openUsersFile.Close()

	responseMsg.SetContent("You have registered successfully")
	responseMsg.SetStatus(200)
	return responseMsg.BuildResponse()

}

func (user *User) Login() any {

	var username, password string
	var allUsers []User
	isUserFound, isPasswordFound := false, false
	currentUser := User{}

	fmt.Print("Enter username: ")
	fmt.Scanln(&username)

	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	users, readUsersErr := os.ReadFile(USERS_DATABASE_FILE)

	json.Unmarshal(users, &allUsers)

	if readUsersErr != nil {
		log.Fatal("Error in reading users file")
	}

	for _, value := range allUsers {
		if value.User == username {
			isUserFound = true
		}

		checkPassResult := bcrypt.CompareHashAndPassword([]byte(value.Password), []byte(password))

		if checkPassResult == nil {
			isPasswordFound = true
		}

		if isUserFound && isPasswordFound {
			currentUser = value
			break
		}
	}

	if !isUserFound || !isPasswordFound {
		responseMsg := response.Response{}
		responseMsg.SetContent("username or password is wrong")
		responseMsg.SetStatus(401)
		return responseMsg.BuildResponse()
	}

	user.insertCookie(currentUser)

	responseMsg := response.Response{}
	responseMsg.SetContent("You have Loged in successfully")
	responseMsg.SetStatus(200)
	return responseMsg.BuildResponse()
}

func (user *User) insertCookie(currentUser User) {

	var users []User
	var oldUsers []User

	allUsers, _ := os.ReadFile("users.json")
	json.Unmarshal(allUsers, &users)

	for _, value := range users {
		if currentUser.User == value.User {
			value.Cookie = helper.GetUserHardwareSerial()
		}
		oldUsers = append(oldUsers, value)
	}

	usersFile, _ := os.OpenFile("users.json", os.O_WRONLY|os.O_TRUNC, 0644)

	cc, _ := json.Marshal(&oldUsers)

	_, _ = usersFile.Write(cc)
}
