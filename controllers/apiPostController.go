package controllers

import (
	"net/http"

	"github.com/Vinnnm/golang-crud/dao"
	"github.com/Vinnnm/golang-crud/initializers"
	"github.com/Vinnnm/golang-crud/models"
	"github.com/gin-gonic/gin"
)

// CREATE POST
func ApiPostsCreate(ctx *gin.Context) {
	// Get data from request body
	var body dao.ApiBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	// Return the created post
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Post created successfully",
		"post":    post,
	})
	// ctx.Redirect(http.StatusFound, "/index")
}

// FIND ALL POSTS
func ApiPostsIndex(ctx *gin.Context) {
	// Get the Posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	ctx.JSON(200, gin.H{
		"posts": posts,
	})
}

// FIND POST BY ID
func ApiPostsShow(ctx *gin.Context) {
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

// UPDATE POST
func ApiPostsUpdate(ctx *gin.Context) {
	// Get id of url
	id := ctx.Param("id")

	// Get the data of the req body
	var body dao.ApiBody

	ctx.ShouldBindJSON(&body)

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

// DELETE POST
func ApiPostsDelete(ctx *gin.Context) {
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
