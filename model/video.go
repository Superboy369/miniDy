package model

import (
	"sync"
	"time"
)

type Video struct {
	Id            int64       `json:"id,omitempty"`
	UserInfoId    int64       `json:"-"`
	Author        UserInfo    `json:"author,omitempty" gorm:"-"` //这里应该是作者对视频的一对多的关系，而不是视频对作者，故gorm不能存他，但json需要返回它
	PlayUrl       string      `json:"play_url,omitempty"`
	CoverUrl      string      `json:"cover_url,omitempty"`
	FavoriteCount int64       `json:"favorite_count,omitempty"`
	CommentCount  int64       `json:"comment_count,omitempty"`
	IsFavorite    bool        `json:"is_favorite,omitempty"`
	Title         string      `json:"title,omitempty"`
	Users         []*UserInfo `json:"-" gorm:"many2many:user_favor_videos;"`
	Comments      []*Comment  `json:"-"`
	CreatedAt     time.Time   `json:"-"`
	UpdatedAt     time.Time   `json:"-"`
}

type VideoDAO struct {
}

var (
	VideoDao     *VideoDAO
	VideoDaoOnce sync.Once
)

func NewVideoDao() *VideoDAO {
	VideoDaoOnce.Do(func() {
		VideoDao = new(VideoDAO)
	})
	return VideoDao
}

func (*VideoDAO) CountUserVideoById(userId int64) (int64, error) {
	result := DB.Model(&Video{}).Where("user_info_id=?", userId)
	return result.RowsAffected, result.Error
}

func (*VideoDAO) CreateVideo(userId int64, playUrl string, title string) error {
	return DB.Create(&Video{
		UserInfoId: userId,
		PlayUrl:    playUrl,
		Title:      title,
	}).Error
}
