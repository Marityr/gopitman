package handler

import (
	"log"
	"net/http"

	"github.com/Marityr/gopitman"
	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body gopitman.User true "account info"
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input gopitman.User

	if err := c.BindJSON(&input); err != nil {
		log.Println(err)
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type sigInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary SignIn
// @Tags auth
// @Description Auth token
// @ID aunhetication
// @Accept  json
// @Produce  json
// @Param input body sigInInput true "account info"
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input sigInInput

	if err := c.BindJSON(&input); err != nil {
		log.Println(err)
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)

	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
