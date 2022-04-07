package category

import (
	"github.com/taoqun8316/goblog/app/models"
	"github.com/taoqun8316/goblog/pkg/model"
	"github.com/taoqun8316/goblog/pkg/types"
)

type Category struct {
	models.BaseModel

	Name string `gorm:"type:varchar(255);not null;" valid:"name"`
}

func Get(idstr string) (Category, error) {
	var category Category
	id := types.StringToUint64(idstr)
	if err := model.DB.First(&category, id).Error; err != nil {
		return category, err
	}

	return category, nil
}
