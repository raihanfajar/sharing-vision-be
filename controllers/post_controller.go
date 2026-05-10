package controllers

import (
	"net/http"
	"strconv"

	"sharing-vision-api/enums"
	"sharing-vision-api/services"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	service *services.PostService
}

func NewPostController(service *services.PostService) *PostController {
	return &PostController{service: service}
}

type createPostRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status"`
}

func validatePostRequest(req createPostRequest) string {
	if req.Title == "" {
		return "title is required"
	}
	if len(req.Title) < 20 {
		return "title must be at least 20 characters"
	}
	if req.Content == "" {
		return "content is required"
	}
	if len(req.Content) < 200 {
		return "content must be at least 200 characters"
	}
	if req.Category == "" {
		return "category is required"
	}
	if len(req.Category) < 3 {
		return "category must be at least 3 characters"
	}
	if req.Status == "" {
		return "status is required"
	}
	if !enums.IsValidStatus(req.Status) {
		return "status must be one of: publish, draft, thrash"
	}
	return ""
}

func (pc *PostController) CreatePost(c *gin.Context) {
	var req createPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body",
		})
		return
	}

	if msg := validatePostRequest(req); msg != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": msg,
		})
		return
	}

	if err := pc.service.CreatePost(req.Title, req.Content, req.Category, req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Post created successfully",
	})
}

func (pc *PostController) GetAllPosts(c *gin.Context) {
	limitStr := c.Param("limit")
	offsetStr := c.Param("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "limit must be a non-negative integer",
		})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "offset must be a non-negative integer",
		})
		return
	}

	posts, err := pc.service.GetAllPosts(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Posts retrieved successfully",
		"data":    posts,
	})
}

func (pc *PostController) GetPostByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "id must be a positive integer",
		})
		return
	}

	post, err := pc.service.GetPostByID(id)
	if err != nil {
		if err.Error() == "post not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": "Post not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Post retrieved successfully",
		"data":    post,
	})
}

func (pc *PostController) UpdatePost(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "id must be a positive integer",
		})
		return
	}

	var req createPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body",
		})
		return
	}

	if msg := validatePostRequest(req); msg != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": msg,
		})
		return
	}

	if err := pc.service.UpdatePost(id, req.Title, req.Content, req.Category, req.Status); err != nil {
		if err.Error() == "post not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": "Post not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Post updated successfully",
	})
}

func (pc *PostController) DeletePost(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "id must be a positive integer",
		})
		return
	}

	if err := pc.service.DeletePost(id); err != nil {
		if err.Error() == "post not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": "Post not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Post deleted successfully",
	})
}
