package http

import (
	"devwithsmile/gin-ecommerce/internal/customer"
	"log"
)

type AppDeps struct {
	CustomerHandler customer.Handler
}

func buildDeps() AppDeps {
	customeRrepo := customer.NewMemRepository()
	customerService := customer.NewService(customeRrepo)
	custHandler := customer.NewHandler(customerService)

	return AppDeps{
		CustomerHandler: custHandler,
	}
}

func Run() {
	appDeps := buildDeps()
	r := newRouter(RouteDeps(appDeps))

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
