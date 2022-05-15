package repository

import "gorm.io/gorm"

// User 用户模型
type User struct {
	Id            int64
	Name          string
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
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
		// util.Logger.Error("find user by id err:" + err.Error())
		return nil, err
	}

	return &user, nil
}
