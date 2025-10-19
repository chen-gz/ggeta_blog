package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	renders "go_blog/render"
	"net/http"
)

func V5Render(c *gin.Context, dbUser *sql.DB) {
	type request struct {
		Content string `json:"content"`
		Format  string `json:"format"` // can be markdown, tex
	}
	type response struct {
		Message  string `json:"message"`
		Rendered string `json:"rendered"`
		Format   string `json:"format"`
	}
	var req request
	if c.BindJSON(&req) != nil {
		c.JSON(http.StatusBadRequest, response{
			Message: "invalid request",
		})
		return
	}

	// only support markdown now
	if req.Format != "markdown" {
		c.JSON(http.StatusBadRequest, response{
			Message: "invalid format",
		})
		return
	}
	_, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, response{
			Message: "invalid token",
		})
		return
	}
	// render markdown
	rendered := string(renders.RenderMd([]byte(req.Content)))
	c.JSON(http.StatusOK, response{
		Rendered: rendered,
		Format:   "markdown",
	})
}
