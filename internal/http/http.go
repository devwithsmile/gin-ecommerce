package http

import (
	"devwithsmile/gin-ecommerce/internal/auth"
	"devwithsmile/gin-ecommerce/internal/customer"
	"log"
)

type AppDeps struct {
	CustomerHandler customer.Handler
	AuthHanlder     auth.Handler
}

func buildDeps() AppDeps {
	//customer dependencies
	customeRrepo := customer.NewMemRepository()
	customerService := customer.NewService(customeRrepo)
	custHandler := customer.NewHandler(customerService)

	//auth dependencies
	authHandler := auth.NewHandler()
	return AppDeps{
		CustomerHandler: custHandler,
		AuthHanlder:     authHandler,
	}
}

func Run() {
	appDeps := buildDeps()
	r := newRouter(RouteDeps(appDeps))

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
