package dto

import (
	"time"
	"user-management/pkg/enums"
)

type UserUpdateRequest struct {
	FirstName   *string       `json:"first_name" validate:"omitempty,max=100"`
	LastName    *string       `json:"last_name" validate:"omitempty,max=100"`
	Email       *string       `json:"email" validate:"omitempty,max=100"`
	Phone       *string       `json:"phone" validate:"omitempty,max=20"`
	BirthDate   *string       `json:"birth_date" validate:"omitempty"`
	Gender      *enums.GENDER `json:"gender" validate:"omitempty,oneof=male female other"`
	Bio         *string       `json:"bio" validate:"omitempty,max=500"`
	Description *string       `json:"description" validate:"omitempty"`
}
type UserResponse struct {
	ID           int64             `json:"id"`
	AuthID       string            `json:"auth_id"`
	Username     string            `json:"username"`
	Name         *Username         `json:"name"`
	Contact      *UserContact      `json:"contact"`
	About        *UserAbout        `json:"about"`
	Verification *UserVerification `json:"verification"`
	AvatarURL    *string           `json:"avatar_url"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
}
type Username struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}
type UserContact struct {
	Email *string `json:"email"`
	Phone *string `json:"phone"`
}
type UserVerification struct {
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	PhoneVerifiedAt *time.Time `json:"phone_verified_at"`
}
type UserAbout struct {
	Bio         *string       `json:"bio"`
	Description *string       `json:"description"`
	BirthDate   *time.Time    `json:"birth_date"`
	Gender      *enums.GENDER `json:"gender"`
}
