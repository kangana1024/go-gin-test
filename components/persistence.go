package components

import (
	"fmt"
	"time"

	"github.com/allegro/bigcache"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	DB          *gorm.DB
	GlobalCache *bigcache.BigCache
)

func init() {
	// Connect to DB
	//var err error
	//DB, err = gorm.Open("mysql", "your_db_url")
	//if err != nil {
	//	panic(fmt.Errorf("failed to connect to DB: %w", err))
	//}

	// Initialize cache
	var err error
	GlobalCache, err = bigcache.NewBigCache(bigcache.DefaultConfig(30 * time.Minute))
	if err != nil {
		panic(fmt.Errorf("failed to initialize cahce: %w", err))
	}
}
