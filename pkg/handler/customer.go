package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Marityr/gopitman"
	"github.com/gin-gonic/gin"
)

type createCustomerinput struct {
	FirstName    string
	LastName     string
	Phone        string
	Email        string
	Birthday     string
	ReferrerCode string
}

// @Summary Create customer
// @Security ApiKeyAuth
// @Tags customer
// @Description create customer
// @ID create-customer
// @Accept  json
// @Produce  json
// @Param input body createCustomerinput true "list info"
// @Router /api/v1/customer [post]
func (h *Handler) createCustomer(c *gin.Context) {
	var input createCustomerinput
	if err := c.BindJSON(&input); err != nil {
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}

	customerId, err := h.services.Customer.Create(input.FirstName, input.LastName, input.Birthday, input.ReferrerCode, input.Phone, input.Email)
	if err != nil {
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"customerId": customerId,
	})
}

// Customer
// @Summary  CustomerAll
// @Security ApiKeyAuth
// @Description  Customer
// @Tags     customer
// @Accept   json
// @Produce  json
// @Param    page  query  string  false  "Page"
// @Param    limit  query  string  false  "Limit"
// @Router   /api/v1/customer [get]
func (h *Handler) getAllCustomer(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("limit", "100"))
	if err != nil {
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "100"))
	if err != nil {
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}

	lists, err := h.services.Customer.GetAll(page, limit)
	if err != nil {
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}

	c.JSON(http.StatusOK, lists)
}

// Customer
// @Summary  Customer By Id
// @Security ApiKeyAuth
// @Description  Customer get by id
// @Tags     customer
// @Accept   json
// @Produce  json
// @Param    id  path  string  true  "id"
// @Router   /api/v1/customer/{id} [get]
func (h *Handler) getCustomerById(c *gin.Context) {
	userId, err := getCustomerId(c)
	if err != nil {
		log.Panicln(err)
		c.AbortWithError(400, err)
		return
	}

	list, err := h.services.Customer.GetById(userId)
	if err != nil {

		c.JSON(400, "not user")
		return
	}

	c.JSON(http.StatusOK, list)

}

// Customer
// @Summary  Customer Update
// @Security ApiKeyAuth
// @Description  Customer update
// @Tags     customer
// @Accept   json
// @Produce  json
// @Param    id  path  string  true  "ID"
// @Param input body gopitman.UpdateCustomer true "customer info"
// @Router   /api/v1/customer/{id} [put]
func (h *Handler) updateCustomer(c *gin.Context) {
	userId, err := getCustomerId(c)
	if err != nil {
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}

	var input gopitman.UpdateCustomer
	if err := c.BindJSON(&input); err != nil {
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}
	err = h.services.Customer.Update(userId, input)
	if err != nil {
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}

	c.JSON(http.StatusOK, "update ok!")
}

// Customer
// @Summary  Customer Delete
// @Security ApiKeyAuth
// @Description  Customer delete
// @Tags     customer
// @Accept   json
// @Produce  json
// @Param    id  path  string  true  "ID"
// @Router   /api/v1/customer/{id} [delete]
func (h *Handler) deleteCustomer(c *gin.Context) {
	userId, err := getCustomerId(c)
	if err != nil {
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}

	err = h.services.Customer.Delete(userId)
	if err != nil {
		log.Println(err)
		c.AbortWithError(400, err)
		return
	}

	c.JSON(http.StatusOK, "delete ok!")
}
