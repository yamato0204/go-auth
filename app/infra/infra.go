package infra

import (
	"github.com/yamato0204/go-up/app/entity"
    "gorm.io/gorm"

)

type Infra interface {
	
	CreateUser(user *entity.User) error
	GetUserByEmail(user *entity.User, email string) error
}

type infra struct {
	db *gorm.DB
}


func NewInfra(db *gorm.DB) Infra {
	return &infra{db} //暗黙的　構造体を返す
}




func  (i *infra)CreateUser(user *entity.User) error {
	if err := i.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

 func (i *infra)GetUserByEmail(user *entity.User, email string) error {
	if err := i.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
 }
