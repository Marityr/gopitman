package handler

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		err := errors.New("empty auth header")
		log.Println(err)
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	headerPath := strings.Split(header, " ")
	if len(headerPath) != 2 {
		err := errors.New("invalid auth header")
		log.Println(err)
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerPath[1])
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	c.Set(userCtx, userId)
}

func getCustomerId(c *gin.Context) (uuid.UUID, error) {
	id, ok := c.Params.Get("id")
	if !ok {
		return uuid.Nil, errors.New("user id not found")
	}

	idconvert, err := uuid.FromString((id))
	if err != nil {
		return uuid.Nil, errors.New("user id is of invalid type")
	}

	idconvert = uuid.Must(idconvert, err)

	return idconvert, nil
}
