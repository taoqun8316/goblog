package article

import (
	"github.com/taoqun8316/goblog/pkg/model"
	"github.com/taoqun8316/goblog/pkg/types"
)

func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToUint64(idstr)
	if err := model.DB.Preload("User").First(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil
}

func GetAll() ([]Article, error) {
	var articles []Article
	if err := model.DB.Debug().Preload("User").Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}
