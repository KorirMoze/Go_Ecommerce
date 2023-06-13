package main

import (
	"Ecommerce/controllers"
	"Ecommerce/database"
	"Ecommerce/middleware"
	"Ecommerce/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main (){
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"),database.UserData(database.Client, "Users"))

	router := gin.new()
	router.Use(gin.logger())
	routes.UserRoutes(router)
	router.User(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.ButFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":"+port))
}
