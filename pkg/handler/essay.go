package handler

import (
	"log"
	"net/http"

	"github.com/Marityr/gopitman"
	"github.com/gin-gonic/gin"
)

// @Summary Create Essay
// @Tags essay
// @Description create essay
// @ID create-essay
// @Accept  json
// @Produce  json
// @Param input body gopitman.Essay true "create essay"
// @Router /api/v1/essay [post]
func (h *Handler) createEssay(c *gin.Context) {
	var input gopitman.Essay
	if err := c.BindJSON(&input); err != nil {
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}

	err := h.services.Essay.Create(input)
	if err != nil {
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status create": "ok",
	})
}
