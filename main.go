package main

import (
	"cli_todo/app"
)

func main() {

	var application app.App

	application.Bootstrap()
	application.ShowStartupMenu()

}
