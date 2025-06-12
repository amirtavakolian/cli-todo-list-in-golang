package app

import (
	"cli_todo/Services/authentication"
"cli_todo/helper"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

type App struct{}

const USERS_DATABASE_FILE = "users.json"
const CATEGORIES_DATABASE_FILE = "categories.json"

func (app App) Bootstrap() {
	app.createUsersFile()
	app.createCategoriesFile()
}

func (app App) createUsersFile() {
	_, err := os.Stat(USERS_DATABASE_FILE)

	if err != nil {
		_, createFileErr := os.OpenFile(USERS_DATABASE_FILE, os.O_CREATE, 0777)

		if errors.Is(createFileErr, err) {
			log.Fatal("Error in creating users file")
		}
	}
}

func (app App) ShowStartupMenu() {
	for {
		isUserRegistredBefore := app.checkIfUserRegistredBefore()

		if !isUserRegistredBefore {
			app.showAuthenticationMenu()
		}

		app.showTodoListMenu()
	}
}

func (app App) showTodoListMenu() {
	var selectOption int

	fmt.Println("\n1- Create Category\n2- Create Task\n")
	fmt.Print("Select Option: ")
	fmt.Scanln(&selectOption)

	switch selectOption {
	case 1:

	case 2:

	default:
		fmt.Println("Wrong number")
		os.Exit(1)
	}
}

func (app App) showAuthenticationMenu() {
	var selectOption int
	auth := authentication.User{}

	fmt.Println("\n1- Register\n2- Login\n")
	fmt.Print("Select Option: ")
	fmt.Scanln(&selectOption)

	switch selectOption {
	case 1:
		authResult := auth.Register()
		fmt.Println(authResult)

	case 2:
		loginResult := auth.Login()
		fmt.Println(loginResult)

	default:
		fmt.Println("Wrong number")
		os.Exit(1)
	}
}

func (app App) checkIfUserRegistredBefore() bool {

	var users []authentication.User

	userHardwareSerial := helper.GetUserHardwareSerial()

	userFile, userFileErr := os.ReadFile(USERS_DATABASE_FILE)

	if userFileErr != nil {
		log.Fatal("Error in opening cookies file")
	}

	json.Unmarshal(userFile, &users)

	for _, value := range users {
		if value.Cookie == userHardwareSerial {
			return true
		}
	}
	return false
}

func (app App) createCategoriesFile() {
	_, fileExistErr := os.Stat(CATEGORIES_DATABASE_FILE)

	if fileExistErr != nil {
		_, createCategoriesFileErr := os.OpenFile(CATEGORIES_DATABASE_FILE, os.O_CREATE, 0777)

		if createCategoriesFileErr != nil {
			log.Fatal("Error in creating categories file")
		}
	}
}
