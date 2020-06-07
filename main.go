package main

import (
	"fmt"
	"net/url"
	"time"

	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	router *gin.Engine
)

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
type Users struct {
	gorm.Model
	Name     string
	Username string
	Password string
}

type CasbinRole struct {
	gorm.Model
	Ptype string
	V0    string
	V1    string
	V2    string
}

func init() {
	// Initialize  casbin adapter
	adapter := fileadapter.NewAdapter("configs/basic_policy.csv")

	// Initialize gin router
	router = gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	router.Use(cors.New(corsConfig)) // CORS configuraion
	router.POST("/user/login", handler.Login)
	resource := router.Group("/api")
	resource.Use(middleware.Authenticate())
	{
		resource.GET("/resource", middleware.Authorize("resource", "read", adapter), handler.ReadResource)
		resource.POST("/resource", middleware.Authorize("resource", "write", adapter), handler.WriteResource)
	}
}
func main() {
	dsn := url.URL{
		User:     url.UserPassword("postgres", "admin"),
		Scheme:   "postgres",
		Host:     "127.0.0.1:5433",
		Path:     "usertest",
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	db, err := gorm.Open("postgres", dsn.String())
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("conected")

	db.AutoMigrate(&CasbinRole{}, &Users{})
	defer db.Close()

}
