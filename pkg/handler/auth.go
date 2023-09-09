package handler

import (
	"github.com/gin-gonic/gin"
	todo "github.com/tuanneli/REST-api.git"
	"net/http"
)

func (h *Handler) singIn(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

}

func (h *Handler) singUp(c *gin.Context) {

}
