package dto

import "time"

type EventUserPayload[T any] struct {
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
	Data      T         `json:"data"`
}
type EventUserData struct {
	UserId        string    `json:"user_id"`
	Username      string    `json:"username"`
	Registered_at time.Time `json:"registered_at"`
}
type EventUserAvatar struct {
	UserId    int64  `json:"user_id"`
	AvatarUrl string `json:"avatar_url"`
}
