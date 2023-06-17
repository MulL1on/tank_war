package mysql

import (
	"gorm.io/gorm"
	"tank_war/server/shared/errno"
)

type User struct {
	ID       int64  `gorm:"column:id;primary_key" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Kill     int    `gorm:"column:kill" json:"kill"`
	Death    int    `gorm:"column:death" json:"death"`
}

type UserManager struct {
	salt string
	db   *gorm.DB
}

func NewUserManager(db *gorm.DB, salt string) *UserManager {
	return &UserManager{
		db:   db,
		salt: salt,
	}
}

func (m *UserManager) CreateUser(user *User) error {
	err := m.db.Create(&user).Error
	if err != nil {
		return err
	}
	return err
}

func (m *UserManager) GetUserByUsername(username string) (*User, error) {
	var u User
	err := m.db.Where(&User{Username: username}).First(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.RecordNotFound
		}
		return nil, err
	}
	return &u, nil
}

func (m *UserManager) GetUserById(id int64) (*User, error) {
	var u User
	err := m.db.Where(&User{ID: id}).First(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.RecordNotFound
		}
		return nil, err
	}
	return &u, nil
}
