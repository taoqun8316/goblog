package database

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/taoqun8316/goblog/pkg/logger"
)

func Initiallize() {
	initDB()
	createTables()
}

var DB *sql.DB

func initDB() {
	var err error
	config := mysql.Config{
		User:                 "root",
		Passwd:               "taoqun8316",
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		DBName:               "goblog",
		AllowNativePasswords: true,
	}

	DB, err = sql.Open("mysql", config.FormatDSN())
	logger.LogError(err)

	// 设置最大连接数
	DB.SetMaxOpenConns(25)
	// 设置最大空闲连接数
	DB.SetConnMaxIdleTime(25)
	// 设置每个链接的过期时间
	DB.SetConnMaxLifetime(5 * time.Minute)

	// 尝试连接，失败会报错
	err = DB.Ping()
	logger.LogError(err)
}

func createTables() {
	createArticlesSQL := `CREATE TABLE IF NOT EXISTS articles(
		id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
		body longtext COLLATE utf8mb4_unicode_ci
	); `
	_, err := DB.Exec(createArticlesSQL)
	logger.LogError(err)
}
