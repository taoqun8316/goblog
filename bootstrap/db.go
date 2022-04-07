package bootstrap

import (
	"time"

	"github.com/taoqun8316/goblog/app/models/article"
	"github.com/taoqun8316/goblog/app/models/category"
	"github.com/taoqun8316/goblog/app/models/user"
	"github.com/taoqun8316/goblog/pkg/config"
	"github.com/taoqun8316/goblog/pkg/model"
	"gorm.io/gorm"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() {
	db := model.ConnectDB()

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)

	// 创建和维护数据表结构
	migration(db)

}

func migration(db *gorm.DB) {

	// 自动迁移
	db.AutoMigrate(
		&user.User{},
		&article.Article{},
		&category.Category{},
	)
}
