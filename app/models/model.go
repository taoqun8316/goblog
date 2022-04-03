package models

import "github.com/taoqun8316/goblog/pkg/types"

type BaseModel struct {
	ID uint64
}

// GetStringID 获取 ID 的字符串格式
func (a BaseModel) GetStringID() string {
	return types.Uint64ToString(a.ID)
}
