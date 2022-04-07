package category

import (
	"github.com/taoqun8316/goblog/pkg/logger"
	"github.com/taoqun8316/goblog/pkg/model"
	"github.com/taoqun8316/goblog/pkg/route"
)

func (category *Category) Create() (err error) {
	if err = model.DB.Create(&category).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

func All() ([]Category, error) {
	var categories []Category
	if err := model.DB.Find(&categories).Error; err != nil {
		return categories, err
	}
	return categories, nil
}

func (c Category) Link() string {
	return route.Name2URL("categories.show", "id", c.GetStringID())
}
