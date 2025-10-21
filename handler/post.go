package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"go_blog/database"
	renders "go_blog/render"
	"log"
	"net/http"
)

func Login(c *gin.Context, db_user *sql.DB) {
	type structLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type LoginResponse struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Token   string `json:"token"`
		Name    string `json:"name"`
		Email   string `json:"email"`
	}
	var login structLogin
	if c.BindJSON(&login) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid request",
		})
		return
	}
	if database.Login(db_user, login.Email, login.Password) {
		c.JSON(http.StatusOK, gin.H{
			"msg":   "log in success",
			"token": database.GenerateToken(login.Email),
			"name":  database.GetUserByEmail(db_user, login.Email).Name,
			"email": login.Email,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "log in failed",
		})
	}
}

func VerifyToken(c *gin.Context, db_user *sql.DB) {
	// get auth header
	auth := c.Request.Header.Get("Authorization")
	user := database.GetUserByAuthHeader(db_user, auth)
	if user.Email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "invalid token",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "valid token",
		})
	}

}

func GetPost(c *gin.Context, db_user *sql.DB, db_post *sql.DB) {
	type GetPostRequest struct {
		Url      string `json:"url"`
		Rendered bool   `json:"rendered"`
	}
	type GetPostResponse struct {
		Status  string              `json:"status"`
		Message string              `json:"message"`
		Post    database.V4PostData `json:"post"`
		Html    string              `json:"html"`
	}
	var postRequest GetPostRequest
	var response GetPostResponse
	//if c.BindJSON(&GetPostRequest{}) != nil {
	if c.BindJSON(&postRequest) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid request",
		})
		return
	}
	user, _ := c.Get("user")
	// get post
	postData, err := database.V4GetPostByUrlUser(db_post, postRequest.Url, user.(database.User))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "permission denied",
		})
		return
	}
	response.Status = "success"
	response.Message = "ok"
	response.Post = postData
	response.Html = string(renders.RenderMd([]byte(postData.Content)))
	c.JSON(http.StatusOK, response)
}

func UpdatePost(c *gin.Context, db_user *sql.DB, db_post *sql.DB) {
	type PostUpdateRequest database.V4PostData
	type GetPostResponse struct {
		Status  string              `json:"status"`
		Message string              `json:"message"`
		Post    database.V4PostData `json:"post"`
		Html    string              `json:"html"`
	}
	user, _ := c.Get("user")
	log.Println("UpdatePost: user: ", user)
	var postUpdateRequest PostUpdateRequest
	if c.BindJSON(&postUpdateRequest) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid request",
		})
		return
	}
	log.Println("UpdatePost: postUpdateRequest: ", postUpdateRequest)
	err := database.V4UpdatePosByUser(db_post, database.V4PostData(postUpdateRequest), user.(database.User))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "permission denied",
		})
		return
	}
	postData, err := database.V4GetPostByUrlUser(db_post, postUpdateRequest.Url, user.(database.User))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "permission denied",
		})
		return
	}
	c.JSON(http.StatusOK, GetPostResponse{
		Status:  "success",
		Message: "ok",
		Post:    postData,
		Html:    string(renders.RenderMd([]byte(postData.Content))),
	})
}

func NewPost(c *gin.Context, dbUser *sql.DB, dbPost *sql.DB) {
	type NewPostResponse struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Url     string `json:"url"`
	}
	user, _ := c.Get("user")
	url, err := database.V4NewPostUser(dbPost, user.(database.User))
	if err != nil {
		c.JSON(http.StatusForbidden, NewPostResponse{
			Message: "permission denied",
		})
		return
	}
	c.JSON(http.StatusOK, NewPostResponse{
		Status:  "success",
		Message: "ok",
		Url:     url,
	})
}

func GetDistinct(c *gin.Context, dbUser *sql.DB, dbPost *sql.DB) {
	type GetDistinctRequest struct {
		Field string `json:"field"`
	}
	type GetDistinctResponse struct {
		Status  string   `json:"status"`
		Message string   `json:"message"`
		Values  []string `json:"values"`
		Length  int      `json:"length"`
	}
	user, _ := c.Get("user")
	var request GetDistinctRequest
	if c.BindJSON(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid request",
		})
		return
	}
	print(request.Field)
	values, err := database.V4GetDistinctUser(dbPost, request.Field, user.(database.User))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "permission denied",
		})
		return
	}
	c.JSON(http.StatusOK, GetDistinctResponse{
		Status:  "success",
		Message: "ok",
		Values:  values,
		Length:  len(values),
	})
}
