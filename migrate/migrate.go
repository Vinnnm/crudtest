package main

import (
	"github.com/Vinnnm/golang-crud/initializers"
	"github.com/Vinnnm/golang-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
