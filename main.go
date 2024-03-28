package main

import (
	"log"

	"github.com/Vinnnm/golang-crud/controllers"
	"github.com/Vinnnm/golang-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

// Reference Page : https://gin-gonic.com/docs/introduction/
func main() {

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", initializers.DisplayIndexPage)
	r.GET("/add", initializers.DisplayAddPage)
	r.GET("/view", initializers.DisplayViewPage)
	r.GET("/update/:id", initializers.DisplayUpdatePage)

	// For UI
	api := r.Group("/api")
	{
		// CREATE
		api.POST("/posts", controllers.ApiPostsCreate)

		// READ ALL
		api.GET("/posts", controllers.ApiPostsIndex)

		// READ BY ID
		api.GET("/posts/:id", controllers.ApiPostsShow)

		// UPDATE
		api.PUT("/posts/:id", controllers.ApiPostsUpdate)

		//DELETE
		api.DELETE("/posts/:id", controllers.ApiPostsDelete)
	}
	// For REST Postman
	rest := r.Group("/rest")
	{
		// CREATE
		rest.POST("/posts", controllers.PostsCreate)

		// READ ALL
		rest.GET("/posts", controllers.PostsIndex)

		// READ BY ID
		rest.GET("/posts/:id", controllers.PostsShow)

		// UPDATE
		rest.PUT("/posts/:id", controllers.PostsUpdate)

		//DELETE
		rest.DELETE("/posts/:id", controllers.PostsDelete)
	}

	// for CSS and JS
	r.Static("/static", "./static")

	//r.Run()
	log.Fatal(r.Run(":3000"))
}
