package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/trankyhoathanh/go_gateway/pkg/auth"
	"github.com/trankyhoathanh/go_gateway/pkg/config"
	"github.com/trankyhoathanh/go_gateway/pkg/order"
	"github.com/trankyhoathanh/go_gateway/pkg/product"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
