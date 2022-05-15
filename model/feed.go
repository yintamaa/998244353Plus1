package model

import "github.com/jinzhu/gorm"

type Video struct {
	gorm.Model
	Id            int64
	Author        User
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	IsFavorite    bool
}

type User struct {
	Id            int64
	Name          string
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
}

type Comment struct {
	Id         int64
	User       User
	Content    string
	CreateDate string
}
