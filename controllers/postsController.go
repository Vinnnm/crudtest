package controllers

import (
	"github.com/Vinnnm/golang-crud/initializers"
	"github.com/Vinnnm/golang-crud/models"
	"github.com/gin-gonic/gin"
)

// Reference page : https://gorm.io/docs/

// https://gorm.io/docs/create.html
// CREATE POST
func PostsCreate(ctx *gin.Context) {
	// Get data off req body
	var body struct {
		Title string
		Body  string
	}

	ctx.Bind(&body)

	// Create a post
	//post := models.Post{Title: "Mingalar Par", Body: "Post body"}
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	// Return it
	ctx.JSON(200, gin.H{
		"post": post,
	})
}

// https://gorm.io/docs/query.html
// FIND ALL POSTS
func PostsIndex(ctx *gin.Context) {
	// Get the Posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	ctx.JSON(200, gin.H{
		"posts": posts,
	})
}

// https://gorm.io/docs/query.html
// FIND POST BY ID
func PostsShow(ctx *gin.Context) {
	// Get id of url
	id := ctx.Param("id")

	// Find the Post
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with them
	ctx.JSON(200, gin.H{
		"post": post,
	})
}

// https://gorm.io/docs/update.html
// UPDATE POST
func PostsUpdate(ctx *gin.Context) {
	// Get id of url
	id := ctx.Param("id")

	// Get the data of the req body
	var body struct {
		Title string
		Body  string
	}

	ctx.Bind(&body)

	// Find the post
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		//Title: "Post body", Body: "Mingalar Par"
		Title: body.Title,
		Body:  body.Body,
	})

	// Respond with it
	ctx.JSON(200, gin.H{
		"post": post,
	})
}

// https://gorm.io/docs/delete.html
// DELETE POST
func PostsDelete(ctx *gin.Context) {
	// Get id of url
	id := ctx.Param("id")

	// Find the Post
	var post models.Post
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		ctx.JSON(404, gin.H{"error": "Post not found."})
		return
	}
	// Delete the posts
	result = initializers.DB.Delete(&models.Post{}, id)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete post."})
		return
	}
	// Respond
	ctx.JSON(200, gin.H{"message": "Post deleted successfully"})
}
