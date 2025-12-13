package http

import (
	"devwithsmile/gin-ecommerce/internal/auth"
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
			//public
			customer.POST("/signup", deps.CustomerHandler.Signup)
			customer.POST("/login", deps.CustomerHandler.Login)

			// protected
			customer.GET("/:email", auth.TokenMiddleware(), deps.CustomerHandler.GetCustomerByEmail)
		}
	}
	return r
}
