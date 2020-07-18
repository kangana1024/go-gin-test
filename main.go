package main

import (
	"fmt"
	"log"
	"net/url"

	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kangana1024/go-gin-test/handlers"
	"github.com/kangana1024/go-gin-test/middleware"
	"github.com/kangana1024/go-gin-test/models"
)

var (
	router *gin.Engine
)

func init() {
	// Initialize  casbin adapter
	adapter := fileadapter.NewAdapter("configs/basic_policy.csv")

	// Initialize gin router
	router = gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	router.Use(cors.New(corsConfig)) // CORS configuraion
	router.POST("/user/login", handlers.Login)
	resource := router.Group("/api")
	resource.Use(middleware.Authenticate())
	{
		resource.GET("/resource", middleware.Authorize("resource", "read", adapter), handlers.ReadResource)
		resource.POST("/resource", middleware.Authorize("resource", "write", adapter), handlers.WriteResource)
	}
}
func main() {

	dsn := url.URL{
		User:     url.UserPassword("postgres", ""),
		Scheme:   "postgres",
		Host:     "127.0.0.1:5432",
		Path:     "usertest",
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	db, err := gorm.Open("postgres", dsn.String())
	defer db.Close()
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("conected")

	err = router.Run(":8081")
	if err != nil {
		log.Fatalln(fmt.Errorf("faild to start Gin application: %w", err))
	}
	db.AutoMigrate(&models.CasbinRole{}, &models.Users{})

}
