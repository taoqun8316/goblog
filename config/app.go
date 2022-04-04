package config

import "github.com/taoqun8316/goblog/pkg/config"

func init() {
	config.Add("app", config.StrMap{

		"name": config.Env("APP_NAME", "GoBlog"),

		// 当前环境，用以区分多环境
		"env": config.Env("APP_ENV", "production"),

		// 是否进入调试模式
		"debug": config.Env("APP_DEBUG", false),

		"port": config.Env("APP_PORT", ":3000"),

		"key": config.Env("APP_KEY", "33446a9dcf9ea060a0a6532b166da32f304af0de"),
	})
}
