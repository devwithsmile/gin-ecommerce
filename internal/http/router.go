package http

import (
	"devwithsmile/gin-ecommerce/internal/customer"

	"github.com/gin-gonic/gin"
)

type RouteDeps struct {
	CustomerHandler customer.Handler
}

func newRouter(deps RouteDeps) *gin.Engine {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	v1 := r.Group("/v1")
	{
		customer := v1.Group("/customer")
		{
			customer.POST("/signup", deps.CustomerHandler.Signup)
			customer.GET("/:email", deps.CustomerHandler.GetCustomerByEmail)
		}
	}
	return r
}
