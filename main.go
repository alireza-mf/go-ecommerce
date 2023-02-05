package main

import (
	"log"

	"github.com/alireza-mf/go-ecommerce/db"
	"github.com/alireza-mf/go-ecommerce/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// Server represents server
type Server struct {
	Instance    *mongo.Database
	Port        string
	ServerReady chan bool
}

func main() {
	// port := configs.Server.Port
	// if port == "" {
	//  	port = "8000"
	// }

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

	log.Fatal(r.Run())
}
