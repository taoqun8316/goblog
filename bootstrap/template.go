package bootstrap

import (
	"embed"

	"github.com/taoqun8316/goblog/pkg/view"
)

// SetupTemplate 模板初始化
func SetupTemplate(tmplFS embed.FS) {

	view.TplFS = tmplFS

}
