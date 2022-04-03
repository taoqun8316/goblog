package bootstrap

import (
	"time"

	"github.com/taoqun8316/goblog/app/models/article"
	"github.com/taoqun8316/goblog/app/models/user"
	"github.com/taoqun8316/goblog/pkg/model"
	"gorm.io/gorm"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() {
	db := model.ConnectDB()

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(25)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	// 创建和维护数据表结构
	migration(db)

}

func migration(db *gorm.DB) {

	// 自动迁移
	db.AutoMigrate(
		&user.User{},
		&article.Article{},
	)
}
