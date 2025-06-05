package app

import (
	"errors"
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
