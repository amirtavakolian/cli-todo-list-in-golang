package Category

import (
	"cli_todo/Services/authentication"
	"cli_todo/Services/response"
	"cli_todo/helper"
	"encoding/json"
	"log"
	"os"
)

type Category struct {
	Id      int
	User_id int
	Title   string
}

const categories_database_file = "categories.json"
const users_database_file = "users.json"

func (category *Category) Store() string {
	var categories []Category
	var responseMsg response.Response

	category.checkCategoriesDatabaseExist()

	categoriesDB, readCategoriesDB := os.ReadFile(categories_database_file)

	if readCategoriesDB != nil {
		log.Fatal("Error in reading categories.json file")
	}

	json.Unmarshal(categoriesDB, &categories)

	category.Id = len(categories) + 1

	category.User_id = category.calculateUserId()

	categories = append(categories, *category)

	marshaledCategories, _ := json.Marshal(categories)

	categoriesFile, _ := os.OpenFile(categories_database_file, os.O_RDWR, 0777)

	categoriesFile.Write(marshaledCategories)

	responseMsg.SetStatus(200)
	responseMsg.SetContent("Category created successfully")
	return responseMsg.BuildResponse()

}

func (category *Category) checkCategoriesDatabaseExist() {
	_, categoriesFileExistErr := os.Stat(categories_database_file)

	if categoriesFileExistErr != nil {
		log.Fatal("Categories.json not found")
	}
}

func (category *Category) calculateUserId() int {
	var users []authentication.User

	hardwareSerial := helper.GetUserHardwareSerial()

	usersFileData, _ := os.ReadFile(users_database_file)

	json.Unmarshal(usersFileData, &users)

	for _, value := range users {
		if value.Cookie == hardwareSerial {
			return value.Id
		}
	}
	log.Fatal("Users file is empty")
	return 0
}
