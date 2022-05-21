package repository

import (
	"sync"

	"github.com/yintamaa/998244353Plus1/utils"
)

type Vedio struct {
	Id            int64  `gorm:"column:id"`
	AuthorId      int64  `json:"author"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	CreateTime    int64
}

func (Vedio) Tablename() string {
	return "video"
}

type VedioDao struct {
}

var vedioDao *VedioDao
var vedioOnce sync.Once

func NewVedioDaoInstance() *VedioDao {
	vedioOnce.Do(
		func() {
			vedioDao = &VedioDao{}
		})
	return vedioDao
}

func (*VedioDao) QueryVedioById(id int64) (*Vedio, error) {
	var vedio Vedio
	err := db.Where("id = ?", id).Find(&vedio).Error
	if err != nil {
		utils.GetLoggerInstance().Error("find vedio by id err:" + err.Error())
		return nil, err
	}
	return &vedio, nil
}

// 返回在lastest_time之前的30条视频
func (*VedioDao) QueryVedioRecent(latest_time int64) ([]*Vedio, error) {
	var vedios []*Vedio
	err := db.Where("CreateTime > ?", latest_time).Limit(30).Order("CreateTime DESV").Find(&vedios).Error
	if err != nil {
		utils.GetLoggerInstance().Error("find vedio err:" + err.Error())
		return nil, err
	}
	return vedios, nil
}
