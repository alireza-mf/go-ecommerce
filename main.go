package main

import (
	"log"

	"github.com/alireza-mf/go-ecommerce/config"
	"github.com/alireza-mf/go-ecommerce/db"
	"github.com/alireza-mf/go-ecommerce/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// todos: register, login with jwt, omit password in get user, req validation, get products, get product, add to card product, confirm card, add user roles and create/update product, app errors

// Server represents server
type Server struct {
	Instance    *mongo.Database
	Port        string
	ServerReady chan bool
}

func main() {
	config.InitConfig()

	r := gin.Default()

	// CORS allows all origins
	conf := cors.DefaultConfig()
	conf.AllowAllOrigins = true
	r.Use(cors.New(conf))

	// init mongodb
	mongodb, err := db.InitMongoDB()
	if err != nil {
		panic(err)
	}

	// init routes
	userAPI := initUserAPI(mongodb)
	routers.UserRouter(r, userAPI)

	address := ":" + config.GetConfig().Port
	log.Fatal(r.Run(address))
}
