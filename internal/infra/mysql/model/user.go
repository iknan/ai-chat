package model

import "gorm.io/gorm"

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}

// 根据手机号获取用户信息
func (u *UserModel) GetUserByPhone(phone string) (*User, error) {
	var user User
	err := u.db.Model(&User{}).Where("phone = ?", phone).First(&user).Error
	return &user, err
}

// 插入User
func (u *UserModel) InsertUser(user *User) error {
	return u.db.Model(&User{}).Omit("create_time", "update_time").Create(&user).Error
}

// 更新User
func (u *UserModel) UpdateUserByPhone(user User) error {
	return u.db.Model(&User{}).Where("phone = ?", user.Phone).Updates(&user).Error
}
