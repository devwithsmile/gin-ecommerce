package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) Handler {
	return Handler{svc: svc}
}

func (h Handler) Signup(c *gin.Context) {
	var req SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "err": err.Error()})
		return
	}

	customer, err := h.svc.Signup(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "err": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data": gin.H{
			"email": customer.Email,
		},
	})

}

func (h Handler) Login(c *gin.Context) {
	// will implement after JWT & repo
	println("Login handler reached")

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	token, err := h.svc.Login(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.SetCookie("refresh-token", token.RefreshToken, 24, "", "", true, true)
	c.JSON(http.StatusOK, gin.H{
		"token": token.AccessToken,
	})

}

func (h Handler) GetCustomerByEmail(c *gin.Context) {
	email := c.Param("email")

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "email is required",
		})
		return
	}

	customer, err := h.svc.GetCustomer(email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"name":  customer.Name,
			"id":    customer.ID,
			"phone": customer.Phone,
			"email": customer.Email,
		},
	})

}
