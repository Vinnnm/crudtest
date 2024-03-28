package initializers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DisplayIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func DisplayAddPage(c *gin.Context) {
	c.HTML(http.StatusOK, "add.html", nil)
}

func DisplayViewPage(c *gin.Context) {
	c.HTML(http.StatusOK, "view.html", nil)
}

func DisplayUpdatePage(c *gin.Context) {
	c.HTML(http.StatusOK, "update.html", nil)
}
