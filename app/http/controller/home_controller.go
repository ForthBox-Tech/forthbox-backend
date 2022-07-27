package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeIndex(c *gin.Context) {
	c.String(http.StatusOK, "Forth Box.")
}


