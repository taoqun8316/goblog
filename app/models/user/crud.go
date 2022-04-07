package user

import "github.com/taoqun8316/goblog/pkg/model"

func All() ([]User, error) {
	var users []User
	if err := model.DB.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
