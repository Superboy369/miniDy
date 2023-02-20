package model

import (
	"errors"
	"miniDy/constant"
	"sync"
	"time"
)

type Message struct {
	Id         int64
	ToUserId   int64 `json:"to_user_id"`
	FromUserId int64 `json:"from_user_id"`
	Content    string
	CreateTime time.Time `json:"create_time"`
}

type MessageDAO struct {
}

var (
	messageDao  *MessageDAO
	messageOnce sync.Once
)

func NewMessageDao() *MessageDAO {
	messageOnce.Do(func() {
		messageDao = &MessageDAO{}
	})

	return messageDao
}

func (m *MessageDAO) CreateMessage(message *Message) error {
	if message == nil {
		return errors.New("CreateMessage message 空指针")
	}

	return DB.Create(message).Error
}

func (m *MessageDAO) QueryMessages(userId, toUserId int64, lastTime time.Time, messageList *[]*Message) error {
	if messageList == nil {
		return errors.New("QueryMessages messageList 空指针")
	}
	return DB.Where("from_user_id = ? AND to_user_id = ? AND create_time  > ?", userId, toUserId, lastTime).
		Limit(constant.MAX_MESSAGE_NUMBER).Find(messageList).Error
}

func (m *MessageDAO) QueryLatestMessageById(userId, toUserId int64, message *Message) error {
	if message == nil {
		return errors.New("QueryLatestMessageById message 空指针")
	}
	return DB.Where("from_user_id = ? AND to_user_id = ?", userId, toUserId).
		Or("to_user_id = ? AND from_user_id = ?", userId, toUserId).
		Order("create_time desc").Find(message).Error
}
