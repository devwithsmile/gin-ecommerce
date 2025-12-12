package http

import (
	"devwithsmile/gin-ecommerce/internal/customer"
	"log"
)

func Run() {
	repo := customer.NewMemRepository()
	svc := customer.NewService(repo)
	custHandler := customer.NewHandler(svc)

	r := newRouter(RouteDeps{CustomerHandler: custHandler})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
