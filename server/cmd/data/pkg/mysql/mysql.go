package pkg

import (
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID       int64  `gorm:"column:id;primary_key" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Kill     int    `gorm:"column:kill" json:"kill"`
	Death    int    `gorm:"column:death" json:"death"`
}

type UserManager struct {
	db *gorm.DB
}

func NewUserManager(db *gorm.DB) *UserManager {
	return &UserManager{
		db: db,
	}
}

func (m *UserManager) UpdateUser(user *User) {
	err := m.db.Model(&user).Updates(user).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Println(err)
		}
	}
}
