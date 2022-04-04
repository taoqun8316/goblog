package article

import (
	"github.com/taoqun8316/goblog/app/models"
	"github.com/taoqun8316/goblog/pkg/logger"
	"github.com/taoqun8316/goblog/pkg/model"
	"github.com/taoqun8316/goblog/pkg/route"
	"github.com/taoqun8316/goblog/pkg/types"
)

type Article struct {
	models.BaseModel

	Title string `gorm:"type:varchar(255);not null;" valid:"title"`
	Body  string `gorm:"type:longtext;not null;" valid:"body"`
}

func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToUint64(idstr)
	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}

func GetAll() ([]Article, error) {
	var articles []Article
	if err := model.DB.Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

func (article *Article) Create() (err error) {
	if err = model.DB.Create(&article).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

func (article *Article) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&article)
	if err := result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}
	return result.RowsAffected, nil
}

func (article *Article) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(article)
	if err := result.Error; err != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}

func (article Article) Link() string {
	return route.Name2URL("article.show", "id", types.Uint64ToString(article.ID))
}
