package mysql

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
	log.Println("update user", user)
	//get user from db
	var u = &User{}
	m.db.Model(&user).Where("username = ?", user.Username).First(&u)

	if user.Death == 1 {
		u.Death += 1
	}
	u.Kill += user.Kill

	err := m.db.Model(&u).Updates(u).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Println(err)
		}
	}
}
