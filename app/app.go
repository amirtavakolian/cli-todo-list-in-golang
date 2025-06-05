package app

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type App struct {
}

func (app App) Bootstrap() {
	app.createUsersFile()
}

func (app App) createUsersFile() {
	_, err := os.Stat("users.json")

	if err != nil {
		_, createFileErr := os.OpenFile("users.json", os.O_CREATE, 0777)

		if errors.Is(createFileErr, err) {
			log.Fatal("Error in creating file")
		}
	}
}

func (app App) ShowStartupMenu() {

	var selectOption int

	for {
		fmt.Println("\n1- Register\n2- Login\n")
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
}
