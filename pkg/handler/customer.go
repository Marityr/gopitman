package handler

import (
	"net/http"
	"strconv"

	"github.com/Marityr/gopitman/pkg/repository"
	"github.com/Marityr/gopitman/schemes"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createCustomer(c *gin.Context) {

}

// Customer
// @Summary  Customer
// @Security ApiKeyAuth
// @Description  Customer
// @Tags     Frontend
// @Accept   json
// @Produce  json
// @Param    page  query  string  false  "Page"
// @Param    limit  query  string  false  "Limit"
// @Router   /api/v1/customer [get]
func (h *Handler) getAllCustomer(c *gin.Context) {
	var cnt []schemes.Customer
	q := repository.DB

	page := c.Copy().DefaultQuery("page", "")
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "100"))
	if err != nil {
		h.logging.Info(err)
		c.AbortWithError(400, err)
		return
	}

	if page != "" {
		offset, err := strconv.Atoi(page)
		if err != nil {
			h.logging.Info(err)
			c.AbortWithError(400, err)
			return
		}
		q = q.Offset(offset * limit)
	}

	err = q.Limit(limit).Find(&cnt).Error
	if err != nil {
		h.logging.Info(err)
		c.AbortWithError(400, err)
		return
	}
	//TODO count param to json
	c.JSON(http.StatusOK, cnt)

}

func (h *Handler) getCustomerById(c *gin.Context) {

}

func (h *Handler) updateCustomer(c *gin.Context) {

}

func (h *Handler) deleteCustomer(c *gin.Context) {

}
