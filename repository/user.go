package repository

import (
	"github.com/yintamaa/998244353Plus1/utils"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	Id            int64  `gorm:"column:id"`
	Name          string `gorm:"column:name"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:follower_count"`
	IsFollow      bool   `gorm:"is_follow"`
}

func (User) Tablename() string {
	return "user"
}

func (*User) QueryUserById(id int16) (*User, error) {
	var user User

	err := db.Where("id = ?", id).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if err != nil {
		logger := utils.GetLoggerInstance()
		logger.Error("find user by id err:" + err.Error())
		return nil, err
	}

	return &user, nil
}
